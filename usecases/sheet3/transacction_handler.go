package transaction_cmd

import (
	"bail/domain/models"
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"

	icmd "bail/usecases/core/cqrs/command"
	irepo "bail/usecases/core/i_repo"
)


type Sheet3Handler struct {
	logRepo irepo.SystemLog
}

// Sheet3Config holds the configuration for creating a Sheet3Handler.
type Sheet3Config struct {
	LogRepo irepo.SystemLog
	
}

// Ensure Sheet3Handler implements icmd.IHandler.
var _ icmd.IHandler[*Sheet3Command, *models.Root] = &Sheet3Handler{}

// NewSheet3Handler creates a new instance of Sheet3Handler with the provided configuration.
func NewSheet3Handler(config Sheet3Config) *Sheet3Handler {
	return &Sheet3Handler{
		logRepo: config.LogRepo,
	}
}

func (s *Sheet3Handler) Handle(cmd *Sheet3Command) (*models.Root, error) {
	fileData := cmd.file
	excelFile, err := excelize.OpenReader(bytes.NewReader(fileData))
	if err != nil {
		return nil, fmt.Errorf("failed to open Excel file: %w", err)
	}

	sheetName := excelFile.GetSheetName(cmd.sheet)
	rows, err := excelFile.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to read rows from sheet %s: %w", sheetName, err)
	}

	// Validate if the sheet has a header row
	if len(rows) < 2 {
		return nil, fmt.Errorf("sheet %s does not contain sufficient rows", sheetName)
	}

	// Extract header and skip it
	headers := rows[0]
	rows = rows[1:]
	records := make([]*models.Record, 0, len(rows))

	for i, row := range rows {
		if len(row) < len(headers) {
			fmt.Printf("Skipping row %d: not enough columns\n", i+2) // Account for header row
			continue
		}

		// Parse branch name and code
		branchName := row[0]
		branchCode := row[1]

		// Parse financial values
		prSystem, err := parseFloat(row[2])
		if err != nil {
			fmt.Printf("Error parsing Pr System for row %d: %v\n", i+2, err)
			continue
		}
		previous, err := parseFloat(row[3])
		if err != nil {
			fmt.Printf("Error parsing Previous for row %d: %v\n", i+2, err)
			continue
		}
		withdrawal, err := parseFloat(row[4])
		if err != nil {
			fmt.Printf("Error parsing Withdrawal for row %d: %v\n", i+2, err)
			continue
		}
		slip, err := parseFloat(row[5])
		if err != nil {
			fmt.Printf("Error parsing Slip for row %d: %v\n", i+2, err)
			continue
		}
		remainingOnSystem, err := parseFloat(row[6])
		if err != nil {
			fmt.Printf("Error parsing Remaining on System for row %d: %v\n", i+2, err)
			continue
		}
		uncollected, err := parseFloat(row[7])
		if err != nil {
			fmt.Printf("Error parsing Uncollected for row %d: %v\n", i+2, err)
			continue
		}

		withdrawal = uncollected+slip 
		remainingOnSystem = prSystem + uncollected

		// Perform updates
		
		

		// Optionally update your repositories or database
		record := models.NewRecord(
			&models.RecordConfig{
				Name:              branchName,
				Code:              branchCode,
				Date:              cmd.date,
				PRSystem:          prSystem,
				Previous:          previous,
				Withdrawal:        withdrawal,
				Slip:              slip,
				RemainingOnSystem: remainingOnSystem,
				Uncollected:       uncollected,
			},
		)
		records = append(records, record)

		
		
	}

	systemLog := models.NewSystemLog(
		&models.SystemLogConfig{
			Date:    cmd.date,
			Records: records,
		},
	)

	err = s.logRepo.AddLog(systemLog)
	if err != nil {
		return nil, fmt.Errorf("failed to save system log: %w", err)
	}


	return &models.Root{}, nil
}

func parseFloat(value string) (float64, error) {
	parsed, err := strconv.ParseFloat(strings.ReplaceAll(value, ",", ""), 64)
	if err != nil {
		return 0, fmt.Errorf("invalid float value: %s", value)
	}
	return parsed, nil
}

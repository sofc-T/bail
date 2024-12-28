package transaction_cmd

import (
	"bail/domain/models"
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"

	icmd "bail/usecases/core/cqrs/command"
	irepo "bail/usecases/core/i_repo"
)


type Sheet2Handler struct {
	userRepo irepo.User
	rootRepo irepo.Root
}

// Sheet2Config holds the configuration for creating a Sheet2Handler.
type Sheet2Config struct {
	UserRepo           irepo.User
	RootRepo 		   irepo.Root
	
}

// Ensure Sheet2Handler implements icmd.IHandler.
var _ icmd.IHandler[*Sheet2Command, *models.Root] = &Sheet2Handler{}

// NewSheet2Handler creates a new instance of Sheet2Handler with the provided configuration.
func NewSheet2Handler(config Sheet2Config) *Sheet2Handler {
	return &Sheet2Handler{
		userRepo: config.UserRepo,
		rootRepo: config.RootRepo,
	}
}

func (s *Sheet2Handler) Handle(cmd *Sheet2Command) (*models.Root, error) {
	fileData := cmd.file
	sheet := cmd.sheet

	// Open Excel file from bytes
	excelFile, err := excelize.OpenReader(bytes.NewReader(fileData))
	if err != nil {
		return nil, fmt.Errorf("failed to open Excel file from bytes: %w", err)
	}

	// Validate the sheet index
	sheetName := excelFile.GetSheetName(sheet)
	if sheetName == "" {
		return nil, fmt.Errorf("sheet %d does not exist in the Excel file", sheet)
	}

	rows, err := excelFile.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to read rows from sheet %s: %w", sheetName, err)
	}

	var totalTransferred float64
	for i, row := range rows {
		// Skip the header row
		if i == 0 {
			continue
		}

		// Skip rows with insufficient data
		if len(row) < 4 {
			fmt.Printf("Skipping row %d: insufficient columns\n", i+1)
			continue
		}

		// Extract relevant columns (customize based on your actual column mapping)
		userType := strings.TrimSpace(row[0]) // Example: "branch", "employee"
		branch := strings.TrimSpace(row[1])  // Branch code or empty
		amountStr := strings.TrimSpace(row[2])
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			fmt.Printf("Skipping row %d: invalid amount %s\n", i+1, amountStr)
			continue
		}

		// Process based on user type and branch presence
		switch userType {
		case "branch":
			// Deduct from branch and update admin balance
			err := s.branchRepo.DeductBalance(branch, amount)
			if err != nil {
				fmt.Printf("Error updating branch %s balance: %v\n", branch, err)
				continue
			}
			err = s.adminRepo.UpdateBalance(amount)
			if err != nil {
				fmt.Printf("Error updating admin balance: %v\n", err)
				continue
			}

		case "employee":
			if branch != "" {
				// Deduct from user and branch, then update admin balance
				err := s.branchRepo.DeductBalance(branch, amount)
				if err != nil {
					fmt.Printf("Error updating branch %s balance: %v\n", branch, err)
					continue
				}
			}
			err := s.userRepo.DeductBalance(row[3], amount) // Assuming user ID is in column 4
			if err != nil {
				fmt.Printf("Error updating user balance: %v\n", err)
				continue
			}
			err = s.adminRepo.UpdateBalance(amount)
			if err != nil {
				fmt.Printf("Error updating admin balance: %v\n", err)
				continue
			}

		default:
			fmt.Printf("Skipping row %d: unknown user type %s\n", i+1, userType)
			continue
		}

		// Record the transaction
		transaction := models.Transaction{
			UserType: userType,
			Branch:   branch,
			Amount:   amount,
		}
		err = s.transactionRepo.Create(transaction)
		if err != nil {
			fmt.Printf("Error creating transaction for row %d: %v\n", i+1, err)
			continue
		}

		// Update total transferred
		totalTransferred += amount
	}

	// Return the total transferred and any final error
	return &models.Root{
		TotalTransferred: totalTransferred,
	}, nil
}


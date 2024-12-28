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

	// Open the Excel file
	excelFile, err := excelize.OpenReader(bytes.NewReader(fileData))
	if err != nil {
		return nil, fmt.Errorf("failed to open Excel file: %w", err)
	}

	// Validate the sheet name
	sheetName := excelFile.GetSheetName(sheet)
	if sheetName == "" {
		return nil, fmt.Errorf("sheet %d does not exist in the Excel file", sheet)
	}

	rows, err := excelFile.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to read rows from sheet %s: %w", sheetName, err)
	}

	if len(rows) < 2 {
		return nil, fmt.Errorf("sheet %s has insufficient data", sheetName)
	}

	// Extract header row (user codes)
	headerRow := rows[0]
	if len(headerRow) == 0 {
		return nil, fmt.Errorf("no user codes found in the header row")
	}

	var totalTransferred float64

	// Iterate over columns
	for colIndex, userCode := range headerRow {
		if strings.TrimSpace(userCode) == "" {
			// Stop when an empty column is encountered
			break
		}

		// Process rows in the column
		for rowIndex := 1; rowIndex < len(rows); rowIndex++ {
			row := rows[rowIndex]

			// Ensure the column exists in the row
			if len(row) <= colIndex {
				continue
			}

			amountStr := strings.TrimSpace(row[colIndex])
			if amountStr == "" {
				// Stop processing this column when an empty cell is found
				break
			}

			// Parse the amount
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				fmt.Printf("Skipping invalid amount in column %d, row %d: %s\n", colIndex+1, rowIndex+1, amountStr)
				continue
			}

			// Fetch the user by code (you'll need to implement this function)
			user, err := s.userRepo.FindByCode(userCode)
			if err != nil {
				fmt.Printf("Error fetching user for code %s in column %d: %v\n", userCode, colIndex+1, err)
				continue
			}

			// Process the deduction and updates based on user type
			switch user.Role() {
			case "branch":
				// Deduct from branch and update admin balance
				_,err = s.userRepo.AddTransaction(userCode, -1 * amount)
				if err != nil {
					fmt.Printf("Error deducting from branch %s: %v\n", user.BranchCode(), err)
					continue
				}

				err = s.rootRepo.AddTransaction(amount)
				if err != nil {
					fmt.Printf("Error updating admin balance: %v\n", err)
					continue
				}

			case "employee":
				if user.BranchCode() != "" {
					// Deduct from user and branch, then update admin balance
					_,err = s.userRepo.AddTransaction(userCode, -1 * amount)
					if err != nil {
						fmt.Printf("Error deducting from branch %s: %v\n", user.CodeNumber(), err)
						continue
					}
				}

				_,err = s.userRepo.AddTransaction(userCode, -1 * amount)
				if err != nil {
					fmt.Printf("Error deducting from user %s: %v\n", user.CodeNumber(), err)
					continue
				}

				err = s.rootRepo.AddTransaction(amount)
				if err != nil {
					fmt.Printf("Error updating admin balance: %v\n", err)
					continue
				}

			default:
				fmt.Printf("Unknown user type for code %s in column %d\n", userCode, colIndex+1)
				continue
			}

			totalTransferred += amount
		}
	}

	return &models.Root{}, nil
}

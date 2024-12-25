package transaction_cmd

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
)

func processTransactionsFromBytes(fileData []byte, sheet int) error {
	// Open the Excel file from []byte
	excelFile, err := excelize.OpenReader(bytes.NewReader(fileData))
	if err != nil {
		return fmt.Errorf("failed to open Excel file from bytes: %v", err)
	}

	// Get the first sheet
	sheetName := excelFile.GetSheetName(sheet)
	rows, err := excelFile.GetRows(sheetName)
	if err != nil {
		return fmt.Errorf("failed to read rows: %v", err)
	}

	// Process each row
	// Skip the header row
	for i, row := range rows {

		if i == 0 { 
			continue
		}

		// Limit to the first four columns
		if len(row) < 4 {
			fmt.Printf("Skipping row %d: not enough columns\n", i+1)
			continue
		}

		// Only use the first 4 columns
		limitedRow := row[:4]

		transaction, err := parseTransaction(limitedRow)
		if err != nil {
			fmt.Printf("Error parsing row %d: %v\n", i+1, err)
			continue
		}

		// Extract agent code (first part of the Agent string)
		agentCode := strings.Split(transaction.Agent, " ")[0]

		// Handle PaidIn or Withdrawal
		if transaction.PaidIn > 0 {
			fmt.Printf("Notify agent %s: Deduct %.2f from PaidIn. Update admin to new balance %.2f\n", agentCode, transaction.PaidIn, transaction.Balance)
		} else {
			fmt.Printf("Notify agent %s: Deduct %.2f from Withdrawal. Update admin to new balance %.2f\n", agentCode, transaction.Withdrawal, transaction.Balance)
		}
	}

	return nil
}

func parseTransaction(limitedRow []string) (any, any) {
	panic("unimplemented")
}

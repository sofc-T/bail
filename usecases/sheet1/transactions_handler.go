package transaction_cmd

import (
	"bail/domain/models"
	"bytes"
	"fmt"
	"strconv"
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
		agentCode := strings.Split(transaction.Agent(), " ")[0]

		// Handle PaidIn or Withdrawal
		if transaction.PaidIn() > 0 {
			fmt.Printf("Notify agent %s: Deduct %.2f from PaidIn. Update admin to new balance %.2f\n", agentCode, transaction.PaidIn(), transaction.Balance())
		} else {
			fmt.Printf("Notify agent %s: Deduct %.2f from Withdrawal. Update admin to new balance %.2f\n", agentCode, transaction.Withdrawal(), transaction.Balance())
		}
	}

	return nil
}

func parseTransaction(row []string) (models.Transaction, error) {
	var transaction models.Transaction
	var err error

	// Parse PaidIn
	paidin, err := strconv.ParseFloat(strings.ReplaceAll(row[0], ",", ""), 64)
	if err != nil && row[0] != "" {
		return transaction, fmt.Errorf("failed to parse PaidIn: %v", err)
	}
	transaction.SetPaidIn(paidin)

	// Parse Balance
	balance, err := strconv.ParseFloat(strings.ReplaceAll(row[1], ",", ""), 64)
	if err != nil {
		return transaction, fmt.Errorf("failed to parse Balance: %v", err)
	}
	transaction.SetBalance(balance)


	// Parse Withdrawal
	withdrawal, err := strconv.ParseFloat(strings.ReplaceAll(row[2], ",", ""), 64)
	if err != nil && row[2] != "" {
		return transaction, fmt.Errorf("failed to parse Withdrawal: %v", err)
	}

	transaction.SetWithdrawal(withdrawal)

	// Set Agent
	transaction.SetAgent( row[3])

	return transaction, nil
}

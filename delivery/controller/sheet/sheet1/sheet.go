package sheet1_controller 






import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
)

type Transaction struct {
	PaidIn     float64
	Balance    float64
	Withdrawal float64
	Agent      string
}

func parseTransaction(row []string) (Transaction, error) {
	var transaction Transaction
	var err error

	// Parse PaidIn
	transaction.PaidIn, err = strconv.ParseFloat(strings.ReplaceAll(row[0], ",", ""), 64)
	if err != nil && row[0] != "" {
		return transaction, fmt.Errorf("failed to parse PaidIn: %v", err)
	}

	// Parse Balance
	transaction.Balance, err = strconv.ParseFloat(strings.ReplaceAll(row[1], ",", ""), 64)
	if err != nil {
		return transaction, fmt.Errorf("failed to parse Balance: %v", err)
	}

	// Parse Withdrawal
	transaction.Withdrawal, err = strconv.ParseFloat(strings.ReplaceAll(row[2], ",", ""), 64)
	if err != nil && row[2] != "" {
		return transaction, fmt.Errorf("failed to parse Withdrawal: %v", err)
	}

	// Set Agent
	transaction.Agent = row[3]

	return transaction, nil
}

// ProcessTransactions processes the first four columns only
func processTransactions(filePath string) error {
	// Open the Excel file
	excelFile, err := excelize.OpenFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to open Excel file: %v", err)
	}

	// Get the first sheet
	sheetName := excelFile.GetSheetName(1)
	rows, err := excelFile.GetRows(sheetName)
	if err != nil {
		return fmt.Errorf("failed to read rows: %v", err)
	}

	// Process each row
	for i, row := range rows {
		if i == 0 { // Skip the header row
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

func main() {
	filePath := "transactions.xlsx" // Replace with the actual file path

	err := processTransactions(filePath)
	if err != nil {
		fmt.Printf("Error processing transactions: %v\n", err)
	}
}

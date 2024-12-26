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


type Sheet1Handler struct {
	userRepo irepo.User
	rootRepo irepo.Root
}

// Sheet1Config holds the configuration for creating a Sheet1Handler.
type Sheet1Config struct {
	UserRepo           irepo.User
	RootRepo 		   irepo.Root
	
}

// Ensure Sheet1Handler implements icmd.IHandler.
var _ icmd.IHandler[*Sheet1Command, *models.Root] = &Sheet1Handler{}

// NewSheet1Handler creates a new instance of Sheet1Handler with the provided configuration.
func NewSheet1Handler(config Sheet1Config) *Sheet1Handler {
	return &Sheet1Handler{
		userRepo: config.UserRepo,
		rootRepo: config.RootRepo,
	}
}

func (s *Sheet1Handler) Handle(cmd *Sheet1Command) (*models.Root, error) {
	fileData := cmd.file
	sheet := cmd.sheet

	excelFile, err := excelize.OpenReader(bytes.NewReader(fileData))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to open Excel file from bytes: %v", err))
	}

	sheetName := excelFile.GetSheetName(sheet)
	rows, err := excelFile.GetRows(sheetName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to read rows: %v", err))
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}

		if len(row) < 4 {
			fmt.Printf("Skipping row %d: not enough columns\n", i+1)
			continue
		}

		limitedRow := row[:4]
		transaction, err := parseTransaction(limitedRow)
		if err != nil {
			fmt.Printf("Error parsing row %d: %v\n", i+1, err)
			continue
		}

		agentCode := strings.Split(transaction.Agent(), " ")[0]

		if transaction.PaidIn() > 0 {
			fmt.Printf("Notify agent %s: Deduct %.2f from PaidIn. Update admin to new balance %.2f\n", agentCode, transaction.PaidIn(), transaction.Balance())
		} else {
			fmt.Printf("Notify agent %s: Deduct %.2f from Withdrawal. Update admin to new balance %.2f\n", agentCode, transaction.Withdrawal(), transaction.Balance())
		}
	}
	return &models.Root{}, nil
}


func parseTransaction(row []string) (models.Transaction, error) {
	var transaction models.Transaction
	var err error

	paidin, err := strconv.ParseFloat(strings.ReplaceAll(row[0], ",", ""), 64)
	if err != nil && row[0] != "" {
		return transaction, fmt.Errorf("failed to parse PaidIn: %v", err)
	}
	transaction.SetPaidIn(paidin)

	balance, err := strconv.ParseFloat(strings.ReplaceAll(row[1], ",", ""), 64)
	if err != nil {
		return transaction, fmt.Errorf("failed to parse Balance: %v", err)
	}
	transaction.SetBalance(balance)


	withdrawal, err := strconv.ParseFloat(strings.ReplaceAll(row[2], ",", ""), 64)
	if err != nil && row[2] != "" {
		return transaction, fmt.Errorf("failed to parse Withdrawal: %v", err)
	}

	transaction.SetWithdrawal(withdrawal)

	transaction.SetAgent( row[3])

	return transaction, nil
}

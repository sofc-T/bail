package models


type Transaction struct {
	file []byte 
	sheet int
	paidIn float64
	balance float64
	withdrawal float64
	agent string


}


type TransactionConfig struct {
	File []byte 
	Sheet int 
	PaidIn float64 
	Balance float64 
	Withdrawal float64 
	Agent string 


}


func (t Transaction) File () []byte {
	return t.file
}

func (t Transaction) Sheet () int {
	return t.sheet
}

func (t *Transaction) PaidIn () float64 {
	return t.paidIn
}

func (t *Transaction) Balance () float64 {
	return t.balance
}

func (t *Transaction) Withdrawal () float64 {
	return t.withdrawal
}

func (t *Transaction) Agent () string {
	return t.agent
}

func NewTransaction (config TransactionConfig) *Transaction {
	return &Transaction {
		file: config.File,
		sheet: config.Sheet,
		paidIn: config.PaidIn,
		balance: config.Balance,
		withdrawal: config.Withdrawal,
		agent: config.Agent,
	}
}

func (t *Transaction) SetFile (file []byte) {
	t.file = file
}

func (t *Transaction) SetSheet (sheet int) {
	t.sheet = sheet
}

func (t *Transaction) SetPaidIn (paidIn float64) {
	t.paidIn = paidIn
}

func (t *Transaction) SetBalance (balance float64) {
	t.balance = balance
}

func (t *Transaction) SetWithdrawal (withdrawal float64) {
	t.withdrawal = withdrawal
}

func (t *Transaction) SetAgent (agent string) {
	t.agent = agent
}


package models

type Root struct {
	id int
	newTransactions int
	balance float64
	
}

func (r Root) NewRoot(balance float64) *Root {
	return &Root{
		id: 1,
		newTransactions:  0,
		balance: balance,
	}
}


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

type RootConfig struct {
	Id int 
	NewTransaction float64
	Balance float64
}


func MapConfig(rootcofig RootConfig) *Root{
	return &Root{
		id:              rootcofig.Id,
		newTransactions: int(rootcofig.NewTransaction),
		balance:         rootcofig.Balance,
	}
}


func (r *Root) SetID(id int) {
	r.id = id
}

func (r *Root) GetID() int {
	return r.id
}

func (r *Root) SetNewTransactions(newTransactions int) {
	r.newTransactions = newTransactions
}

func (r *Root) GetNewTransactions() int {
	return r.newTransactions
}

func (r *Root) SetBalance(balance float64) {
	r.balance = balance
}

func (r *Root) GetBalance() float64 {
	return r.balance
}


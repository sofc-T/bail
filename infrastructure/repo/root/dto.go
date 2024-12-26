package rootrepo

import (
	"bail/domain/models"

)

type adminDTO struct {
	Id  int `bson:"_id"`
	Balance float64 `bson:"balance"`
	NewTransactions int `bson:"new_transactions"`
}

func fromAdmin(root models.Root) adminDTO {
	return adminDTO{
		Id:             root.GetID(),
		Balance:        root.GetBalance(),
		NewTransactions: root.GetNewTransactions(),
	}
}


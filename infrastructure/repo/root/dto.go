package rootrepo

import (
	"bail/domain/models"

	"github.com/google/uuid"
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
func toAdmin(dto adminDTO) models.Root{
	return dto.Id, dto.Balance, dto.NewTransactions
}


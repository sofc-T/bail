package usercontroller

import (
	"bail/domain/models"

	"github.com/google/uuid"
)

type SignUpDto struct {
	Name                string  `json:"name" binding:"required"`
	Email               string  `json:"email" binding:"required"`
	Salary              float64 `json:"salary" `
	Age                 int     `json:"age" `
	Role                string  `json:"role" `
	CoSignerName        string  `json:"co_signer_name" `
	CodeNumber          string  `json:"code_number" `
	CoSignerDocument    []byte  `json:"co_signer_document" `
	EducationalDocument []byte  `json:"educational_document" `
	Password            string  `json:"password" `
}

type UpdateUserDto struct {
	ID                  uuid.UUID `json:"id" `
	Name                string    `json:"name"`
	Email               string    `json:"email"`
	Salary              float64   `json:"salary"`
	Age                 int       `json:"age"`
	Role                string    `json:"role"`
	CoSignerName        string    `json:"co_signer_name"`
	CodeNumber          string    `json:"code_number"`
	CoSignerDocument    []byte    `json:"co_signer_document"`
	EducationalDocument []byte    `json:"educational_document"`
	Password            string    `json:"password"`
}

type resTokenDto struct {
	ID                  uuid.UUID `json:"id"`
	Name                string    `json:"name"`
	Email               string    `json:"email"`
	Salary              float64   `json:"salary"`
	Age                 int       `json:"age"`
	Role                string    `json:"role"`
	CoSignerName        string    `json:"co_signer_name"`
	CodeNumber          string    `json:"code_number"`
	CoSignerDocument    []byte    `json:"co_signer_document"`
	EducationalDocument []byte    `json:"educational_document"`
}

func toUserTokenDto(user *models.User) resTokenDto {
	return resTokenDto{
		ID:                  user.ID(),
		Name:                user.Name(),
		Email:               user.Email(),
		Salary:              user.Salary(),
		Age:                 user.Age(),
		Role:                user.Role(),
		CoSignerName:        user.CoSignerName(),
		CodeNumber:          user.CodeNumber(),
		CoSignerDocument:    user.CoSignerDocument(),
		EducationalDocument: user.EducationalDocument(),
	}
}

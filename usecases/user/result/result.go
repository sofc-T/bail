package result

import (
	"bail/domain/models"

	"github.com/google/uuid"
)

type SignUpResult struct {
	ID                  uuid.UUID
	Name                string
	Email               string
	Salary              float64
	Age                 int
	Role                string
	CoSignerName        string
	CodeNumber          string
	CoSignerDocument    []byte
	EducationalDocument []byte
}

func NewSignUpResult(ID uuid.UUID, name string, email string, salary float64, age int, role string, coSignerName string, codeNumber string, coSignerDocument []byte, educationalDocument []byte) SignUpResult {
	return SignUpResult{
		ID:                  ID,
		Name:                name,
		Email:               email,
		Salary:              salary,
		Age:                 age,
		Role:                role,
		CoSignerName:        coSignerName,
		CodeNumber:          codeNumber,
		CoSignerDocument:    coSignerDocument,
		EducationalDocument: educationalDocument,
	}
}

type LoginInResult struct {
	User  *models.User
	Token string
}

func NewLoginInResult(token string,  user *models.User) LoginInResult {
	return LoginInResult{
		Token: token,
		User:  user,
	}
}

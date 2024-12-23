package usercmd

import (
	"github.com/google/uuid"
)

// SignUpCommand represents the command to register a new user with necessary details.
type UpdateUserCommand struct {
	id                  uuid.UUID
	name                string
	salary              float64
	age                 int
	role                string
	coSignerName        string
	codeNumber          string
	coSignerDocument    []byte 
	educationalDocument []byte
}

// NewSignUpCommand creates a new SignUpCommand instance with the provided user details.
func NewUpdateCommand(id uuid.UUID, name string, email string, salary float64, age int, role string, coSignerName string, codeNumber string, coSignerDocument []byte, educationalDocument []byte) *UpdateUserCommand {
	return &UpdateUserCommand{
		id:                  id,
		name:                name,
		salary:              salary,
		age:                 age,
		role:                role,
		coSignerName:        coSignerName,
		codeNumber:          codeNumber,
		coSignerDocument:    coSignerDocument,
		educationalDocument: educationalDocument,
	}
}

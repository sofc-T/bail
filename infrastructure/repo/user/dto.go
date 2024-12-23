package userrepo

import (
	"time"

	"bail/domain/models"

	"github.com/google/uuid"
)

type userDTO struct {
	Id                  uuid.UUID `bson:"_id"`
	Name                string    `bson:"name"`
	Email               string    `bson:"email"`
	Salary              float64   `bson:"salary"`
	Age                 int       `bson:"age"`
	Role                string    `bson:"role"`
	CoSignerName        string    `bson:"co_signer_name"`
	CodeNumber          string    `bson:"code_number"`
	CoSignerDocument    []byte    `bson:"co_signer_document"`
	EducationalDocument []byte    `bson:"educational_document"`
	CreatedAt           time.Time `bson:"created_at"`
	UpdatedAt           time.Time `bson:"updated_at"`
	password            string    `bson:"password"`
}

func FromUser(u *models.User) *userDTO {
	return &userDTO{
		Id:                  u.ID(),
		Name:                u.Name(),
		Email:               u.Email(),
		Salary:              u.Salary(),
		Age:                 u.Age(),
		Role:                u.Role(),
		CoSignerName:        u.CoSignerName(),
		CodeNumber:          u.CodeNumber(),
		CoSignerDocument:    u.CoSignerDocument(),
		EducationalDocument: u.EducationalDocument(),
		CreatedAt:           u.CreatedAt(),
		UpdatedAt:           u.UpdatedAt(),
		password:            u.Password(),
	}
}

func ToUser(u *userDTO) *models.User {
	user := models.MapUser(
		models.UserConfig{
			ID:                  u.Id,
			Name:                u.Name,
			Email:               u.Email,
			Salary:              u.Salary,
			Age:                 u.Age,
			Role:                u.Role,
			CoSignerName:        u.CoSignerName,
			CodeNumber:          u.CodeNumber,
			CoSignerDocument:    u.CoSignerDocument,
			EducationalDocument: u.EducationalDocument,
		},
	)
	return user
}

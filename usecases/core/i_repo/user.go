package irepo

import (
	"github.com/google/uuid"
	"bail/domain/models"
)

// User defines methods to interact with User storage.
type User interface {
	// Save saves a User.
	Save(user *models.User) error

	// Delete removes a User by ID.
	Delete(id uuid.UUID) error

	//Find User by Email
	FindByEmail(email string) (*models.User, error)

	//Find by Id
	FindById(userid uuid.UUID) (*models.User, error)

	//Get all
	GetAll(int) ([]*models.User, error)

	//add transaction
	AddTransaction(string, float64) (*models.User, error)

	//Find by Code
	FindByCode(code string) (*models.User, error)

}

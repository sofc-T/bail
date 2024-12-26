package irepo

import (
	"bail/domain/models"
)

// User defines methods to interact with User storage.
type Root interface {
	// Save saves a User.
	Save(*models.Root) error

    // adds a transaction
    AddTransaction(float64) error


}

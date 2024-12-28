package irepo

import (
	"bail/domain/models"
)

// User defines methods to interact with User storage.
type SystemLog interface {
	// Save saves a User.
	Save(*models.SystemLog) error

    // adds a transaction
    


}

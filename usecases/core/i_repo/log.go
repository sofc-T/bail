package irepo

import (
	"bail/domain/models"
)

// User defines methods to interact with User storage.
type SystemLog interface {
	// Save saves a User.
	Save(*models.SystemLog) error

    //add a log 
	AddLog(*models.SystemLog) error
    


}

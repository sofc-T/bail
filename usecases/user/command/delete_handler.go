package usercmd

import (
	icmd "bail/usecases/core/cqrs/command"
	irepo "bail/usecases/core/i_repo"

	"github.com/google/uuid"
)

// // DeleteHandler handles user Delete requests.
type DeleteHandler struct {
	repo         irepo.User
}

// // DeleteCommand is a command for logging in an user.

// NewDeleteHandler creates a new instance of DeleteHandler with the provided configuration.
func NewDeleteHandler(userepo irepo.User) *DeleteHandler {
	return &DeleteHandler{
		repo:         userepo,
	}
}

// // Ensure DeleteHandler implements icmd.IHandler
var _ icmd.IHandler[uuid.UUID, error] = &DeleteHandler{}


// // Handle processes the Delete command and returns the error.
func (h *DeleteHandler) Handle(id uuid.UUID) (error, error) {
	err := h.repo.Delete(id)
	if err != nil {
		return nil, err
	}

	return nil, nil
	
}	
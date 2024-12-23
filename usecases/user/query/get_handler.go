package userqry

import (
	"bail/domain/models"

	"github.com/google/uuid"

	iqry "bail/usecases/core/cqrs/query"
	irepo "bail/usecases/core/i_repo"
)

// GetHandler is responsible for handling the Get User query by its ID.
type GetHandler struct {
	repo irepo.User
}

// Ensure Handler implements the IHandler interface
var _ iqry.IHandler[uuid.UUID, *models.User] = &GetHandler{}

// NewGetHandler creates a new instance of Handler with the provided businessOwner repository.
func NewGetHandler(reationRepo irepo.User) *GetHandler {
	return &GetHandler{
		repo: reationRepo,
	}
}

// Handle processes the Get query by its ID and returns the corresponding User.
func (h *GetHandler) Handle(id uuid.UUID) (*models.User, error) {

	var err error
	var User *models.User

	User, err = h.repo.FindById(id)

	if err != nil {
		return nil, err
	}

	return User, nil
}

package userqry

import (
	"bail/domain/models"

	iqry "bail/usecases/core/cqrs/query"
	irepo "bail/usecases/core/i_repo"
)

// GetInfuencersHandler is responsible for handling the Get user query by its ID.
type GetInfuencersHandler struct {
	repo irepo.User
}

// Ensure Handler implements the IHandler interface
var _ iqry.IHandler[int, []*models.User] = &GetInfuencersHandler{}

// NewGetInfuencersHandler creates a new instance of Handler with the provided users repository.
func NewGetusersHandler(userRepo irepo.User) *GetInfuencersHandler {
	return &GetInfuencersHandler{
		repo: userRepo,
	}
}

func (h *GetInfuencersHandler) Handle(page int) ([]*models.User, error) {
	users, err := h.repo.GetAll(page)
	if err != nil {
		return nil, err
	}
	return users, nil
}

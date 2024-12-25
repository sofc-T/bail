package usercmd

import (
	"bail/domain/models"
	irepo "bail/usecases/core/i_repo"
	"errors"

	"github.com/google/uuid"
)

// PromoteUserCommand is a command for promoting an user.
type PromoteUserCommand struct {
	ID uuid.UUID
	Role string

}

// NewPromoteUserCommand creates a new instance of PromoteUserCommand with the provided configuration.
func NewPromoteUserCommand(id uuid.UUID, role string) *PromoteUserCommand {
	return &PromoteUserCommand{
		ID: id,
		Role: role,
	}
}


func NewPromoteHandler(userrepo irepo.User) *PromoteHandler {
	return &PromoteHandler{
		Userrepo: userrepo,
	}
}

// Handle processes the promote command to promote an user.
func (h *PromoteHandler) Handle(command *PromoteUserCommand) (*models.User, error) {
	user, err := h.Userrepo.FindById(command.ID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("User not found")
	}

	if command.Role != "admin" && command.Role != "user" && command.Role != "manager" && command.Role != "HR" {
		return nil, errors.New("Invalid role")
	}

	user.SetRole(command.Role)

	err = h.Userrepo.Save(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
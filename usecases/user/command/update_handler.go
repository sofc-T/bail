package usercmd

import (
	"errors"
	"log"

	icmd "bail/usecases/core/cqrs/command"
	irepo "bail/usecases/core/i_repo"
	"bail/usecases/user/result"
)

type UpdateUserHandler struct {
	userrepo          irepo.User
}

type UpdateUserConfig struct {
	UserRepo          irepo.User
}

var _ icmd.IHandler[*UpdateUserCommand, *result.SignUpResult] = &UpdateUserHandler{}

func NewUpdateUserHandler(config UpdateUserConfig) *UpdateUserHandler {
	return &UpdateUserHandler{
		userrepo:          config.UserRepo,
	}
}

func (h *UpdateUserHandler) Handle(command *UpdateUserCommand) (*result.SignUpResult, error) {
	log.Println("Starting Set user process")

	user, err := h.userrepo.FindById(command.id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("UserNotFound")
	}


	if  command.age  != 0{
		user.SetAge(command.age)
	}

	if "" != command.name {
		user.SetName(command.name)
	}

	if 0 != command.salary {
		user.SetSalary(command.salary)
	}

	if "" != command.role {
		user.SetRole(command.role)
	}

	if "" != command.coSignerName {
		user.SetCoSignerName(command.coSignerName)
	}

	if "" != command.codeNumber {
		user.SetCodeNumber(command.codeNumber)
	}
	
	if nil != command.coSignerDocument {
		user.SetCoSignerDocument(command.coSignerDocument)
	}
	
	
	if nil != command.educationalDocument {
		user.SetEducationalDocument(command.educationalDocument)
	}



	err = h.userrepo.Save(user)
	if err != nil {
		return nil, err
	}

	return &result.SignUpResult{
		ID:                  user.ID(),
		Name:                user.Name(),
		Email:               user.Email(),
		Salary:              user.Salary(),
		Age:                 user.Age(),
		Role:                user.Role(),
		CoSignerName:        user.CoSignerName(),
		CodeNumber:          user.CodeNumber(),
		CoSignerDocument:    user.CoSignerDocument(),
		EducationalDocument: user.EducationalDocument(),
	
	}, nil

}
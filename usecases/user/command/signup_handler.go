package usercmd

import (
	"errors"
	"log"
	// "time"

	// "math/rand"

	ihash "bail/domain/i_hash"
	"bail/domain/models"
	icmd "bail/usecases/core/cqrs/command"
	ijwt "bail/usecases/core/i_jwt"
	irepo "bail/usecases/core/i_repo"
	result "bail/usecases/user/result"

	"go.mongodb.org/mongo-driver/mongo"
)

// SignUpHandler handles user sign-up logic.
type SignUpHandler struct {
	Userrepo           irepo.User
	jwtService         ijwt.Service
	hashService        ihash.Service
}

// SignUpConfig holds the configuration for creating a SignUpHandler.
type SignUpConfig struct {
	UserRepo           irepo.User
	JwtService         ijwt.Service
	HashService        ihash.Service
}

// Ensure SignUpHandler implements icmd.IHandler.
var _ icmd.IHandler[*SignUpCommand, *result.SignUpResult] = &SignUpHandler{}

// NewSignUpHandler creates a new instance of SignUpHandler with the provided configuration.
func NewSignUpHandler(config SignUpConfig) *SignUpHandler {
	return &SignUpHandler{
		Userrepo:           config.UserRepo,
		jwtService:         config.JwtService,
		hashService:        config.HashService,
	}
}

// Handle processes the sign-up command to register a new user.
// It creates a new user, checks for conflicts in Email and email,
// generates a validation link, and sends a sign-up email.
func (h *SignUpHandler) Handle(command *SignUpCommand) (*result.SignUpResult, error) {
	log.Println("Starting sign-up process")

	email := command.email

	res, err := h.Userrepo.FindByEmail(email)
	if err != nil {
		//mongo no document found err
		if err != errors.New("UserNotFound") && err != mongo.ErrNoDocuments {
			log.Printf("Error finding user by Email: %v", err.Error())
			return nil, err
		}
	} else if res != nil {
		log.Printf("Email %s is already taken", command.email)
		return nil, errors.New("Email taken")
	}

	log.Printf("Email %s is available", command.email)

	cfg := models.UserConfig{
		Email:    email,
		Password: command.password,
		Name:    command.name,
		Age:    command.age,
		Salary:    command.salary,
		Role:    command.role,
		CoSignerName:    command.coSignerName,
		CodeNumber:    command.codeNumber,
		CoSignerDocument:    command.coSignerDocument,
		EducationalDocument:    command.educationalDocument,

	}


	user := models.NewUser(cfg)
	log.Println("New being user created")


	// Save the new user
	if err := h.Userrepo.Save(user); err != nil {
		log.Printf("Error saving new user: %v", err)
		return nil, err
	}
	log.Printf("New user %s saved successfully", user.Email())

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

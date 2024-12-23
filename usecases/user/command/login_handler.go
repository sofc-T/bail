package usercmd

// import (
// 	"errors"
// 	"log"
// 	ihash "bail/domain/i_hash"
// 	icmd "bail/usecases/core/cqrs/command"
// 	ijwt "bail/usecases/core/i_jwt"
// 	irepo "bail/usecases/core/i_repo"
// 	result "bail/usecases/user/result"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// // LoginHandler handles user login requests.
// type LoginHandler struct {
// 	repo         irepo.User
// 	jwtService   ijwt.Service
// 	hashService  ihash.Service
// }

// // LoginConfig holds the configuration for creating a LoginHandler.
// type LoginConfig struct {
// 	UserRepo     irepo.User
// 	JwtService   ijwt.Service
// 	HashService  ihash.Service
// }

// // LoginCommand is a command for logging in an user.

// // NewLoginHandler creates a new instance of LoginHandler with the provided configuration.
// func NewLoginHandler(config LoginConfig) *LoginHandler {
// 	return &LoginHandler{
// 		repo:         config.UserRepo,
// 		jwtService:   config.JwtService,
// 		hashService:  config.HashService,
// 	}
// }

// // Ensure LoginHandler implements icmd.IHandler
// var _ icmd.IHandler[*LoginCommand, *result.LoginInResult] = &LoginHandler{}

// // Handle processes the login command and returns the login result with tokens.
// func (h *LoginHandler) Handle(command *LoginCommand) (*result.LoginInResult, error) {
// 	// Find user by email
// 	log.Printf("Finding user by email: %s", command)
// 	user, err := h.repo.FindByEmail(command.email)

// 	// check for not found error if not not found error return the error
// 	if err != nil && err != mongo.ErrNoDocuments {
// 		return nil, err
// 	}
// 	if user == nil {
// 		return nil, errors.New("user not found")

// 	}

// 	// Verify password
// 	log.Println(user, "found this user", user.PasswordHash(), "now checking password")
// 	ok, err := h.hashService.Match(user.PasswordHash(), command.password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if !ok {
// 		return nil, errors.New("password is incorrect")
// 	}

// 	// Mark user as active
	



// 	user.SetIsActive(true)
// 	err = h.repo.Save(user)
// 	if err != nil {
// 		return nil, errors.New("internal server Error")
// 	}

// 	// Generate tokens
// 	token, err := h.jwtService.Generate(user, ijwt.Access)
// 	if err != nil {
// 		return nil, err
// 	}

// 	refreshToken, err := h.jwtService.Generate(user, ijwt.Refresh)
// 	if err != nil {
// 		return nil, err
// 	}



// 	// Return login result
// 	result := result.NewLoginInResult(token, refreshToken, user)
// 	return &result, nil
// }


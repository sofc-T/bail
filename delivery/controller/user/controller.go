package usercontroller

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	

	// "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	// "github.com/google/uuid"
	basecontroller "bail/delivery/base"
	"bail/domain/models"
	icmd "bail/usecases/core/cqrs/command"

	// passwordreset "bail/usecases/password_reset"
	// resultValidate "bail/usecases/password_reset/result"
	usercmd "bail/usecases/user/command"
	"bail/usecases/user/result"
)

// UserController handles user-related HTTP requests.
type UserController struct {
	basecontroller.BaseHandler
	loginUserHandler     icmd.IHandler[*usercmd.LoginCommand, *result.LoginInResult]
	signupUserHandler    icmd.IHandler[*usercmd.SignUpCommand, *result.SignUpResult]
	updateUserHandler          icmd.IHandler[*usercmd.UpdateUserCommand, *result.SignUpResult]
	getEmployeeHandler       icmd.IHandler[int, []*models.User]
	getUserHandler			 icmd.IHandler[uuid.UUID, *models.User]
	deleteEmployeeHandler    icmd.IHandler[uuid.UUID, error]
	promoteUserHandler       icmd.IHandler[*usercmd.PromoteUserCommand, *models.User]

}

// Config holds the configuration for creating a new UserController.
type Config struct {
	SignupUserHandler    icmd.IHandler[*usercmd.SignUpCommand, *result.SignUpResult]
	UpdateUserHandler          icmd.IHandler[*usercmd.UpdateUserCommand, *result.SignUpResult]
	GetEmployeeHandler       icmd.IHandler[int, []*models.User]
	GetUserHandler			 icmd.IHandler[uuid.UUID, *models.User]
	DeleteEmployeeHandler    icmd.IHandler[uuid.UUID, error]
	LoginUserHandler         icmd.IHandler[*usercmd.LoginCommand, *result.LoginInResult]
	PromoteUserHandler       icmd.IHandler[*usercmd.PromoteUserCommand, *models.User]
	
}

// New creates a new UserController with the given CQRS handlers.
func New(config Config) *UserController {
	return &UserController{
		signupUserHandler:    config.SignupUserHandler,
		updateUserHandler:          config.UpdateUserHandler,
		getEmployeeHandler:       config.GetEmployeeHandler,
		getUserHandler:			 config.GetUserHandler,
		deleteEmployeeHandler:    config.DeleteEmployeeHandler,
		loginUserHandler:         config.LoginUserHandler,
	}
}

func (u UserController) RegisterPrivilegedAdmin(router *gin.RouterGroup) {
	router = router.Group("/auth")
	router.POST("/promote", u.promote)
}

func (u UserController) RegisterProtected(router *gin.RouterGroup) {

}

func (u UserController) RegisterPrivilegedHR(router *gin.RouterGroup) {
	
	
}

func (u UserController) RegisterPrivilegedManager(router *gin.RouterGroup) {
	
	
}

func (u UserController) RegisterPublic(router *gin.RouterGroup) {
	router = router.Group("/auth")
	router.GET("/employee/list/:page", u.getAllEmployees)
	router.GET("/employee/:id", u.getAllEmployees)
	router.POST("/signup", u.signUp)
	router.DELETE("/delete/:id", u.deleteEmployee)
	router.PATCH("/update", u.update)
	// router.POST("/login", u.login)
}

func (u *UserController) signUp(ctx *gin.Context) {
	var user SignUpDto
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		log.Println("User input could not be bound -- UserController", err)
		return
	}

	user.Email = strings.ToLower(user.Email)
	
	command := usercmd.NewSignUpCommand( user.Name, user.Email, user.Salary, user.Age, user.Role, user.CoSignerName, user.CodeNumber, user.CoSignerDocument, user.EducationalDocument, user.Password)

	_, err := u.signupUserHandler.Handle(command)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		log.Println("User use case invalidated data -- UserController")
		return
	}
	log.Println("User signed up successfully -- UserController")
	ctx.JSON(http.StatusCreated, "Signed Up successfully")

}



func (u UserController) update(ctx *gin.Context) {
	var user UpdateUserDto
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		log.Println("User input could not be bound -- UserController")
		return
	}

	user.Email = strings.ToLower(user.Email)
	command := usercmd.NewUpdateCommand(user.ID, user.Name, user.Email, user.Salary, user.Age, user.Role, user.CoSignerName, user.CodeNumber, user.CoSignerDocument, user.EducationalDocument)

	res, err := u.updateUserHandler.Handle(command)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		log.Println("User use case invalidated data -- UserController")
		return
	}

	log.Println("User updated successfully -- UserController")
	ctx.JSON(http.StatusCreated, res)
}



func (u UserController) getAllEmployees(ctx *gin.Context) {

	page, err := strconv.Atoi(ctx.Param("page"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		log.Println("User input could not be bound -- UserController")
		return
	}

	res, err := u.getEmployeeHandler.Handle(page)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		log.Println("User use case invalidated data -- UserController")
		return
	}

	var users []resTokenDto 
	for _, user := range res {
		if res == nil{
			continue
			
		}
		users = append(users, toUserTokenDto(user))
	}

	log.Println("users fetched successfully -- UserController")
	ctx.JSON(http.StatusOK, users)

}


func (u UserController) getUser(ctx *gin.Context) {
	idstr := ctx.Param("id")
	id, err := uuid.Parse(idstr)
	log.Println(idstr, id, err)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		log.Println("User input could not be bound -- UserController")
		return
	}

	res, err := u.getUserHandler.Handle(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		log.Println("User use case invalidated data -- UserController")
		return
	}

	if res == nil{
		ctx.JSON(http.StatusNotFound, "User not found")
		return 
	}
	userres := toUserTokenDto(res)

	log.Println("User fetched successfully -- UserController")
	ctx.JSON(http.StatusOK, userres)

}

func (u UserController) deleteEmployee(ctx *gin.Context) {
	idstr := ctx.Param("id")
	id, err := uuid.Parse(idstr)
	log.Println(idstr, id, err)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		log.Println("User input could not be bound -- UserController")
		return
	}

	_, err = u.deleteEmployeeHandler.Handle(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		log.Println("User use case invalidated data -- UserController")
		return
	}

	log.Println("User deleted successfully -- UserController")
	ctx.JSON(http.StatusOK, "User deleted successfully")

}


func (u UserController) login(ctx *gin.Context) {
	var user LoginDto
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		log.Println("User input could not be bound -- UserController")
		return
	}

	user.Email = strings.ToLower(user.Email)
	
	command := usercmd.NewLoginCommand(user.Email, user.Password)

	res, err := u.loginUserHandler.Handle(command)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		log.Println("User use case invalidated data -- UserController")
		return
	}

	log.Println("User logged in successfully -- UserController")
	ctx.JSON(http.StatusCreated, res)

}

func (u UserController) promote(ctx *gin.Context) {
	var promote promoteDto
	if err := ctx.BindJSON(&promote); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		log.Println("User input could not be bound -- UserController")
		return
	}



	command := usercmd.NewPromoteUserCommand(promote.ID, promote.Role)

	res, err := u.promoteUserHandler.Handle(command)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		log.Println("User use case invalidated data -- UserController")
		return
	}

	result := toUserDto(res)
	log.Println("User promoted successfully -- UserController")
	ctx.JSON(http.StatusCreated, result)

}

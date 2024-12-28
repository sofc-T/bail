package sheet3_controller

import (
	basecontroller "bail/delivery/base"
	"bail/domain/models"
	"net/http"

	icmd "bail/usecases/core/cqrs/command"
	transaction_cmd "bail/usecases/sheet3"

	"github.com/gin-gonic/gin"
	
)



type sheet3_controller struct {
	basecontroller.BaseHandler	
	parseconntroller icmd.IHandler[ transaction_cmd.sheet3Command , models.Root]
}

type Config struct{
	ParseHandler icmd.IHandler[ transaction_cmd.sheet3Command , models.Root]
}

func New(config Config) *sheet3_controller {
	return &sheet3_controller{
		parseconntroller: config.ParseHandler,
	}
}

func (u sheet3_controller) RegisterPrivilegedAdmin(router *gin.RouterGroup) {

}

func (u sheet3_controller) RegisterProtected(router *gin.RouterGroup) {
	
}

func (u sheet3_controller) RegisterPrivilegedHR(router *gin.RouterGroup) {
	
	
}

func (u sheet3_controller) RegisterPrivilegedManager(router *gin.RouterGroup) {
	
	
}

func (u sheet3_controller) RegisterPublic(router *gin.RouterGroup) {
	router = router.Group("/parse1")
	router.POST("/promote", u.sheet3)
}



func (u sheet3_controller) sheet3(c *gin.Context) {
	var  input sheet3Dto

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := transaction_cmd.Newsheet3Command(input.File, input.Sheet)
	result, err := u.parseconntroller.Handle(*cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	c.JSON(http.StatusOK, result)
}

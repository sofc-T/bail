package sheet1_controller

import (
	basecontroller "bail/delivery/base"
	"bail/domain/models"
	"net/http"

	icmd "bail/usecases/core/cqrs/command"
	transaction_cmd "bail/usecases/sheet1"

	"github.com/gin-gonic/gin"
	
)



type sheet1_controller struct {
	basecontroller.BaseHandler	
	parseconntroller icmd.IHandler[ transaction_cmd.Sheet1Command , models.Root]
}

type Config struct{
	ParseHandler icmd.IHandler[ transaction_cmd.Sheet1Command , models.Root]
}

func New(config Config) *sheet1_controller {
	return &sheet1_controller{
		parseconntroller: config.ParseHandler,
	}
}

func (u sheet1_controller) RegisterPrivilegedAdmin(router *gin.RouterGroup) {
}

func (u sheet1_controller) RegisterProtected(router *gin.RouterGroup) {
	router = router.Group("/parse1")
	router.POST("/promote", u.sheet1)
}

func (u sheet1_controller) RegisterPrivilegedHR(router *gin.RouterGroup) {
	
	
}

func (u sheet1_controller) RegisterPrivilegedManager(router *gin.RouterGroup) {
	
	
}

func (u sheet1_controller) RegisterPublic(router *gin.RouterGroup) {

}



func (u sheet1_controller) sheet1(c *gin.Context) {
	var  input Sheet1Dto

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := transaction_cmd.NewSheet1Command(input.File, input.Sheet)
	result, err := u.parseconntroller.Handle(*cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	c.JSON(http.StatusOK, result)
}

package sheet2_controller

import (
	basecontroller "bail/delivery/base"
	"bail/domain/models"
	"net/http"

	icmd "bail/usecases/core/cqrs/command"
	transaction_cmd "bail/usecases/sheet2"

	"github.com/gin-gonic/gin"
	
)



type sheet2_controller struct {
	basecontroller.BaseHandler	
	parseconntroller icmd.IHandler[ transaction_cmd.Sheet2Command , models.Root]
}

type Config struct{
	ParseHandler icmd.IHandler[ transaction_cmd.Sheet2Command , models.Root]
}

func New(config Config) *sheet2_controller {
	return &sheet2_controller{
		parseconntroller: config.ParseHandler,
	}
}

func (u sheet2_controller) RegisterPrivilegedAdmin(router *gin.RouterGroup) {

}

func (u sheet2_controller) RegisterProtected(router *gin.RouterGroup) {
	
}

func (u sheet2_controller) RegisterPrivilegedHR(router *gin.RouterGroup) {
	
	
}

func (u sheet2_controller) RegisterPrivilegedManager(router *gin.RouterGroup) {
	
	
}

func (u sheet2_controller) RegisterPublic(router *gin.RouterGroup) {
	router = router.Group("/parse2")
	router.POST("", u.sheet2)
}



func (u sheet2_controller) sheet2(c *gin.Context) {
	var  input Sheet2Dto

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := transaction_cmd.NewSheet2Command(input.File, input.Sheet)
	result, err := u.parseconntroller.Handle(*cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	c.JSON(http.StatusOK, result)
}

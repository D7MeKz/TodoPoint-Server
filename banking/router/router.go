package router

import (
	"github.com/gin-gonic/gin"
	"todopoint/banking/controller"
)

func NewRouter(controller *controller.BankAccountController) *gin.Engine {
	engine := gin.Default()
	router := engine.Group("/banking")
	{
		router.POST("/account/register", controller.RegisterAccount)
	}
	return engine
}

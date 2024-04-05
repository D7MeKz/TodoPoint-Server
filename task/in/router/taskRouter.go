package router

import (
	"github.com/gin-gonic/gin"
	"todopoint/common/netutils/middleware"
	"todopoint/task/in/router/controller"
)

func NewTaskRouter(controller *controller.TaskController) *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.ErrorHandler())

	router := engine.Group("/tasks")
	{
		router.POST("/create", controller.CreateTask)

	}

	return engine
}

package router

import (
	"github.com/gin-gonic/gin"
	"todopoint/task/in/web"
)

func NewTaskRouter(controller *web.TaskController) *gin.Engine {
	engine := gin.Default()
	router := engine.Group("/tasks")
	{
		router.POST("/create", controller.CreateTask)
		router.GET("/:userId", controller.GetTask)
	}

	return engine
}

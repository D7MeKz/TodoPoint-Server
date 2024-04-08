package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"todopoint/common/netutils/middleware"
	_ "todopoint/task/docs"
	"todopoint/task/in/router/controller"
)

// NewTaskRouter
// @title Task API
// @Version 1.0
// @host localhost:3001
func NewTaskRouter(controller *controller.TaskController) *gin.Engine {
	engine := gin.Default()

	engine.Use(middleware.ErrorHandler())
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router := engine.Group("/tasks")
	{
		router.POST("/create", controller.CreateTask)
	}

	return engine
}

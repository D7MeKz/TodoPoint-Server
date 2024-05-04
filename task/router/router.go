package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"todopoint/common2/netutils/middleware"
	_ "todopoint/task/docs"
)

// NewTaskRouter
// @title Task API
// @Version 1.0
// @host localhost:3001
func NewTaskRouter(controller *TaskController) *gin.Engine {
	engine := gin.Default()

	engine.Use(middleware.ErrorHandler())
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router := engine.Group("/tasks")
	{
		router.POST("/create", controller.CreateTask)
		router.GET("", controller.GetList)
		router.GET("/today", controller.GetToday)

	}

	// SubTask Router
	subRouter := engine.Group("/subtasks")
	{
		subRouter.POST("/create", controller.AddSubtask)
		subRouter.GET("/:subtask_id", controller.CheckSubtask)
	}
	return engine
}

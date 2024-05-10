package router

import (
	"github.com/gin-gonic/gin"
	"todopoint/d7modules/server/middleware"
)

// NewTaskRouter
// @title Task API
// @Version 1.0
// @host localhost:3002
func NewTaskRouter(controller *TaskController) *gin.Engine {
	engine := gin.Default()

	engine.Use(middleware.ErrorHandler())
	engine.GET("/swagger/*any")
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

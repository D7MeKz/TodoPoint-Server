package api

import (
	"github.com/gin-gonic/gin"
	"modules/common/server/middleware"
	"todopoint/user/api/controller"
)

// NewUserRouter
// @title User API
// @Version 1.0
// @host localhost:3000
func NewUserRouter(controller *controller.UserController) *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.ErrorHandler())
	engine.Use(middleware.TokenAuthMiddleware())
	router := engine.Group("/user")
	{
		router.POST("/profile", controller.Create)
		router.GET("/me", controller.Me)
	}
	return engine
}

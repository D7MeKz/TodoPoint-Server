package api

import (
	"github.com/gin-gonic/gin"
	"modules/common/server/middleware"
	"todopoint/auth/api/controller"
)

// NewAuthRouter
// @title Auth API
// @Version 1.0
// @host localhost:3001
func NewAuthRouter(controller *controller.AuthController) *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.ErrorHandler())
	router := engine.Group("/auth")
	{
		router.POST("/login", controller.Login)
		router.POST("/register", controller.Register)
		router.GET("/token", controller.Issue)
		router.GET("/valid", controller.Valid)
	}
	return engine
}

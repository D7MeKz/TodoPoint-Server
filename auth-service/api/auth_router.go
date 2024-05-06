package api

import (
	"github.com/gin-gonic/gin"
	"todopoint/auth/api/controller"
	"todopoint/common/server/middleware"
)

// NewAuthRouter
// @title Auth API
// @Version 1.0
// @host localhost:3001
func NewAuthRouter(controller *controller.AuthController) *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.ErrorHandler())
	router := engine.Group("/auth-service")
	{
		router.GET("/login", controller.Login)
		router.POST("/register", controller.Register)
		router.GET("/token", controller.Issue)
		router.GET("/valid", controller.Valid)
	}
	return engine
}

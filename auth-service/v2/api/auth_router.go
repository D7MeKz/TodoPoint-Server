package api

import (
	"github.com/gin-gonic/gin"
	"modules/v2/common/httputils"
)

// NewAuthRouter
// @title Auth API
// @Version 1.0
// @host localhost:3001
func NewAuthRouter(controller *AuthController) *gin.Engine {
	engine := gin.Default()
	engine.Use(httputils.ErrorMiddleware())
	router := engine.Group("/auth")
	{
		router.POST("/login", controller.Login)
		router.POST("/register", controller.Register)
		router.GET("/token", controller.Issue)
		router.GET("/valid", controller.Valid)
	}
	return engine
}

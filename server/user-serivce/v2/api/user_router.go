package api

import (
	"github.com/gin-gonic/gin"
	"modules/v2/common/httputils"
	"modules/v2/common/security/d7jwt"
)

// NewUserRouter
// @title User API
// @Version 1.0
// @host localhost:3000
func NewUserRouter(controller *UserController) *gin.Engine {
	engine := gin.Default()
	engine.Use(httputils.ErrorMiddleware())
	engine.Use(d7jwt.TokenAuthMiddleware())
	router := engine.Group("/users")
	{
		router.GET("/me", controller.GetMe)
		router.POST("/profile", controller.UpdateProfile)
	}
	return engine
}

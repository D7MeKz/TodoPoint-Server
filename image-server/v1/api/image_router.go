package api

import (
	"github.com/gin-gonic/gin"
	"modules/v2/common/httputils"
	"modules/v2/common/security/d7jwt"
)

// NewImageRouter
// @title Image API
// @Version 1.0
// @host localhost:3003
func NewImageRouter(controller *ImageController) *gin.Engine {
	engine := gin.Default()
	engine.MaxMultipartMemory = 8 << 20 // 8 MiB
	engine.Use(httputils.ErrorMiddleware())
	engine.Use(d7jwt.TokenAuthMiddleware())

	router := engine.Group("/image")
	{
		router.POST("/upload", controller.Upload)
		// router.GET("/download", controller.Download)
	}

	return engine
}

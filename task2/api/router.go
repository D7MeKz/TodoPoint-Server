package api

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Create(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	GetMany(ctx *gin.Context)
}

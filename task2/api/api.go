package api

import (
	"github.com/gin-gonic/gin"
	"todopoint/common/errorutils"
)

type Service interface {
	Create(ctx *gin.Context, data interface{}) (string, *errorutils.NetError)
	Delete(ctx *gin.Context, id int) *errorutils.NetError
	Modify(ctx *gin.Context, data interface{}) *errorutils.NetError
	IsExist(ctx *gin.Context, id int) (bool, *errorutils.NetError)
	GetOne(ctx *gin.Context, id int) (interface{}, *errorutils.NetError)
	GetMany(ctx *gin.Context, count int) ([]interface{}, *errorutils.NetError)
}

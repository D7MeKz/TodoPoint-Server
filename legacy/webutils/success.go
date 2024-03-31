package webutils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type successWithData struct {
	Code int         `json:"codes"`
	Data interface{} `json:"data"`
}

type success struct {
	Code int `json:"codes"`
}

func SuccessWith(ctx *gin.Context, data any) {
	res := successWithData{
		Code: http.StatusOK,
		Data: data,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.AbortWithStatusJSON(res.Code, res)
}
func Success(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	ctx.AbortWithStatusJSON(http.StatusOK, success{http.StatusOK})
}

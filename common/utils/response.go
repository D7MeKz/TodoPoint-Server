package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type errorResponse struct {
	Code   int    `json:"code"`
	Status bool   `json:"status"`
	Error  string `json:"error"`
}

type validResponse struct {
	Code   int         `json:"code"`
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

func Success(ctx *gin.Context, data any) {
	res := validResponse{
		Code:   http.StatusOK,
		Status: true,
		Data:   data,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(res.Code, res)
}
func Error(ctx *gin.Context, code int, msg error) {
	res := errorResponse{
		Code:   code,
		Status: false,
		Error:  msg.Error(),
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(res.Code, res)
}

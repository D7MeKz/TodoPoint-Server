package netutils

import (
	"github.com/gin-gonic/gin"
	"todopoint/common/netutils/codes"
)

type SuccessResponse struct {
	Code    codes.WebCode `json:"codes"`
	Message string        `json:"message"`
}

func NewSuccessResponse(code codes.WebCode) *SuccessResponse {
	return &SuccessResponse{Code: code, Message: "Success"}
}

type SuccessResponseWith struct {
	Code    codes.WebCode `json:"codes"`
	Message string        `json:"message"`
	Data    any           `json:"data"`
}

func NewSuccessResponseWith(data any) *SuccessResponseWith {
	return &SuccessResponseWith{Code: 0, Message: "Success", Data: data}
}

func SuccessWith(ctx *gin.Context, res SuccessResponseWith) {
	status := codes.GetStatus(res.Code)
	ctx.Header("Content-Type", "application/json")
	ctx.AbortWithStatusJSON(status, res)
}
func Success(ctx *gin.Context, res SuccessResponse) {
	status := codes.GetStatus(res.Code)
	ctx.Header("Content-Type", "application/json")
	ctx.AbortWithStatusJSON(status, res)
}

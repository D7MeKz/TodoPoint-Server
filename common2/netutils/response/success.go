package response

import (
	"github.com/gin-gonic/gin"
	"todopoint/common/errorutils/codes"
)

type SuccessResponse struct {
	Code    codes.WebCode `json:"code"`
	Message string        `json:"message"`
}

func NewSuccessResponse(code codes.WebCode) *SuccessResponse {
	return &SuccessResponse{Code: code, Message: "Success"}
}

type SuccessResponseWith struct {
	Code    codes.WebCode `json:"code"`
	Message string        `json:"message"`
	Data    any           `json:"data"`
}

func NewSuccessResponseWith(code codes.WebCode, data any) *SuccessResponseWith {
	return &SuccessResponseWith{Code: code, Message: "Success", Data: data}
}

func SuccessWith(ctx *gin.Context, code codes.WebCode, data any) {
	status := codes.GetStatus(code)
	res := NewSuccessResponseWith(code, data)
	ctx.Header("Content-Type", "application/json")
	ctx.AbortWithStatusJSON(status, res)
}
func Success(ctx *gin.Context, code codes.WebCode) {
	status := codes.GetStatus(code)
	ctx.Header("Content-Type", "application/json")

	res := NewSuccessResponse(code)
	ctx.JSON(status, res)
}

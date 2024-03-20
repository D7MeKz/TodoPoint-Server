package webutils

import "github.com/gin-gonic/gin"

//	type errorResponse struct {
//		Code    int    `json:"code"`
//		Error   string `json:"error"`
//		Message string `json:"message"`
//	}
//
//	func InvalidDataError(ctx *gin.Context, err error) {
//		res := errorResponse{Code: http.StatusUnauthorized, Error: err.Error()}
//		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
//	}
//
//	func InternalDBError(ctx *gin.Context, err error) {
//		res := errorResponse{Code: http.StatusInternalServerError, Error: err.Error()}
//		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
//	}
//
//	func BadRequestError(ctx *gin.Context, err error) {
//		res := errorResponse{Code: http.StatusBadRequest, Error: err.Error()}
//		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
//	}
type ErrorType struct {
	Code int
}

func ErrorFunc(ctx *gin.Context, err error, errType ErrorType) {
	res := getErrorMsg(errType.Code)
	// Todo Logging

	ctx.AbortWithStatusJSON(res.StatusCode, res)
}

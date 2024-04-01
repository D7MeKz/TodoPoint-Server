package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"todopoint/common/netutils/codes"
	"todopoint/common/netutils/errorutils"
	"todopoint/common/netutils/response"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		// JSON이 두번 쓰이는 것을 대비해서 Body 확인
		isBodyWritten := ctx.Writer.Written()
		err := ctx.Errors.Last()

		if err != nil {
			var netErr *errorutils.NetError
			if errors.As(err, &netErr) {
				convertedError, ok := errorutils.Convert(err)

				code := convertedError.GetCode()
				statusCode := codes.GetStatus(code)
				res := response.NewErrorResponse(code)

				// Abort with the appropriate status code and response
				if !isBodyWritten {
					ctx.AbortWithStatusJSON(int(statusCode), res)
				}
			} else {
				res := response.NewErrorResponse(codes.GlobalInternalServerError)
				if !isBodyWritten {
					ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
				}
			}
			//convertedError := err.Unwrap()
			//ok := errorutils.IsNetError(convertedError)
			//
			//logrus.Warn("Warning!")
			//if !ok {
			//	res := response.NewErrorResponse(codes.GlobalInternalServerError)
			//	if !isBodyWritten {
			//		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
			//	}
			//}
			//// if NetError
			//code := convertedError.GetCode()
			//statusCode := codes.GetStatus(code)
			//res := response.NewErrorResponse(code)
			//
			//if !isBodyWritten {
			//	ctx.AbortWithStatusJSON(int(statusCode), res)
			//}
		}
	}

}

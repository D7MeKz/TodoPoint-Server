package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todopoint/common/utils/errorutils"
	"todopoint/common/utils/netutils"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		// JSON이 두번 쓰이는 것을 대비해서 Body 확인
		isBodyWritten := ctx.Writer.Written()
		err := ctx.Errors.Last()
		if err != nil {
			convertedError, ok := errorutils.ConvertNetError(err)
			if !ok {
				res := netutils.NewErrorResponse(errorutils.GlobalInternalServerError)
				if !isBodyWritten {
					ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
				}
			}
			// if NetError
			statusCode := convertedError.GetStatus()
			code := convertedError.GetCode()
			res := netutils.NewErrorResponse(code)
			if !isBodyWritten {
				ctx.AbortWithStatusJSON(int(statusCode), res)
			}
		}
	}

}

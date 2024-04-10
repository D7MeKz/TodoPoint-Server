package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"todopoint/common/errorutils"
	"todopoint/common/errorutils/codes"
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
				code := netErr.GetCode()
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

		}
	}
}

func SetHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0, max-age=0")
	c.Header("Last-Modified", time.Now().String())
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "-1")
	c.Next()
}

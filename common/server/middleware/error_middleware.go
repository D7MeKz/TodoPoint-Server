package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"todopoint/common/d7errors"
	"todopoint/common/d7errors/codes"
	"todopoint/common/server/httpdata/response"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		// JSON이 두번 쓰이는 것을 대비해서 Body 확인
		isBodyWritten := ctx.Writer.Written()
		err := ctx.Errors.Last()

		if err != nil {
			var netErr *d7errors.NetError
			// Get data from NetError
			if errors.As(err, &netErr) {
				code := netErr.GetCode()
				logrus.Errorf("ErrorCode: %d", code)
				statusCode := codes.GetStatus(code)
				res := response.NewErrorResponse(code)

				// Abort with the appropriate status code and domain
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

func SetHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0, max-age=0")
		c.Header("Last-Modified", time.Now().String())
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "-1")
		c.Next()
	}

}

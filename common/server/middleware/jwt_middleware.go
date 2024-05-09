package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"todopoint/common/security/d7jwt"
	"todopoint/common/server/httpdata/d7errors/codes"
	"todopoint/common/server/httpdata/d7success"
)

const (
	BearerSchema string = "BEARER "
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" {
			logrus.Error("Authorization value is Empty")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := fromHeader(header)
		if err != nil {
			logrus.Errorf("Invalid Authorization value: %v", err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Validate Token
		ok, err := d7jwt.IsExpired(token)
		if err != nil {
			logrus.Errorf("Invalid Token : %v", err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// if token expired
		if !ok {
			logrus.Error("Token expired")
			res := d7success.NewErrorResponse(codes.TokenExpired)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}
	}
}

func fromHeader(header string) (string, error) {
	bearerLen := len(BearerSchema)
	if strings.ToUpper(header[0:bearerLen]) == BearerSchema {
		return header[bearerLen:], nil
	}
	return "", errors.New("invalid Authorization value")
}

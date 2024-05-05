package httputils

import (
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
	"todopoint/common/server/httpdata/auth"
)

func getAuthorization(ctx *gin.Context) (string, error) {
	auth := ctx.GetHeader("Authorization")
	if auth == "" {
		return "", errors.New("authorization value is Empty")
	}
	return auth, nil
}

func splitBasic(token string) (string, error) {
	if len(token) < 6 {
		return "", errors.New("invalid Authorization value")
	}
	return token[6:], nil
}

func splitBearer(token string) (string, error) {
	if len(token) < 7 {
		return "", errors.New("invalid Authorization value")
	}
	return token[7:], nil
}

func convert(token string) (*auth.Credential, error) {
	// Decode
	b64data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, errors.New("authorization Header:Base64 Decoding error")
	}

	// split :
	cred := strings.Split(string(b64data), ":")
	return &auth.Credential{
		Email:    cred[0],
		Password: cred[1],
	}, nil
}

func GetCredential(ctx *gin.Context) (*auth.Credential, error) {
	auth, err := getAuthorization(ctx)
	if err != nil {
		return nil, err
	}

	token, err := splitBasic(auth)
	if err != nil {
		return nil, err
	}

	cred, err := convert(token)
	if err != nil {
		return nil, err
	}
	return cred, nil
}

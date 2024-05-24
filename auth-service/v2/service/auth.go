package service

import (
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"modules/v2/common/security/d7jwt"
	"strings"
	"todopoint/auth/v2/data"
)

func GenerateTokenPair(uid int) (*data.TokenPair, error) {
	claim := d7jwt.NewTokenClaims(uid)
	access, err := claim.Generate()
	if err != nil {
		logrus.Errorf("Token creation Error : %v", err)
		return nil, err
	}
	logrus.Debug("Success : Access Token generation")

	// Create Access, Refresh Token
	refresh := uuid.NewString()
	return &data.TokenPair{Access: access, Refresh: refresh}, nil
}

// GetCredential extracts token from header and decode it.
func GetCredential(ctx *gin.Context) (*data.Credential, error) {
	// Extract Basic token
	token, err := d7jwt.GetBasic(ctx)
	if err != nil {
		return nil, err

	}
	// Decode base64 from token
	cred, err := convert(token)
	if err != nil {
		return nil, err
	}
	return cred, nil
}

type UserId struct {
	Id int `json:"user_id"`
}

// extractIdFrom extracts user id from Authorization header.
func extractIdFrom(ctx *gin.Context) (*UserId, error) {
	token, tokenErr := d7jwt.GetBearerToken(ctx)
	if tokenErr != nil {
		return nil, tokenErr
	}

	uid, validErr := d7jwt.GetIdFrom(token)
	if validErr != nil {
		return nil, validErr
	}
	return &UserId{uid}, nil
}

// convert is a function that convert base64 encoded token to Credential struct.
func convert(token string) (*data.Credential, error) {
	// Decode
	print(token)
	b64data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, errors.New("authorization Header:Base64 Decoding error")
	}

	cred := strings.Split(string(b64data), ":")
	return &data.Credential{
		Email:    cred[0],
		Password: cred[1],
	}, nil
}

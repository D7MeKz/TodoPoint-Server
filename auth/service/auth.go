package service

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"todopoint/auth/data/domain"
	"todopoint/common/server/middleware"
)

func GenerateTokenPair(uid int) (*domain.TokenPair, error) {
	claim := middleware.NewTokenClaims(uid)
	access, err := claim.Generate()
	if err != nil {
		logrus.Errorf("Token creation Error : %v", err)
		return nil, err
	}
	logrus.Debug("Success : Access Token generation")

	// Create Access, Refresh Token
	refresh := uuid.NewString()
	return &domain.TokenPair{Access: access, Refresh: refresh}, nil
}

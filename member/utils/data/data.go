package data

import (
	"github.com/google/uuid"
	"time"
)

type TokenDetail struct {
	AccessToken  string
	RefreshToken string
	RtExpires    int64
}

func NewTokenDetail(accessToken string) *TokenDetail {
	return &TokenDetail{
		AccessToken:  accessToken,
		RefreshToken: uuid.NewString(),
		RtExpires:    time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
}

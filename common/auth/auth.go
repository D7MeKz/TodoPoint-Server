package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var (
	secret = "todopoint123"
)

type TokenClaims struct {
	TokenUUID string   `json:"tid"`
	UserID    int      `json:"user_id"`
	Email     string   `json:"email"`
	Role      []string `json:"role"`
	jwt.MapClaims
}

func NewTokenClaims(uid int) *TokenClaims {
	claim := TokenClaims{
		TokenUUID: uuid.NewString(),
		UserID:    uid,

		MapClaims: jwt.MapClaims{
			"iss": "d7mekz",
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	return &claim
}

func (t TokenClaims) Generate() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, t)
	return token.SignedString([]byte(secret))
}

func validate(token string) (*TokenClaims, error) {
	parsed, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			// unexpected algorithm
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	// Validate Token Claim type.
	if claims, ok := parsed.Claims.(*TokenClaims); ok && parsed.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func IsNotExpired(token string) (bool, error) {
	_, err := validate(token)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

package auth

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
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

// Generate Generate()
// Create Access token
func (t TokenClaims) Generate() (string, error) {
	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, t)
	token, err := claim.SignedString([]byte(secret))
	return token, err
}

// validate()
// Validate right claims
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

// getIdFrom
// Get UserId from Access Token
func getIdFrom(token string) (int, error) {
	claim, err := validate(token)
	if err != nil {
		return -1, err
	}
	return claim.UserID, nil
}

// IsNotExpired
// Check Claim is expired or not
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

// getTokenFrom
// Get Token from header in context
// Return token string.
func getTokenFrom(ctx *gin.Context) (string, error) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		return "", fmt.Errorf("token is empty in header")
	}

	token, err := fromHeader(header)
	if err != nil {
		return "", fmt.Errorf("invalid Authorization value: %v", err)
	}
	return token, nil
}

// GetToken
// Get token from context and extract userId from claim
func GetToken(ctx *gin.Context) (int, error) {
	token, tokenErr := getTokenFrom(ctx)
	if tokenErr != nil {
		return -1, tokenErr
	}
	uid, validErr := getIdFrom(token)
	if validErr != nil {
		return -1, validErr
	}
	return uid, nil
}

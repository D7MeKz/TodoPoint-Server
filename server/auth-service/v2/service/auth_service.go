package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"modules/common/httputils"
	"modules/common/security/d7jwt"
	"strconv"
	"time"
	"todopoint/auth/v2/data"
)

//go:generate mockery --name Storer --case underscore
type Storer interface {
	Create(ctx *gin.Context, d interface{}) error
	IsExist(ctx *gin.Context, d interface{}) (bool, error)
	Modify(ctx *gin.Context, d interface{}) error
}

//go:generate mockery --name MysqlStorer --case underscore
type MysqlStorer interface {
	Storer

	GetId(ctx *gin.Context, data interface{}) (int, error)
}

type AuthService struct {
	redisStore Storer
	mysqlStore MysqlStorer
}

func NewAuthService(redis Storer, mysql MysqlStorer) *AuthService {
	return &AuthService{
		redisStore: redis,
		mysqlStore: mysql,
	}
}

func (a *AuthService) Login(ctx *gin.Context) (*httputils.BaseResponse, *httputils.NetError) {
	// Get crediential from header
	cred, err := GetCredential(ctx)
	if err != nil {
		return nil, httputils.NewNetError(httputils.BadAuthenticationData, err)
	}

	// Find user exist
	// if user does not exist, return error
	uid, err := a.mysqlStore.GetId(ctx, *cred)
	if err != nil {
		return nil, httputils.NewNetError(httputils.NotFound, err)
	}

	// if token generation failed, return error
	pair, err := GenerateTokenPair(uid)
	if err != nil {
		return nil, &httputils.NetError{Code: httputils.TokenCreateFailed, Err: err}
	}

	// Store token in redis
	param := data.RedisParams{Key: strconv.Itoa(uid), Value: pair.Refresh, Expires: time.Now().Add(time.Hour * 24 * 7).Unix()}
	err = a.redisStore.Create(ctx, param)
	if err != nil {
		return nil, httputils.NewNetError(httputils.TokenCreateFailed, err)
	}

	return httputils.NewSuccessBaseResponse(nil), nil
}

func (a *AuthService) Register(ctx *gin.Context, req data.RegisterRequest) (*httputils.BaseResponse, *httputils.NetError) {

	// Check user exist
	ok, err := a.mysqlStore.IsExist(ctx, data.Credential{Email: req.Email, Password: req.Password})
	if err != nil {
		return nil, httputils.NewNetError(httputils.NotFound, err)
	}

	// if user does not exist, create user
	if !ok {
		err = a.mysqlStore.Create(ctx, req)
		if err != nil {
			return nil, httputils.NewNetError(httputils.CreateFailed, err)
		}
		return httputils.NewSuccessBaseResponse(nil), nil
	}

	return nil, httputils.NewNetError(httputils.CreateFailed, errors.New("User already exist"))
}

func (a *AuthService) Issue(ctx *gin.Context) (*httputils.BaseResponse, *httputils.NetError) {
	// Check Refresh token is expired
	uid, err := extractIdFrom(ctx)
	if err != nil {
		return nil, httputils.NewNetError(httputils.TokenExpired, err)
	}

	// Issue new access token
	claim := d7jwt.NewTokenClaims(uid.Id)
	access, err := claim.Generate()
	if err != nil {
		return nil, httputils.NewNetError(httputils.TokenCreateFailed, err)
	}

	// Modify redis
	err = a.redisStore.Modify(ctx, data.RedisParams{Key: strconv.Itoa(uid.Id), Value: access})
	if err != nil {
		return nil, httputils.NewNetError(httputils.UpdateFailed, err)
	}
	return httputils.NewSuccessBaseResponse(access), nil
}

func (a *AuthService) Valid(ctx *gin.Context) (*httputils.BaseResponse, *httputils.NetError) {

	uid, err := extractIdFrom(ctx)

	if err != nil {
		return nil, httputils.NewNetError(httputils.TokenExpired, err)
	}

	return httputils.NewSuccessBaseResponse(uid), nil
}

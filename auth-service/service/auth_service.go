package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"modules/common/security/d7jwt"
	"modules/common/server/httpdata"
	"modules/common/server/httpdata/d7errors"
	"modules/common/server/httpdata/d7errors/codes"
	"strconv"
	"time"
	"todopoint/auth/data/dio"
	"todopoint/auth/data/params"
	"todopoint/auth/data/request"
)

type Storer interface {
	Create(ctx *gin.Context, data interface{}) error
	IsExist(ctx *gin.Context, data interface{}) (bool, error)
	Modify(ctx *gin.Context, data interface{}) error
}

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

func (a *AuthService) Login(ctx *gin.Context) (*httpdata.BaseResponse, *d7errors.NetError) {
	// Get crediential from header
	cred, err := GetCredential(ctx)
	if err != nil {
		return nil, d7errors.NewNetError(codes.BadAuthorizationHeader, err)
	}

	// Find user exist
	// if user does not exist, return error
	uid, err := a.mysqlStore.GetId(ctx, *cred)
	if err != nil {
		return nil, d7errors.NewNetError(codes.UserNotFound, err)
	}

	// if token generation failed, return error
	pair, err := GenerateTokenPair(uid)
	if err != nil {
		return nil, &d7errors.NetError{Code: codes.TokenCreationErr, Err: err}
	}

	// Store token in redis
	param := params.RedisParams{Key: strconv.Itoa(uid), Value: pair.Refresh, Expires: time.Now().Add(time.Hour * 24 * 7).Unix()}
	err = a.redisStore.Create(ctx, param)
	if err != nil {
		return nil, d7errors.NewNetError(codes.TokenCreationError, err)
	}

	return httpdata.NewBaseResponse(codes.UserLoginSuccess, pair), nil
}

func (a *AuthService) Register(ctx *gin.Context, req request.RegisterRequest) (*httpdata.BaseResponse, *d7errors.NetError) {

	// Check user exist
	ok, err := a.mysqlStore.IsExist(ctx, dio.Credential{Email: req.Email, Password: req.Password})
	if err != nil {
		return nil, d7errors.NewNetError(codes.UserInternalServerError, err)
	}

	// if user does not exist, create user
	if !ok {
		err = a.mysqlStore.Create(ctx, req)
		if err != nil {
			return nil, d7errors.NewNetError(codes.UserCreationError, err)
		}
		return httpdata.NewBaseResponse(codes.UserCreationSuccess, nil), nil
	}

	return nil, d7errors.NewNetError(codes.UserAlreadyExist, errors.New("User already exist"))
}

func (a *AuthService) Issue(ctx *gin.Context) (*httpdata.BaseResponse, *d7errors.NetError) {
	// Check Refresh token is expired
	uid, err := extractIdFrom(ctx)
	if err != nil {
		return nil, d7errors.NewNetError(codes.TokenInvalid, err)
	}

	// Issue new access token
	claim := d7jwt.NewTokenClaims(uid.Id)
	access, err := claim.Generate()
	if err != nil {
		return nil, d7errors.NewNetError(codes.TokenCreationError, err)
	}

	// Modify redis
	err = a.redisStore.Modify(ctx, params.RedisParams{Key: strconv.Itoa(uid.Id), Value: access})
	if err != nil {
		return nil, d7errors.NewNetError(codes.UserRedisSetErr, err)
	}
	return httpdata.NewBaseResponse(codes.UserTokenSetupSuccess, access), nil
}

func (a *AuthService) Valid(ctx *gin.Context) (*httpdata.BaseResponse, *d7errors.NetError) {

	uid, err := extractIdFrom(ctx)

	if err != nil {
		return nil, d7errors.NewNetError(codes.TokenExpired, err)
	}

	return httpdata.NewBaseResponse(codes.UserTokenSetupSuccess, uid), nil
}

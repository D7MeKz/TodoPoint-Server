package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"todopoint/auth/data/params"
	"todopoint/auth/data/request"
	"todopoint/common/d7errors"
	"todopoint/common/d7errors/codes"
	"todopoint/common/server/httpdata/auth"
	"todopoint/common/server/httpdata/response"
	"todopoint/common/server/httputils"
)

type Storer interface {
	Create(ctx *gin.Context, data interface{}) error
	IsExist(ctx *gin.Context, data interface{}) (bool, error)
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

func (a *AuthService) Login(ctx *gin.Context) (*response.BaseResponse, *d7errors.NetError) {
	// Get crediential from header
	cred, err := httputils.GetCredential(ctx)
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

	return response.NewBaseResponse(codes.UserLoginSuccess, pair), nil
}

func (a *AuthService) Register(ctx *gin.Context, req request.RegisterRequest) *d7errors.NetError {

	// Check user exist
	ok, err := a.mysqlStore.IsExist(ctx, auth.Credential{Email: req.Email, Password: req.Password})
	if err != nil {
		return d7errors.NewNetError(codes.UserInternalServerError, err)
	}

	// if user does not exist, create user
	if !ok {
		err = a.mysqlStore.Create(ctx, req)
		if err != nil {
			return d7errors.NewNetError(codes.UserCreationError, err)
		}
		return nil
	}

	return d7errors.NewNetError(codes.UserAlreadyExist, errors.New("User already exist"))
}

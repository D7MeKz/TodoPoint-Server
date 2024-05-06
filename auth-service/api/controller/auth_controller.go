package controller

import (
	"github.com/gin-gonic/gin"
	"todopoint/auth/data/request"
	"todopoint/common/d7errors"
	"todopoint/common/d7errors/codes"
	"todopoint/common/server/httpdata/response"
)

//go:generate mockery --name AuthOperator --case underscore
type AuthOperator interface {
	Login(ctx *gin.Context) (*response.BaseResponse, *d7errors.NetError)
	Register(ctx *gin.Context, req request.RegisterRequest) (*response.BaseResponse, *d7errors.NetError)
	Issue(ctx *gin.Context) (*response.BaseResponse, *d7errors.NetError)
	Valid(ctx *gin.Context) (*response.BaseResponse, *d7errors.NetError)
}

type AuthController struct {
	service AuthOperator
}

func NewAuthController(s AuthOperator) *AuthController {
	return &AuthController{
		service: s,
	}
}

// Login
// @Summary Login
// @Description Login
// @Tags auth-service
// @Accept json
// @Produce json
// @Router /auth-service/login [post]
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func (controller *AuthController) Login(ctx *gin.Context) {
	// Login
	res, err := controller.service.Login(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	res.Success(ctx)
	return
}

// Register
// @Summary Register user
// @Description Register user using email, password and username
// @Tags auth-service
// @Accept json
// @Produce json
// @Router /auth-service/register [post]
func (controller *AuthController) Register(ctx *gin.Context) {
	// Check body
	req := request.RegisterRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		_ = ctx.Error(d7errors.NewNetError(codes.UserInvalidJson, err))
		return
	}

	// Register user
	res, err := controller.service.Register(ctx, req)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	res.Success(ctx)
	return
}

// Issue
// @Summary Issue token
// @Description Issue refresh token when access token is expired
// @Tags auth-service
// @Accept json
// @Produce json
// @Router /auth-service/token [get]
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func (controller *AuthController) Issue(ctx *gin.Context) {
	// Issue token
	res, err := controller.service.Issue(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	res.Success(ctx)
	return

}

// Valid
// @Summary Check token is valid
// @Description Check token is valid
// @Tags auth-service
// @Accept json
// @Produce json
// @Router /auth-service/valid [get]
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func (controller *AuthController) Valid(ctx *gin.Context) {
	// Check token
	res, err := controller.service.Valid(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	res.Success(ctx)
	return
}

package api

import (
	"github.com/gin-gonic/gin"
	"modules/v2/common/httputils"
	"modules/v2/common/httputils/codes"
	"todopoint/auth/v2/data"
)

//go:generate mockery --name AuthOperator --case underscore
type AuthOperator interface {
	Login(ctx *gin.Context) (*httputils.BaseResponse, *httputils.NetError)
	Register(ctx *gin.Context, req data.RegisterRequest) (*httputils.BaseResponse, *httputils.NetError)
	Issue(ctx *gin.Context) (*httputils.BaseResponse, *httputils.NetError)
	Valid(ctx *gin.Context) (*httputils.BaseResponse, *httputils.NetError)
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
// @Tags auth
// @Accept json
// @Produce json
// @Router /auth/login [post]
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

	res.OKSuccess(ctx)
	return
}

// Register
// @Summary Register user
// @Description Register user using email, password and username
// @Tags auth
// @Accept json
// @Produce json
// @Router /auth/register [post]
func (controller *AuthController) Register(ctx *gin.Context) {
	// Check body
	req := data.RegisterRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		_ = ctx.Error(httputils.NewNetError(codes.CreateFailed, err))
		return
	}

	// Register user
	res, err := controller.service.Register(ctx, req)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	res.CreatedSuccess(ctx)
	return
}

// Issue
// @Summary Issue token
// @Description Issue refresh token when access token is expired
// @Tags auth
// @Accept json
// @Produce json
// @Router /auth/token [get]
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

	res.CreatedSuccess(ctx)
	return

}

// Valid
// @Summary Check token is valid
// @Description Check token is valid
// @Tags auth
// @Accept json
// @Produce json
// @Router /auth/valid [get]
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

	res.OKSuccess(ctx)
	return
}

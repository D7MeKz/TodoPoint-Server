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
	Register(ctx *gin.Context, req request.RegisterRequest) *d7errors.NetError
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

	res.Success(ctx)
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
	req := request.RegisterRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		_ = ctx.Error(d7errors.NewNetError(codes.UserInvalidJson, err))
		return
	}

	// Register user
	err = controller.service.Register(ctx, req)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	response.Success(ctx, codes.UserCreationSuccess)
	return
}

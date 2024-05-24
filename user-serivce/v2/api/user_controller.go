package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"modules/v2/common/httputils"
	"modules/v2/common/httputils/codes"
	"modules/v2/common/security/d7jwt"
	"todopoint/user/v2/data"
)

//go:generate mockery --name UserOperator --case underscore
type UserOperator interface {
	GetMe(ctx *gin.Context, uid int) (*httputils.BaseResponse, *httputils.NetError)
	Update(ctx *gin.Context, uid int, me data.Me) (*httputils.BaseResponse, *httputils.NetError)
}

type UserController struct {
	service UserOperator
}

func NewUserController(s UserOperator) *UserController {
	return &UserController{
		service: s,
	}
}

// GetMe
// @Summary GetMe
// @Description GetMe
// @Tags user
// @Accept json
// @Produce json
// @Router /user/get/me [get]
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func (controller *UserController) GetMe(ctx *gin.Context) {
	// GetMe
	// Get uid from token
	uid, err := d7jwt.Validate(ctx)
	fmt.Println(uid)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	res, err := controller.service.GetMe(ctx, uid)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	res.OKSuccess(ctx)
}

func (controller *UserController) UpdateProfile(ctx *gin.Context) {
	// GetMe
	// Get uid from token
	uid, err := d7jwt.Validate(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	req := data.Me{}
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		_ = ctx.Error(httputils.NewNetError(codes.InvalidBody, err))
		return
	}

	res, uerr := controller.service.Update(ctx, uid, req)
	if uerr != nil {
		_ = ctx.Error(err)
		return
	}

	res.OKSuccess(ctx)
	return
}

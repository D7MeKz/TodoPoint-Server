package api

import (
	"github.com/gin-gonic/gin"
	"modules/common/httputils"
	"modules/common/security/d7jwt"
)

//go:generate mockery --name UserOperator --case underscore
type UserOperator interface {
	GetMe(ctx *gin.Context, uid int) (*httputils.BaseResponse, *httputils.NetError)
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
	uid, err := d7jwt.GetIdFromHeader(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	res, err := controller.service.GetMe(ctx, uid)
	if err != nil {
		_ = ctx.Error
		return
	}
	res.OKSuccess(ctx)
}

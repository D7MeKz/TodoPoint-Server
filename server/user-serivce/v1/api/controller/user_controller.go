package controller

import (
	"github.com/gin-gonic/gin"
	"modules/common/server/httpdata"
	"modules/common/server/httpdata/d7errors"
)

type UserOperator interface {
	Create(ctx *gin.Context) (*httpdata.BaseResponse, *d7errors.NetError)
	Me(ctx *gin.Context) (*httpdata.BaseResponse, *d7errors.NetError)
}

type UserController struct {
	service UserOperator
}

func NewUserController(s UserOperator) *UserController {
	return &UserController{
		service: s,
	}
}

// Me
// @Summary Me
// @Description Get user profile information
// @Tags user
// @Accept json
// @Produce json
// @Router /user/me [get]
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func (controller *UserController) Me(ctx *gin.Context) {
	// Me
	res, err := controller.service.Me(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	res.Success(ctx)
	return
	
}

func (controller *UserController) Create(ctx *gin.Context) {
	// Create
	res, err := controller.service.Create(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	res.Success(ctx)
	return
}

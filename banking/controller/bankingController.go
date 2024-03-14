package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todopoint/banking/data/request"
	"todopoint/banking/data/response"
	"todopoint/banking/service"
)

// 추상화
type BankAccountController struct {
	service service.BankAccountService
}

func NewBankAccountController(s *service.BankAccountService) *BankAccountController {
	return &BankAccountController{
		service: s,
	}
}

func (controller *BankAccountController) RegisterAccount(ctx *gin.Context) {
	req := request.CreateReqData{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := controller.service.CreateBankAccount(ctx, req)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	response.Success(ctx, res)
}

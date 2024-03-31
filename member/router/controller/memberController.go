package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"todopoint/common/netutils/codes"
	"todopoint/common/netutils/errorutils"
	"todopoint/common/netutils/response"
	"todopoint/member/data"
	"todopoint/member/service"
)

type MemberController struct {
	service service.MemberService
}

func NewMemberController(s service.MemberService) *MemberController {
	return &MemberController{
		service: s,
	}
}

func (controller *MemberController) RegisterMember(ctx *gin.Context) {
	req := data.RegisterReq{}
	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		_ = ctx.Error(errorutils.NewNetError(codes.MemberInvalidJson, err))
		return
	}

	// Create member
	mem, err := controller.service.CreateMember(ctx, req)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	mid := data.MemberId{MemberId: mem.ID}
	response.SuccessWith(ctx, codes.MemberCreationSuccess, mid)
}

func (controller *MemberController) LoginMember(ctx *gin.Context) {
	req := data.LoginReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		_ = ctx.Error(errorutils.NewNetError(codes.MemberInvalidJson, err))
		return
	}
	// login Member
	memId, err := controller.service.LoginMember(ctx, req)
	if err != nil {
		_ = ctx.Error(err)
	}

	res := data.MemberId{MemberId: memId}
	response.SuccessWith(ctx, codes.MemberLoginSuccess, res)

}

func (controller *MemberController) IsValidMember(ctx *gin.Context) {
	strId, ok := ctx.Params.Get("memId")
	if ok == false {
		_ = ctx.Error(errorutils.NewNetError(codes.MemberInvalidUri, nil))
		return
	}
	log.Println(strId)

	memId, err := strconv.Atoi(strId)
	if err != nil {
		_ = ctx.Error(errorutils.NewNetError(codes.MemberInvalidUri, err))
		return
	}
	ok, err = controller.service.CheckIsValid(ctx, memId)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	response.Success(ctx, codes.MemberOK)
}

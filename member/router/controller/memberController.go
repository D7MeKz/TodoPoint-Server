package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"todopoint/common/errorutils"
	"todopoint/common/errorutils/codes"
	"todopoint/common/netutils/response"
	"todopoint/member/service"
	"todopoint/member/utils/data"
)

type MemberController struct {
	service service.MemberService
}

////go:generate mockery --name MemberService --case underscore
//type MemberService interface {
//	CreateMember(ctx *gin.Context, req data.RegisterReq) (*ent.Member, *errorutils.NetError)
//	LoginMember(ctx *gin.Context, req data.LoginReq) (int, *errorutils.NetError)
//	CheckIsValid(ctx *gin.Context, memId int) (bool, *errorutils.NetError)
//}

func NewMemberController(s service.MemberService) *MemberController {
	return &MemberController{
		service: s,
	}
}

// RegisterMember
// @Summary Register Member
// @Description Register Member
// @Tags members
// @Accept json
// @Produce json
// @Param request body data.RegisterReq true "query params"
// @Success 200 {object} data.MemberId
// @Router /members/register [post]
func (controller *MemberController) RegisterMember(ctx *gin.Context) {
	req := data.RegisterReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		_ = ctx.Error(errorutils.NewNetError(codes.MemberInvalidJson, err))
		return
	}

	// Create member
	mem, err2 := controller.service.CreateMember(ctx, req)
	if err2 != nil {
		_ = ctx.Error(err2)
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

	memId, err := strconv.Atoi(strId)
	if err != nil {
		_ = ctx.Error(errorutils.NewNetError(codes.MemberInvalidUri, err))
		return
	}
	ok, err2 := controller.service.CheckIsValid(ctx, memId)
	if err2 != nil {
		_ = ctx.Error(err2)
		return
	}
	response.Success(ctx, codes.MemberOK)
}

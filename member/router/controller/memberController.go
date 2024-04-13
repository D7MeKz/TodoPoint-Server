package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
// @Router /auth/register [post]
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

// LoginMember
// @Summary Login Member
// @Description If you login, Create tokens(Refresh, Access Token)
// @Tags members
// @Accept json
// @Produce json
// @Param request body data.LoginReq true "query params"
// @Success 200 {object} data.TokenPair
// @Router /auth/login [post]
func (controller *MemberController) LoginMember(ctx *gin.Context) {
	// Get body
	req := data.LoginReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logrus.Errorf("Invalid Json format : %v ", err)
		_ = ctx.Error(errorutils.NewNetError(codes.MemberInvalidJson, err))
		return
	}

	// login Member
	pair, err2 := controller.service.LoginMember(ctx, req)
	if err2 != nil {
		logrus.Errorf("Login failed")
		_ = ctx.Error(err2)
		return
	}
	logrus.Debug(pair)
	response.SuccessWith(ctx, codes.MemberLoginSuccess, pair)

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

// RefreshToken
// @Summary Generate Refresh Token
// @Description
// If access token is expired, client should request with refresh token in body.
// Service checks the refresh token expiration. If it does, generate new one.
// @Tags members
// @Accept json
// @Produce json
// @Param request body data.RefreshToken true "query params"
// @Success 200 {object} data.AccessToken
// @Router /auth/token [post]
func (controller *MemberController) RefreshToken(ctx *gin.Context) {
	req := data.RefreshToken{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		_ = ctx.Error(errorutils.NewNetError(codes.MemberInvalidJson, err))
		return
	}

	// Generate New Token
	access, err := controller.service.GenerateNewToken(ctx, req)
	if err != nil {
		_ = ctx.Error(err)
	}

	response.SuccessWith(ctx, codes.MemberLoginSuccess, access)
}

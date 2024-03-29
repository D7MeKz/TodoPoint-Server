package service

import (
	"github.com/gin-gonic/gin"
	"todopoint/member/out/ent"

	wu "todopoint/common/webutils"
	"todopoint/member/data/request"
)

//go:generate mockery --name MemberStore --case underscore
type MemberStore interface {
	Create(ctx *gin.Context, req request.RegisterReq) (*ent.Member, *wu.Error)
	GetById(ctx *gin.Context, memberId int) (*ent.Member, *wu.Error)
	GetMemberByEmail(ctx *gin.Context, email string) (*ent.Member, *wu.Error)
	GetIDByLogin(ctx *gin.Context, req request.LoginReq) (int, error)
	IsExistByID(ctx *gin.Context, memId int) (bool, error)
}

type MemberService struct {
	Store MemberStore
}

func NewMemberService(s MemberStore) *MemberService {
	return &MemberService{Store: s}
}

func (s *MemberService) CreateMember(ctx *gin.Context, req request.RegisterReq) *ent.Member {
	// Check member Exist
	existedMem, err := s.Store.GetMemberByEmail(ctx, req.Email)
	//if err != nil {
	//	wu.ErrorFunc(ctx, err)
	//}

	// Create Member
	if existedMem != nil {
		return existedMem
	}
	mem, err := s.Store.Create(ctx, req)
	if err != nil {
		wu.ErrorFunc(ctx, err)
		return nil
	}
	return mem
}

func (s *MemberService) LoginMember(ctx *gin.Context, req request.LoginReq) int {
	memId, err := s.Store.GetIDByLogin(ctx, req)
	if err != nil {
		wu.ErrorFunc(ctx, wu.NewError(wu.LOGIN_FAILED, err))
		return -1
	}
	return memId
}

func (s *MemberService) CheckIsValid(ctx *gin.Context, memId int) bool {
	isExist, err := s.Store.IsExistByID(ctx, memId)
	if err != nil {
		wu.ErrorFunc(ctx, wu.NewError(wu.INVALID_MEMBER, err))
		return false
	}
	return isExist
}

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
	IsExistByEmail(ctx *gin.Context, email string) (*ent.Member, *wu.Error)
	IsExistByLogin(ctx *gin.Context, req request.LoginReq) (int, error)
}

type MemberService struct {
	Store MemberStore
}

func NewMemberService(s MemberStore) *MemberService {
	return &MemberService{Store: s}
}

func (s *MemberService) CreateMember(ctx *gin.Context, req request.RegisterReq) *ent.Member {
	// Check member Exist
	existedMem, err := s.Store.IsExistByEmail(ctx, req.Email)
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

func (s *MemberService) LoginMember(ctx *gin.Context, req request.LoginReq) (int, *wu.Error) {
	memId, err := s.Store.IsExistByLogin(ctx, req)
	if err != nil {
		return -1, wu.NewError(wu.LOGIN_FAILED, err)
	}
	return memId, nil
}

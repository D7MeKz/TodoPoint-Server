package service

import (
	"github.com/gin-gonic/gin"
	"todopoint/common/db/ent"
	wu "todopoint/common/webutils"
	"todopoint/member/data/request"
)

//go:generate mockery --name MemberStore --case underscore
type MemberStore interface {
	Create(ctx *gin.Context, req request.RegisterReq) (*ent.Member, *wu.Error)
	GetById(ctx *gin.Context, memberId int) (*ent.Member, *wu.Error)
	IsExist(ctx *gin.Context, email string) (*ent.Member, *wu.Error)
}

type MemberService struct {
	Store MemberStore
}

func NewMemberService(s MemberStore) *MemberService {
	return &MemberService{Store: s}
}

func (s *MemberService) CreateMember(ctx *gin.Context, req request.RegisterReq) *ent.Member {

	// Check member Exist
	isExist, err := s.Store.IsExist(ctx, req.Email)
	//if err != nil {
	//	wu.ErrorFunc(ctx, err)
	//}

	// Create Member
	if isExist != nil {
		return isExist
	}
	mem, err := s.Store.Create(ctx, req)
	if err != nil {
		wu.ErrorFunc(ctx, err)
		return nil
	}
	return mem
}

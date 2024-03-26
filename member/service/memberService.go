package service

import (
	"github.com/gin-gonic/gin"
	"todopoint/common/db/ent"
	"todopoint/common/errorutils"
	"todopoint/member/data/request"
)

//go:generate mockery --name MemberStore --case underscore
type MemberStore interface {
	Create(ctx *gin.Context, req request.RegisterReq) error
	GetById(ctx *gin.Context, memberId int) (*ent.Member, error)
}

type MemberService struct {
	Store MemberStore
}

func NewMemberService(s MemberStore) *MemberService {
	return &MemberService{Store: s}
}

func (s *MemberService) CreateMember(ctx *gin.Context, req request.RegisterReq) *errorutils.ErrorBox {
	err := s.Store.Create(ctx, req)
	if err != nil {
		return errorutils.NewErrorBox(errorutils.ERROR_TASK_DB, err, "")
	}
	return nil
}

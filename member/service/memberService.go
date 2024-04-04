package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"todopoint/common/netutils/codes"
	"todopoint/common/netutils/errorutils"
	"todopoint/member/data"
	"todopoint/member/out/ent"
)

//go:generate mockery --name MemberStore --case underscore
type MemberStore interface {
	Create(ctx *gin.Context, req data.RegisterReq) (*ent.Member, error)
	GetById(ctx *gin.Context, memberId int) (*ent.Member, error)
	GetMemberByEmail(ctx *gin.Context, email string) (*ent.Member, error)
	GetIDByLogin(ctx *gin.Context, req data.LoginReq) (int, error)
	IsExistByID(ctx *gin.Context, memId int) (bool, error)
}

type MemberService struct {
	Store MemberStore
}

func NewMemberService(s MemberStore) *MemberService {
	return &MemberService{Store: s}
}

func (s *MemberService) CreateMember(ctx *gin.Context, req data.RegisterReq) (*ent.Member, *errorutils.NetError) {
	// Check member Exist
	_, err := s.Store.GetMemberByEmail(ctx, req.Email)
	if err != nil && !ent.IsNotFound(err) {
		return nil, &errorutils.NetError{Code: codes.MemberInternalServerError, Err: err}
	}

	mem, err := s.Store.Create(ctx, req)
	if err != nil {
		return nil, &errorutils.NetError{Code: codes.MemberCreationError, Err: err}
	}
	return mem, nil
}

func (s *MemberService) LoginMember(ctx *gin.Context, req data.LoginReq) (int, *errorutils.NetError) {
	memId, err := s.Store.GetIDByLogin(ctx, req)
	if err != nil {
		if ent.IsNotFound(err) {
			return -1, &errorutils.NetError{Code: codes.MemberNotFound, Err: err}
		} else {
			return -1, &errorutils.NetError{Code: codes.MemberInternalServerError, Err: err}
		}
	}
	return memId, nil
}

func (s *MemberService) CheckIsValid(ctx *gin.Context, memId int) (bool, *errorutils.NetError) {
	ok, err := s.Store.IsExistByID(ctx, memId)

	if ok == false || err != nil {
		if ent.IsNotFound(err) {
			return false, &errorutils.NetError{Code: codes.MemberNotFound, Err: err}
		}
		return false, &errorutils.NetError{Code: codes.MemberInternalServerError, Err: err}
	}
	logrus.Warn(ok)

	uuid.New()

	return true, nil
}

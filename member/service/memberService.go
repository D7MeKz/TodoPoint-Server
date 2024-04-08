package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"todopoint/common/errorutils"
	"todopoint/common/errorutils/codes"
	"todopoint/member/out/ent"
	"todopoint/member/utils/data"
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
	existedMem, err := s.Store.GetMemberByEmail(ctx, req.Email)
	if err != nil {
		logrus.Print("Get By Email Error")
		return nil, &errorutils.NetError{Code: codes.MemberInternalServerError, Err: err}
	}
	if ent.IsNotFound(err) {
		logrus.Print("Member does not exist")
		mem, err2 := s.Store.Create(ctx, req)
		if err2 != nil {
			return nil, &errorutils.NetError{Code: codes.MemberCreationError, Err: err2}
		}
		return mem, nil
	}
	return existedMem, nil
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
			logrus.Warn("Member Does not Exist")
			return false, &errorutils.NetError{Code: codes.MemberNotFound, Err: err}
		}
		return false, &errorutils.NetError{Code: codes.MemberInternalServerError, Err: err}
	}

	return true, nil
}

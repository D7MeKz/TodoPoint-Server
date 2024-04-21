package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
	"todopoint/common/auth"
	"todopoint/common/errorutils"
	"todopoint/common/errorutils/codes"
	"todopoint/member/out/ent"
	"todopoint/member/out/persistence"
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

	if ent.IsNotFound(err) {
		logrus.Print("Member does not exist")
		mem, err2 := s.Store.Create(ctx, req)
		if err2 != nil {
			return nil, &errorutils.NetError{Code: codes.MemberCreationError, Err: err2}
		}
		return mem, nil
	}
	if err != nil {
		logrus.Print("Get By Email Error")
		return nil, &errorutils.NetError{Code: codes.MemberInternalServerError, Err: err}
	}
	return existedMem, nil
}

func (s *MemberService) LoginMember(ctx *gin.Context, req data.LoginReq) (*data.TokenPair, *errorutils.NetError) {
	// Verify User Exist
	memId, err := s.Store.GetIDByLogin(ctx, req)
	if err != nil {
		if ent.IsNotFound(err) {
			logrus.Errorf("Member not found : %v", err)
			return nil, &errorutils.NetError{Code: codes.MemberNotFound, Err: err}
		} else {
			logrus.Errorf("Internal server error : %v", err)
			return nil, &errorutils.NetError{Code: codes.MemberInternalServerError, Err: err}
		}
	}
	logrus.Debugf("Get memberId from login : %d", memId)

	// Create Token
	claim := auth.NewTokenClaims(memId)
	access, err := claim.Generate()
	if err != nil {
		logrus.Errorf("Token creation Error : %v", err)
		return nil, &errorutils.NetError{Code: codes.TokenCreationErr, Err: err}
	}
	logrus.Debug("Success : Access Token generation")

	// Create Access, Refresh Token
	refresh := uuid.NewString()
	redisStore := persistence.NewRedisStore()
	expires := time.Now().Add(time.Hour * 24 * 7).Unix()
	redisErr := redisStore.Create(ctx, refresh, strconv.Itoa(memId), expires)
	if redisErr != nil {
		logrus.Error(redisErr)
		return nil, &errorutils.NetError{Code: codes.TokenCreationError, Err: err}
	}
	logrus.Debug("Success : Refresh Token generation")

	return &data.TokenPair{AccessToken: access, RefreshToken: refresh}, nil
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

func (s *MemberService) GenerateNewToken(ctx *gin.Context, token data.RefreshToken) (*data.AccessToken, *errorutils.NetError) {
	// Check refresh token validation
	redisStore := persistence.NewRedisStore()
	memId, err := redisStore.Find(ctx, token.RefreshToken)
	// If redis value did not exist, response error. Login again
	if err != nil {
		return nil, &errorutils.NetError{Code: codes.TokenExpired, Err: err}
	}

	// Generate new access token
	claim := auth.NewTokenClaims(memId)
	access, err := claim.Generate()
	if err != nil {
		return nil, &errorutils.NetError{Code: codes.TokenCreationErr, Err: err}
	}

	return &data.AccessToken{AccessToken: access}, nil
}

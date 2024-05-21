package service

import (
	"github.com/gin-gonic/gin"
	"modules/v2/common/httputils"
	"modules/v2/common/httputils/codes"
	"modules/v2/d7mysql/ent"
	"todopoint/user/v2/data"
)

//go:generate mockery --name Store --case underscore
type Store interface {
	Create(ctx *gin.Context, info *data.UserInfo) error
	FindOne(ctx *gin.Context, uid int) (*data.Me, error)
	Update(ctx *gin.Context, uid int, me data.Me) error
}

type UserService struct {
	store Store
}

func NewUserService(store Store) *UserService {
	return &UserService{
		store: store,
	}
}

func (s *UserService) GetMe(ctx *gin.Context, uid int) (*httputils.BaseResponse, *httputils.NetError) {
	me, err := s.store.FindOne(ctx, uid)
	if ent.IsNotFound(err) {
		return nil, httputils.NewNetError(codes.NotFound, err)
	} else if err != nil {
		return nil, httputils.NewNetError(codes.FindFail, err)
	}

	return httputils.NewSuccessBaseResponse(me), nil
}

func (s *UserService) Update(ctx *gin.Context, uid int, me data.Me) (*httputils.BaseResponse, *httputils.NetError) {
	err := s.store.Update(ctx, uid, me)
	if err != nil {
		return nil, httputils.NewNetError(codes.UpdateFailed, err)
	}

	return httputils.NewSuccessBaseResponse(nil), nil
}

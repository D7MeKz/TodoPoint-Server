package service

import (
	"github.com/gin-gonic/gin"
	"modules/common/httputils"
	"modules/common/httputils/codes"
	"modules/d7mysql/ent"
	"todopoint/user/v2/data"
)

//go:generate mockery --name Store --case underscore
type Store interface {
	Create(ctx *gin.Context, info *data.UserInfo) error
	FindOne(ctx *gin.Context, uid int) (*data.Me, error)
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

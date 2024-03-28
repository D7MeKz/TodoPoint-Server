package persistence

import (
	"github.com/gin-gonic/gin"
	"todopoint/common/db/config"
	"todopoint/common/db/ent"
	"todopoint/common/db/ent/member"
	wu "todopoint/common/webutils"
	"todopoint/member/data/request"
)

type Store struct {
	client *ent.Client
}

func NewStore() *Store {
	return &Store{
		client: config.GetClient(),
	}
}

func (s *Store) Create(ctx *gin.Context, req request.RegisterReq) (*ent.Member, *wu.Error) {
	// create Task
	mem, err := s.client.Member.Create().SetEmail(req.Email).SetUsername("").SetPassword(req.Password).Save(ctx)
	if err != nil {
		return nil, wu.NewError(wu.ERROR_MEMBER_DB, err)
	}
	return mem, nil
}

func (s *Store) GetById(ctx *gin.Context, memberId int) (*ent.Member, *wu.Error) {
	mem, err := s.client.Member.Get(ctx, memberId)
	if err != nil {
		return nil, wu.NewError(wu.ERROR_MEMBER_DB, err)
	}
	return mem, nil
}

func (s *Store) IsExist(ctx *gin.Context, email string) (*ent.Member, *wu.Error) {
	mem, err := s.client.Member.Query().Where(member.EmailEQ(email)).First(ctx)
	if err != nil {
		return nil, wu.NewError(wu.ERROR_MEMBER_DB, err)
	}
	if mem != nil {
		return mem, nil
	}
	return nil, nil
}

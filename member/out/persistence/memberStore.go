package persistence

import (
	"github.com/gin-gonic/gin"
	"todopoint/common/db/config"
	"todopoint/common/db/ent"
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

func (s *Store) Create(ctx *gin.Context, req request.RegisterReq) error {
	// create Task
	_, err := s.client.Member.Create().SetEmail(req.Email).SetUsername(req.Username).SetPassword(req.Password).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetById(ctx *gin.Context, memberId int) (*ent.Member, error) {
	member, err := s.client.Member.Get(ctx, memberId)
	if err != nil {
		return nil, err
	}
	return member, nil
}

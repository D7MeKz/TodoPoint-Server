package persistence

import (
	"github.com/gin-gonic/gin"
	"todopoint/member/config"
	"todopoint/member/out/ent"
	"todopoint/member/out/ent/member"

	wu "todopoint/common/webutils"
	"todopoint/member/data/request"
)

//go:generate mockery --name Store --case underscore --inpackage
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
	mem, err := s.client.Member.Create().SetEmail(req.Email).SetUsername(req.Username).SetPassword(req.Password).Save(ctx)
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

func (s *Store) IsExistByEmail(ctx *gin.Context, email string) (*ent.Member, *wu.Error) {
	mem, err := s.client.Member.Query().Where(member.EmailEQ(email)).First(ctx)
	if err != nil {
		return nil, wu.NewError(wu.ERROR_MEMBER_DB, err)
	}
	if mem != nil {
		return mem, nil
	}
	return nil, nil

}

func (s *Store) IsExistByLogin(ctx *gin.Context, req request.LoginReq) (int, error) {
	mem, err := s.client.Member.Query().Where(member.EmailEQ(req.Email), member.PasswordEQ(req.Password)).First(ctx)

	if err != nil {
		return -1, err
	}
	return mem.ID, nil
}

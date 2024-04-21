package persistence

import (
	"github.com/gin-gonic/gin"
	"todopoint/member/out/ent"
	"todopoint/member/out/ent/member"
	ent2 "todopoint/member/utils/config"
	"todopoint/member/utils/data"
)

//go:generate mockery --name Store --case underscore --inpackage
type Store struct {
	client *ent.Client
}

func NewStore() *Store {
	return &Store{
		client: ent2.GetClient(),
	}
}

func (s *Store) Create(ctx *gin.Context, req data.RegisterReq) (*ent.Member, error) {
	// create Task
	mem, err := s.client.Member.Create().SetEmail(req.Email).SetUsername(req.Username).SetPassword(req.Password).Save(ctx)
	if err != nil {
		return nil, err
	}
	return mem, nil
}

func (s *Store) GetById(ctx *gin.Context, memberId int) (*ent.Member, error) {
	mem, err := s.client.Member.Get(ctx, memberId)
	if err != nil {
		return nil, err
	}
	return mem, nil
}

func (s *Store) GetMemberByEmail(ctx *gin.Context, email string) (*ent.Member, error) {
	mem, err := s.client.Member.Query().Where(member.EmailEQ(email)).First(ctx)
	if err != nil {
		return nil, err
	}
	return mem, nil
}

func (s *Store) GetIDByLogin(ctx *gin.Context, req data.LoginReq) (int, error) {
	mem, err := s.client.Member.Query().Where(member.EmailEQ(req.Email), member.PasswordEQ(req.Password)).First(ctx)
	if err != nil {
		return -1, err
	}
	return mem.ID, nil
}

func (s *Store) IsExistByID(ctx *gin.Context, memId int) (bool, error) {
	_, err := s.client.Member.Get(ctx, memId)
	if err != nil {
		return false, err
	}
	return true, nil
}

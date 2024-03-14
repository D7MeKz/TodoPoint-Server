package repo

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"todopoint/banking/config"
	"todopoint/banking/data/request"
	"todopoint/banking/models/ent"
)

type Store struct {
	ctx    context.Context
	client *ent.BankAccountClient
}

func NewStore() *Store {
	return &Store{
		client: config.GetClient().BankAccount,
	}
}

func (s *Store) Create(ctx *gin.Context, b request.CreateReqData) (*ent.BankAccount, error) {
	account, err := s.client.Create().
		SetUserID(b.UserId).
		SetBankName(b.BankName).
		Save(ctx)
	if err != nil {
		log.Println("Create Member Error")
		return nil, err
	}
	return account, nil
}

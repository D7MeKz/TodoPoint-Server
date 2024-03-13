package service

import (
	"context"
	"todopoint/banking/config"
	"todopoint/banking/models/ent"
)

type BankAccountRepo struct {
	ctx    context.Context
	client *ent.BankAccountClient
}

func NewMemberRepo(ctx context.Context) *BankAccountRepo {
	return &BankAccountRepo{
		ctx:    ctx,
		client: config.GetClient().BankAccount,
	}
}

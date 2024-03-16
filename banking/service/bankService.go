package service

import (
	"github.com/gin-gonic/gin"
	"todopoint/banking/data/request"
	"todopoint/banking/data/response"
	"todopoint/banking/models/ent"
)

type BankAccountStore interface {
	Create(ctx *gin.Context, b request.CreateReqData) (*ent.BankAccount, error)
}

// Abstraction
type BankAccountService struct {
	store BankAccountStore
}

func NewBankAccountService(s BankAccountStore) *BankAccountService {
	return &BankAccountService{
		store: s,
	}
}

func (bs *BankAccountService) CreateBankAccount(ctx *gin.Context, data request.CreateReqData) (response.CreateAccountRes, error) {
	account, err := bs.store.Create(ctx, data)
	if err != nil {
		// Cannot use 'nil' as the type
		// return empty value
		return response.CreateAccountRes{}, err
	}
	res := response.CreateAccountRes{
		BankAccount: account.BankAccount.String(), // UUID to string
	}
	return res, nil

}

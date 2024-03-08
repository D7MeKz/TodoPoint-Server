package main

import (
	"context"
	"net/http"
)

type Service interface {
	GetMembership(context.Context) (*Member, error)
}

type MemberService struct {
	url string
}

func NewMembershipService(url string) Service {
	return &MemberService{
		url: url,
	}
}

func (s *MemberService) GetMembership(ctx context.Context) (*Member, error) {
	resp, err := http.Get()
	if err != nil {
		return nil, err
	}
	fact := &Member{}

}

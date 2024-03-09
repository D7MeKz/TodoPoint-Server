package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	GetMembership(context.Context) (*Member, error)
}

type MemberService struct {
	url string
}

func NewMemberService(url string) Service {
	return &MemberService{
		url: url,
	}
}

func (s *MemberService) GetMembership(ctx context.Context) (*Member, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}
	member := &Member{}
	if err := json.NewDecoder(resp.Body).Decode(member); err != nil {
		return nil, err
	}
	return member, nil
}

package service

import (
	"context"
	"log"
	"todopoint/member/internal/config"
	"todopoint/member/internal/ent"
	"todopoint/member/internal/model"
)

type MemberRepo struct {
	ctx    context.Context
	client *ent.MemberClient
}

func NewMemberRepo(ctx context.Context) *MemberRepo {
	return &MemberRepo{
		ctx:    ctx,
		client: config.GetClient().Member,
	}
}

func (m *MemberRepo) CreateMember(newMember model.Member) (*ent.Member, error) {
	// Create Member
	member, err := m.client.Create().
		SetUsername(newMember.Username).
		SetEmail(newMember.Email).
		SetPassword(newMember.Password).
		Save(m.ctx)
	if err != nil {
		log.Println("Create Member Error")
		return nil, err
	}
	return member, nil
}

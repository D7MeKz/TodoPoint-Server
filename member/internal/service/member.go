package service

import (
	"context"
	"log"
	"todopoint/db/ent"
	"todopoint/member/internal/config"
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

func (m *MemberRepo) GetAllMembers() ([]*ent.Member, error) {
	members, err := m.client.Query().All(m.ctx)
	if err != nil {
		return nil, err
	}
	return members, nil
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

func (m *MemberRepo) GetMemberByID(id int) (*ent.Member, error) {
	member, err := m.client.Get(m.ctx, id)
	if err != nil {
		return nil, err
	}
	return member, nil
}

func (m *MemberRepo) DeleteMember(id int) (int, error) {
	err := m.client.DeleteOneID(id).Exec(m.ctx)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *MemberRepo) UpdateMember(id int, member model.Member) (*ent.Member, error) {
	updatedMember, err := m.client.UpdateOneID(id).
		SetEmail(member.Email).
		SetPassword(member.Password).
		SetUsername(member.Username).Save(m.ctx)
	if err != nil {
		return nil, err
	}
	return updatedMember, nil
}

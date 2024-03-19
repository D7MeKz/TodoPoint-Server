package persistence

import (
	"github.com/gin-gonic/gin"
	"todopoint/common/db/config"
	"todopoint/common/db/ent"
	"todopoint/task/data/request"
)

type Store struct {
	client *ent.Client
}

func NewStore() *Store {
	return &Store{
		client: config.GetClient(),
	}
}

func (s *Store) Create(ctx *gin.Context, b request.CreateTask) error {
	// create Task
	_, err := s.client.Task.Create().SetTitle(b.Title).SetTotalStatus(0).SetMemberID(b.UserId).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetAllTasks(ctx *gin.Context, memberId int) ([]*ent.Task, error) {
	// Get All task
	tasks, err := s.client.Task.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, err
}

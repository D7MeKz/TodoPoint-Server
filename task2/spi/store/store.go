package store

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"todopoint/task2/data"
	"todopoint/task2/data/model"
)

type TaskStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewTaskStore(client *mongo.Client, collection *mongo.Collection) *TaskStore {
	return &TaskStore{
		client:     client,
		collection: collection,
	}
}

func (s *TaskStore) Create(
	ctx *gin.Context,
	formatter data.DioFormatter,
) error {
	if data, ok := formatter.(data.UserId); ok {
		// 런타임 에러를 발생시키지 않음
		taskData := model.NewTask(data.UserId)
		_, err := s.collection.InsertOne(ctx, taskData)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("Format error")
}

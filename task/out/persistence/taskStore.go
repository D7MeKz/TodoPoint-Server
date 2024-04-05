package persistence

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"todopoint/common/config/mongodb"
	"todopoint/task/data"
)

type Store struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewStore(colName string) *Store {
	client := mongodb.GetClient()
	return &Store{
		client:     client,
		collection: mongodb.GetCollection(client, colName),
	}
}

func (s *Store) Create(ctx *gin.Context, req data.CreateReq) (*mongo.InsertOneResult, error) {
	taskData := data.NewTaskInfo(req)
	id, err := s.collection.InsertOne(ctx, taskData)
	if err != nil {
		return nil, err
	}
	return id, nil

}

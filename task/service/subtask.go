package service

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"todopoint/task/data"
	"todopoint/task/out/persistence"
)

type SubTaskStore interface {
	Create(ctx *gin.Context, req data.CreateReq) (*mongo.InsertOneResult, error)
	IsExist(ctx *gin.Context, taskId int) (bool, error)
}

type SubTaskService struct {
	Store SubTaskStore
}

func NewSubTaskService(s *persistence.Store) *SubTaskService {
	return &SubTaskService{
		Store: s,
	}
}

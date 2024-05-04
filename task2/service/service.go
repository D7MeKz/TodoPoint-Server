package service

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"todopoint/task2/data"
)

type Store interface {
	Create(ctx *gin.Context, data data.DioFormatter) (*mongo.InsertOneResult, error)
	//Update(ctx *gin.Context, data interface{}) error
	//Delete(ctx *gin.Context, id string) error
	//
	//Finder
}

type Finder interface {
	FindOne(ctx *gin.Context, id string) error
	FindMany(ctx *gin.Context, count int) error
}

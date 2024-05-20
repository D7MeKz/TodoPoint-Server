package spi

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"todopoint/image/data/model"
)

type ImageStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewImageStore(client *mongo.Client) *ImageStore {
	return &ImageStore{
		client:     client,
		collection: client.Database("todopoint").Collection("profile_image"),
	}
}

func (i *ImageStore) Add(ctx *gin.Context, uid int, path string) error {
	img := model.NewImageModel(uid, path)
	_, err := i.collection.InsertOne(ctx, img)
	if err != nil {
		return err
	}
	return nil
}

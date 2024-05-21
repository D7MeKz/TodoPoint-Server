package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ImageModel struct {
	ImageID   primitive.ObjectID `bson:"_id,omitempty" json:"image_id"`
	UserID    int                `bson:"user_id,omitempty" json:"user_id"`
	Path      string             `bson:"path" json:"path"`
	CreatedAt string             `bson:"created_at" json:"created_at"`
}

func NewImageModel(uid int, path string) *ImageModel {
	return &ImageModel{
		UserID:    uid,
		Path:      path,
		CreatedAt: time.Now().Format("2006-01-02 15:01"),
	}
}

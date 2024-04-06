package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Task struct {
	ObjectID   primitive.ObjectID `bson:"_id" json:"_id"`
	UserId     int                `bson:"user_id" json:"user_id,omitempty"`
	Title      string             `bson:"title" json:"title,omitempty"`
	Status     bool               `bson:"status" json:"status,omitempty"`
	CreatedAt  string             `bson:"created_at" json:"created_at"`
	ModifiedAt string             `bson:"modified_at" json:"modified_at"`
}

func NewTaskInfo(req CreateReq) *Task {
	return &Task{
		UserId:     req.UserId,
		Title:      req.Title,
		Status:     false,
		CreatedAt:  time.Now().String(),
		ModifiedAt: time.Now().String(),
	}
}

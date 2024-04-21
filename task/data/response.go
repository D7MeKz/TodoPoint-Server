package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"todopoint/task/data/model"
)

type TaskId struct {
	Id string `json:"task_id"`
}

type SubtaskId struct {
	Id string `json:"subtask_id"`
}

type Subtasks struct {
	Subtasks []primitive.ObjectID `bson:"subtasks"`
}

type TaskInfo struct {
	TaskId    string `json:"task_id" bson:"_id"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	Status    bool   `json:"status" bson:"status"`
}

type TaskDetail struct {
	TaskId    string          `json:"task_id" bson:"_id"`
	CreatedAt string          `json:"created_at" bson:"created_at"`
	Status    bool            `json:"status" bson:"status"`
	Subtasks  []model.SubTask `json:"subtasks"`
}

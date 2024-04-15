package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const TaskNotChecked = false
const TaskChecked = true

const SubTaskNotChecked = false
const SubTaskChecked = true

type Task struct {
	TaskID     primitive.ObjectID `bson:"_id,omitempty" json:"task_id"`
	UserID     int                `bson:"user_id,omitempty" json:"user_id"`
	Status     bool               `bson:"status"`
	CreatedAt  string             `bson:"created_at" json:"created_at"`
	ModifiedAt string             `bson:"modified_at" json:"modified_at"`
	SubTask    []SubTask          `bson:"subtask" json:"subtask"`
}

// NewTask
// Task Constructor
func NewTask(uid int, title string) *Task {
	return &Task{
		UserID:     uid,
		Status:     TaskNotChecked,
		CreatedAt:  time.Now().Format("2006-01-02 15:01"),
		ModifiedAt: time.Now().Format("2006-01-02 15:01"),
	}
}

type SubTask struct {
	SubTaskID primitive.ObjectID `bson:"_id,omitempty" json:"subtask_id"`
	Title     string             `bson:"title" json:"title"`
	Status    bool               `bson:"status" json:"status"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	CheckedAt time.Time          `bson:"checked_at" json:"checked_at"`
	//Record    string             `bson:"record" json:"record"`
}

func NewSubTask(title string) *SubTask {
	return &SubTask{
		Title:     title,
		Status:    SubTaskNotChecked,
		CreatedAt: time.Now(),
	}
}

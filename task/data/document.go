package document

import (
	"time"
	"todopoint/task/data/req"
)

type TaskInfo struct {
	UserId     int
	Title      string
	Status     bool
	CreatedAt  string
	ModifiedAt string
}

func NewTaskInfo(req req.CreateReq) *TaskInfo {
	return &TaskInfo{
		UserId:     req.UserId,
		Title:      req.Title,
		Status:     false,
		CreatedAt:  time.Now().String(),
		ModifiedAt: time.Now().String(),
	}
}

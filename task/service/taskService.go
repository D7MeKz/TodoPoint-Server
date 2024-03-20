package service

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strconv"
	"todopoint/common/db/ent"
	"todopoint/common/networking"
	"todopoint/common/webutils"
	"todopoint/task/data/request"
)

type TaskStore interface {
	Create(ctx *gin.Context, req request.CreateTask) error
	GetAllTasks(ctx *gin.Context, memberId int) ([]*ent.Task, error)
}
type TaskService struct {
	store TaskStore
}

func NewTaskService(s TaskStore) *TaskService {
	return &TaskService{store: s}
}

func isValidMember(baseurl string, userId int) error {
	memberUrl := filepath.Join(baseurl, "members", strconv.Itoa(userId), "valid")
	memberUrl = "http://" + memberUrl
	resp, err := networking.RequestGetToService(memberUrl)
	if err != nil {
		return err
	}
	if resp.StatusCode == 200 {
		return nil
	} else {
		return err
	}
}

func (bs *TaskService) CreateTask(ctx *gin.Context, baseUrl string, req request.CreateTask) (webutils.ErrorType, error) {
	// Check user is valid
	// Panic
	err := isValidMember(baseUrl, req.UserId)
	if err != nil {
		return webutils.ErrorType{Code: webutils.INVALID_MEMBER}, err
	}

	// Create Task
	err = bs.store.Create(ctx, req)
	if err != nil {

		return webutils.ErrorType{Code: webutils.ERROR_TASK_DB}, err
	}
	return webutils.ErrorType{Code: webutils.SUCCESS}, nil
}

func (bs *TaskService) GetTaskByMemId(ctx *gin.Context, memberId int) ([]*ent.Task, webutils.ErrorType, error) {
	// Check user is valid
	err := isValidMember("localhost:3000", memberId)
	if err != nil {
		return nil, webutils.ErrorType{Code: webutils.INVALID_MEMBER}, err
	}

	tasks, err := bs.store.GetAllTasks(ctx, memberId)
	if err != nil {
		return nil, webutils.ErrorType{Code: webutils.ERROR_TASK_DB}, err
	}
	return tasks, webutils.ErrorType{Code: 0}, err
}

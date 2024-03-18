package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strconv"
	"todopoint/common/networking"
	"todopoint/task/data/request"
)

type TaskStore interface {
	Create(ctx *gin.Context, req request.CreateTask) error
}
type TaskService struct {
	store TaskStore
}

func isValidMember(baseurl string, userId int) error {
	memberUrl := filepath.Join(baseurl, "members", strconv.Itoa(userId), "valid")
	resp, err := networking.RequestGetToService(memberUrl)
	if err != nil {
		return err
	}
	if resp.StatusCode == 200 {
		return nil
	} else {
		return errors.New("Member Service Connect Error")
	}

}

func (bs *TaskService) CreateTask(ctx *gin.Context, baseUrl string, req request.CreateTask) error {
	// Check user is valid
	err := isValidMember(baseUrl, req.UserId)
	if err != nil {
		return err
	}

	// Create Task
	err = bs.store.Create(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

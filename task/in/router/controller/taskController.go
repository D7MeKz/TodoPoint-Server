package controller

import (
	"github.com/gin-gonic/gin"
	"todopoint/common/errorutils"
	"todopoint/common/errorutils/codes"
	"todopoint/common/netutils/response"
	"todopoint/task/data"
	"todopoint/task/service"
)

type TaskController struct {
	service service.TaskService
}

func NewTaskController(s service.TaskService) *TaskController {
	return &TaskController{
		service: s,
	}
}

func (c *TaskController) CreateTask(ctx *gin.Context) {
	r := data.CreateReq{}
	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		_ = ctx.Error(errorutils.NewNetError(codes.TaskInvalidJson, err))
		return
	}

	oid, err2 := c.service.CreateTask(ctx, r)
	if err2 != nil {
		_ = ctx.Error(err2)
		return
	}
	response.SuccessWith(ctx, codes.TaskCreationSuccess, oid)
	return
}

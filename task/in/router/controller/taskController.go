package controller

import (
	"github.com/gin-gonic/gin"
	"todopoint/common/auth"
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

// CreateTask
// @Summary Create task
// @Description If user is valid, create task.
// @Tags tasks
// @Accept json
// @Produce json
// @Param request body data.CreateReq true "query params"
// @Success 200 {object} data.TaskId
// @Router /tasks/create [post]
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

// GetList
// @Summary Get Task list
// @Description Get All List
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {object}
// @Router /tasks [get]
func (c *TaskController) GetList(ctx *gin.Context) {
	// Get Token From Header
	uid, tokenErr := auth.GetToken(ctx)
	if tokenErr != nil {
		_ = ctx.Error(tokenErr)
		return
	}
	tasks, err := c.service.GetTasksFrom(ctx, uid)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	response.SuccessWith(ctx, codes.TaskListSuccess, tasks)
}

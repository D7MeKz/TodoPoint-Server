package router

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
// @Summary Create taskA
// @Description If user is valid, create taskA.
// @Tags tasks
// @Accept json
// @Produce json
// @Param request body data.CreateReq true "query params"
// @Success 200 {object} data.TaskId
// @Router /tasks/create [post]
func (c *TaskController) CreateTask(ctx *gin.Context) {
	// Get Token From Header
	uid, tokenErr := auth.GetToken(ctx)
	if tokenErr != nil {
		_ = ctx.Error(tokenErr)
		return
	}

	// Get Request from request
	r := data.CreateReq{}
	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		_ = ctx.Error(errorutils.NewNetError(codes.TaskInvalidJson, err))
		return
	}

	oid, err2 := c.service.CreateTask(ctx, r, uid)
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

// GetToday
// @Summary Get Today's taskA
// @Description Get today's taskA
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {object} data.TaskDetail
// @Router /tasks/today [get]
func (c *TaskController) GetToday(ctx *gin.Context) {
	// Get Token From Header
	uid, tokenErr := auth.GetToken(ctx)
	if tokenErr != nil {
		_ = ctx.Error(tokenErr)
		return
	}

	// Get Today's tasks
	task, err := c.service.GetTodayFrom(ctx, uid)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	response.SuccessWith(ctx, codes.TaskOneSuccess, task)
}

func (c *TaskController) AddSubtask(ctx *gin.Context) {
	uid, tokenErr := auth.GetToken(ctx)
	if tokenErr != nil {
		_ = ctx.Error(tokenErr)
		return
	}

	// Get Request from request
	r := data.AddSubReq{}
	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		_ = ctx.Error(errorutils.NewNetError(codes.SubtaskInvalidJson, err))
		return
	}

	// Add subtask
	tid, err := c.service.AddSubOne(ctx, r, uid)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	response.SuccessWith(ctx, codes.SubtaskOneSuccess, tid)
}

func (c *TaskController) CheckSubtask(ctx *gin.Context) {
	// Get Task Id
	tid, ok := ctx.Params.Get("subtask_id")
	if ok == false {
		_ = ctx.Error(errorutils.NewNetError(codes.TaskInvalidUri, nil))
		return
	}

	// Get query
	status := ctx.Query("checked")
	if status == "" {
		_ = ctx.Error(errorutils.NewNetError(codes.TaskInvalidQuery, nil))
		return
	}

	// Change check status
	ok, err := c.service.CheckSubtask(ctx, tid, status)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	response.Success(ctx, codes.TaskUpdateSuccess)

}

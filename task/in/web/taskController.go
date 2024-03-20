package web

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"todopoint/common/utils"
	"todopoint/common/webutils"
	"todopoint/task/data/request"
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

func (controller *TaskController) CreateTask(ctx *gin.Context) {
	req := request.CreateTask{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		webutils.ErrorFunc(ctx, err, webutils.ErrorType{Code: webutils.ERROR_GO_CODE})
		return
	}

	// Bussiness Logic
	errType, err := controller.service.CreateTask(ctx, "localhost:3000", req)
	if err != nil {
		webutils.ErrorFunc(ctx, err, errType)
		return
	}
	webutils.Success(ctx)

}

func (controller *TaskController) GetTask(ctx *gin.Context) {
	id := ctx.Param("userId")
	// Convert to int
	convertedId, err := strconv.Atoi(id)
	if err != nil {
		webutils.ErrorFunc(ctx, err, webutils.ErrorType{Code: webutils.ERROR_GO_CODE})
		return
	}

	tasks, errType, err := controller.service.GetTaskByMemId(ctx, convertedId)
	if err != nil {
		webutils.ErrorFunc(ctx, err, errType)
	}
	utils.Success(ctx, tasks)
}

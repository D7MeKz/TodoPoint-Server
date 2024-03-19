package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
		utils.Error(ctx, http.StatusBadRequest, err)
		return
	}

	// Bussiness Logic
	err = controller.service.CreateTask(ctx, "localhost:3000", req)
	if err != nil {
		// TODO Error Handling
		utils.Error(ctx, http.StatusInternalServerError, err)
		return
	}
	utils.Success(ctx, nil)
	return

}

func (controller *TaskController) GetTask(ctx *gin.Context) {
	id := ctx.Param("userId")
	// Convert to int
	convertedId, err := strconv.Atoi(id)
	if err != nil {
		webutils.BadRequestError(ctx, err, "Convert Id error")
		return
	}

	tasks, err := controller.service.GetTaskByMemId(ctx, convertedId)
	if err != nil {
		webutils.InternalDBError(ctx, err, "")
		return
	}
	utils.Success(ctx, tasks)
}

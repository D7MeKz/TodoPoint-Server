package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todopoint/common/utils"
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
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, err)
		return
	}

	// Bussiness Logic
	err = controller.service.CreateTask(ctx, "http://localhost:3000", req)
	if err != nil {
		// TODO Error Handling
		utils.Error(ctx, http.StatusInternalServerError, err)
		return
	}

}

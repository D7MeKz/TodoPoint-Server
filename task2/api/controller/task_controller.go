package controller

import (
	"github.com/gin-gonic/gin"
	"todopoint/common2/auth"
	"todopoint/common2/errorutils/codes"
	"todopoint/common2/netutils/response"
	"todopoint/task2/service"
)

type TaskController struct {
	service *service.TaskService
}

func NewTaskController(service service.TaskService) *TaskController {
	return &TaskController{
		service: &service,
	}
}

// Create
// @Summary Create taskA
// @Description If user is valid, create taskA.
// @Tags tasks
// @Accept json
// @Produce json
// @Param request body data.CreateReq true "query params"
// @Success 200 {object} data.TaskId
// @Router /tasks/create [post]
func (t *TaskController) Create(ctx *gin.Context) {
	// Get Token From Header
	uid, tokenErr := auth.GetToken(ctx)
	if tokenErr != nil {
		_ = ctx.Error(tokenErr)
		return
	}

	// TODO Check user is valid

	// Create User
	oid, err := t.service.Create(ctx, uid) // Duck typing
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	response.SuccessWith(ctx, codes.TaskCreationSuccess, oid)
	return
}

func (t *TaskController) GetOne(ctx *gin.Context) {

}

func (t *TaskController) GetMany(ctx *gin.Context) {

}

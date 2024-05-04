package service

import (
	"github.com/gin-gonic/gin"
	"todopoint/common2/errorutils"
	"todopoint/common2/errorutils/codes"
	"todopoint/task2/data"
)

type TaskService struct {
	store Store
}

func (t *TaskService) Create(ctx *gin.Context, data data.DioFormatter) (string, *errorutils.NetError) {
	result, err := t.store.Create(ctx, data)
	if err != nil {
		return "", &errorutils.NetError{Code: codes.TaskCreationError, Err: err}
	}

	oid, err := ConvertOid(result)
	if err != nil {
		return "", &errorutils.NetError{Code: codes.TaskCreationError, Err: err}
	}

	return oid, nil
}

//func (t *TaskService) Delete(ctx *gin.Context, id int) *errorutils.NetError {
//	return nil, nil
//}
//
//func (t *TaskService) Modify(ctx *gin.Context, data interface{}) *errorutils.NetError {
//	return nil, nil
//}
//
//func (t *TaskService) IsExist(ctx *gin.Context, id int) (bool, *errorutils.NetError) {
//	return nil, nil
//}
//
//func (t *TaskService) GetOne(ctx *gin.Context, id int) (interface{}, *errorutils.NetError) {
//	return nil, nil
//}
//
//func (t *TaskService) GetMany(ctx *gin.Context, count int) ([]interface{}, *errorutils.NetError) {
//	tasks, err := t.Store.GetMany(ctx)
//	if err != nil {
//		return nil, &errorutils.NetError{Code: codes.TaskListError, Err: err}
//	}
//
//	return tasks, nil
//}

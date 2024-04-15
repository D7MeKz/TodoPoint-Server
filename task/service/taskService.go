package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"todopoint/common/errorutils"
	"todopoint/common/errorutils/codes"
	"todopoint/task/data"
	"todopoint/task/out/external"
	"todopoint/task/out/persistence"
)

type TaskStore interface {
	Create(ctx *gin.Context, req data.CreateReq) (*mongo.InsertOneResult, error)
	IsExist(ctx *gin.Context, taskId int) (bool, error)
	GetManyFrom(ctx *gin.Context, uid int) ([]data.TaskInfo, error)
}

type TaskService struct {
	Store *persistence.Store
}

func NewTaskService(s *persistence.Store) *TaskService {
	return &TaskService{
		Store: s,
	}
}

func (s *TaskService) CreateTask(ctx *gin.Context, req data.CreateReq) (*data.TaskId, *errorutils.NetError) {
	// Check user isExist
	info, err := external.RequestTo(req.UserId)
	if err != nil {
		logrus.Error(err.Error())
		return nil, &errorutils.NetError{Code: codes.TaskMemberUnavailable, Err: err}
	}

	if !info.IsSuccess() {
		logrus.Error(err.Error())
		return nil, &errorutils.NetError{Code: info.Code, Err: err}
	}

	// Create Task(중복 허용)
	result, err := s.Store.Create(ctx, req)
	if err != nil {
		logrus.Error(err.Error())
		return nil, &errorutils.NetError{Code: codes.TaskCreationError, Err: err}
	}

	// Get oid
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		// Return string
		return &data.TaskId{Id: oid.Hex()}, nil
	}
	return nil, &errorutils.NetError{Code: codes.TaskCreationError, Err: err}
}

// GetTasksFrom
// Get All tasks
// NOTE : Only get 3 tasks, if I learn pagination, I'll change my logic.
func (s *TaskService) GetTasksFrom(ctx *gin.Context, uid int) ([]data.TaskInfo, error) {
	// Get Tasks from users task
	tasks, err := s.Store.GetManyFrom(ctx, uid)
	if err != nil {
		return nil, &errorutils.NetError{Code: codes.TaskListError, Err: err}
	}
	return tasks, nil
}

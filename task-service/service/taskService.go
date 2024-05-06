package service

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"todopoint/common/d7errors"
	"todopoint/common/d7errors/codes"
	"todopoint/task/data"
	"todopoint/task/data/model"
	"todopoint/task/out/persistence"
)

type Store interface {
	Create(ctx *gin.Context, req data.CreateReq) (*mongo.InsertOneResult, error)
	IsExist(ctx *gin.Context, taskId int) (bool, error)
	GetManyFrom(ctx *gin.Context, uid int) ([]data.TaskInfo, error)
}

type TaskService struct {
	Store *persistence.TaskStore
}

func NewTaskService(s *persistence.TaskStore) *TaskService {
	return &TaskService{
		Store: s,
	}
}

func (s *TaskService) GetTodayFrom(ctx *gin.Context, uid int) (*data.TaskDetail, *d7errors.NetError) {
	// Get today's taskA
	today := time.Now().Format("2006-01-02")
	cond := bson.D{{"created_at", today}, {"user_id", uid}}
	task, err := s.Store.GetOneFrom(ctx, cond)
	if err != nil {
		return nil, d7errors.NewNetError(codes.TaskDoesNotFound, err)
	}

	// Get Task detail
	var info = data.TaskInfo{}
	taskErr := task.Decode(&info)
	if taskErr != nil {
		logrus.Errorf("Decoding error : %v", taskErr)
		return nil, d7errors.NewNetError(codes.TaskDecodingErr, taskErr)
	}

	// Get subtask object id
	var subIds = data.Subtasks{}
	idsErr := task.Decode(&subIds)
	if idsErr != nil {
		logrus.Errorf("SubID Decoding error : %v", err)
		return nil, d7errors.NewNetError(codes.TaskDecodingErr, taskErr)
	}

	// Get Subtask info using object id
	var subtasks []model.SubTask
	errFlag := false
	var subtaskErr error
	for _, v := range subIds.Subtasks {
		result, getSubErr := s.Store.GetSubById(ctx, v)
		if getSubErr != nil {
			logrus.Errorf("Get Subtask info error : %v", getSubErr)
			errFlag = true
			subtaskErr = getSubErr
			break
		}
		subtasks = append(subtasks, *result)
	}
	if errFlag == true {
		return nil, d7errors.NewNetError(codes.TaskDecodingErr, subtaskErr)
	}

	// TaskInfo to TaskDetail
	detail := data.TaskDetail{TaskId: info.TaskId, CreatedAt: info.CreatedAt, Status: info.Status, Subtasks: subtasks}
	return &detail, nil

}
func (s *TaskService) CreateTask(ctx *gin.Context, req data.CreateReq, uid int) (*data.TaskId, *d7errors.NetError) {
	// Create Task
	result, err := s.Store.Create(ctx, req, uid)
	if err != nil {
		logrus.Error(err.Error())
		return nil, &d7errors.NetError{Code: codes.TaskCreationError, Err: err}
	}

	// Get oid
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		// Return string
		return &data.TaskId{Id: oid.Hex()}, nil
	}
	return nil, &d7errors.NetError{Code: codes.TaskCreationError, Err: err}
}

// GetTasksFrom
// Get All tasks
// NOTE : Only get 3 tasks, if I learn pagination, I'll change my logic.
func (s *TaskService) GetTasksFrom(ctx *gin.Context, uid int) ([]data.TaskInfo, error) {
	// Get Tasks from users taskA
	tasks, err := s.Store.GetManyFrom(ctx, uid)
	if err != nil {
		return nil, &d7errors.NetError{Code: codes.TaskListError, Err: err}
	}
	return tasks, nil
}

func (s *TaskService) AddSubOne(ctx *gin.Context, req data.AddSubReq, uid int) (*data.SubtaskId, error) {
	// Create Subtask
	result, err := s.Store.CreateSubtask(ctx, req.Title)
	if err != nil {
		return nil, &d7errors.NetError{Code: codes.SubtaskCreationErr, Err: err}
	}
	// Update subtask into taskA
	subOid, subOk := result.InsertedID.(primitive.ObjectID)
	if subOk {
		addOk, addErr := s.Store.Add(ctx, req.TaskId, subOid)
		if addErr != nil && !addOk {
			logrus.Errorf("Can't add subtask in taskA : %v", addErr)
			return nil, &d7errors.NetError{Code: codes.SubtaskAdditionErr, Err: err}
		}

		// return subtask id
		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			// Return string
			return &data.SubtaskId{Id: oid.Hex()}, nil
		}
		return nil, &d7errors.NetError{Code: codes.SubtaskCreationErr, Err: err}
	}
	return nil, &d7errors.NetError{Code: codes.SubtaskAdditionErr, Err: err}

}

func (s *TaskService) CheckSubtask(ctx *gin.Context, stid string, status string) (bool, *d7errors.NetError) {
	// if checkSubtask is exist
	stoid, err := primitive.ObjectIDFromHex(stid)
	if err != nil {
		return false, d7errors.NewNetError(codes.TaskDecodingErr, err)
	}

	// Update status
	ok, err := s.Store.UpdateStatus(ctx, stoid, status)
	if !ok && err != nil {
		return false, d7errors.NewNetError(codes.SubtaskUpdateErr, err)
	}
	return true, nil
}

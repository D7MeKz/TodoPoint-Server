package persistence

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"todopoint/task/data"
	"todopoint/task/data/model"
)

type TaskStore struct {
	client       *mongo.Client
	tCollection  *mongo.Collection
	stCollection *mongo.Collection
}

func (s *TaskStore) IsExist(ctx *gin.Context, taskId int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewTaskStore(client *mongo.Client) *TaskStore {
	return &TaskStore{
		client:       client,
		tCollection:  client.Database("todopoint").Collection("task"),
		stCollection: client.Database("todopoint").Collection("subtask"),
	}
}

func (s *TaskStore) Create(ctx *gin.Context, req data.CreateReq, uid int) (*mongo.InsertOneResult, error) {
	taskData := model.NewTask(uid, req.Title)
	id, err := s.tCollection.InsertOne(ctx, taskData)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (s *TaskStore) GetOneFrom(ctx *gin.Context, filter bson.D) (*mongo.SingleResult, error) {
	// Check taskA is exist
	result := s.tCollection.FindOne(ctx, filter)
	if result.Err() != nil {
		mongoErr := result.Err()
		logrus.Errorf("MongoDB Error : %v", mongoErr)
		return nil, mongoErr
	}
	return result, nil
}

func (s *TaskStore) GetSubById(ctx *gin.Context, oid primitive.ObjectID) (*model.SubTask, error) {
	result := s.stCollection.FindOne(ctx, bson.D{{"_id", oid}})

	// Decode to Subtask
	var subtask = model.SubTask{}
	err := result.Decode(&subtask)
	if err != nil {
		return nil, err
	}

	return &subtask, nil
}
func (s *TaskStore) GetManyFrom(ctx *gin.Context, uid int) ([]data.TaskInfo, error) {
	// Get Cursor
	filter := bson.D{{"user_id", uid}}
	cursor, findErr := s.tCollection.Find(ctx, filter)
	if findErr != nil {
		logrus.Fatal(findErr)
		return nil, findErr
	}

	// Close cursor
	defer cursor.Close(ctx)

	// Get data.TaskInfo from data.Task
	var tasks []data.TaskInfo
	var cnt int = 0
	for cursor.Next(ctx) {
		var task data.TaskInfo
		taskErr := cursor.Decode(&task)

		if taskErr != nil {
			logrus.Fatal(taskErr)
		}
		tasks = append(tasks, task)

		// Only 3 tasks
		cnt++
		if cnt == 2 {
			break
		}
	}

	// Cursor error handling
	// Err returns the last error seen by the Cursor, or nil if no error has occurred.
	if cursorErr := cursor.Err(); cursorErr != nil {
		return nil, cursorErr
	}

	return tasks, nil

}

func (s *TaskStore) CreateSubtask(ctx *gin.Context, title string) (*mongo.InsertOneResult, error) {
	// Create subtask
	subtask := model.NewSubTask(title)
	result, err := s.stCollection.InsertOne(ctx, subtask)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Add
// Update Subtask into taskA
func (s *TaskStore) Add(ctx *gin.Context, tid string, subId primitive.ObjectID) (bool, error) {
	toid, convertErr := primitive.ObjectIDFromHex(tid)
	if convertErr != nil {
		return false, convertErr
	}

	_, err := s.tCollection.UpdateOne(
		ctx,
		bson.M{"_id": toid}, // String to ObjectID
		bson.M{"$push": bson.M{"subtasks": subId}},
	)
	if err != nil {
		return false, err
	}
	return true, err
}

func (s *TaskStore) UpdateStatus(ctx *gin.Context, subId primitive.ObjectID, status string) (bool, error) {
	// Change status type, int to bool
	var st bool
	st, err := strconv.ParseBool(status)
	if err != nil {
		return false, err
	}

	// Update value
	_, err = s.stCollection.UpdateOne(
		ctx,
		bson.M{"_id": subId}, // String to ObjectID
		bson.M{"$set": bson.M{"status": st}},
	)
	if err != nil {
		return false, err
	}
	return true, nil
}

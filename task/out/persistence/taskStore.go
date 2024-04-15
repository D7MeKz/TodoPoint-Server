package persistence

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"todopoint/common/config/mongodb"
	"todopoint/task/data"
)

type Store struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (s *Store) IsExist(ctx *gin.Context, taskId int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewStore(colName string) *Store {
	client := mongodb.GetClient()
	return &Store{
		client:     client,
		collection: mongodb.GetCollection(client, colName),
	}
}

func (s *Store) Create(ctx *gin.Context, req data.CreateReq) (*mongo.InsertOneResult, error) {
	taskData := data.NewTask(req)
	id, err := s.collection.InsertOne(ctx, taskData)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (s *Store) GetManyFrom(ctx *gin.Context, uid int) ([]data.TaskInfo, error) {
	// Get Cursor
	filter := bson.D{{"user_id", uid}}
	cursor, findErr := s.collection.Find(ctx, filter)
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

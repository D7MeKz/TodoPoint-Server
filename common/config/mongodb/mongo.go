package mongodb

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"path/filepath"
)

var (
	client *mongo.Client
)

func GetClient() *mongo.Client { return client }

func SetClient(newClient *mongo.Client) {
	client = newClient
}

func GetCollection(client *mongo.Client, colName string) *mongo.Collection {
	return client.Database("todopoint").Collection(colName)
}

func NewMongoClient(path string) *mongo.Client {
	dir, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}

	environmentPath := filepath.Join(dir, ".env")
	err = godotenv.Load(environmentPath)
	if err != nil {
		log.Fatal("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	username := os.Getenv("MONGODB_USERNAME")
	pw := os.Getenv("MONGODB_PASSWORD")
	if uri == "" || username == "" || pw == "" {
		log.Fatal("Empty .env")
	}

	client, err := mongo.Connect(context.TODO(),
		options.Client().ApplyURI(uri).SetAuth(options.Credential{Username: username, Password: pw}))
	if err != nil {
		panic(err)
	}

	//defer func() {
	//	if err := client.Disconnect(context.TODO()); err != nil {
	//		panic(err)
	//	}
	//}()

	return client
}

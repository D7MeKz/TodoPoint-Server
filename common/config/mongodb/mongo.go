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
	if uri == "" {
		log.Fatal("You must set uri")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	return client
}

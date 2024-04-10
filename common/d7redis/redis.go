package d7redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
)

func GetClient() *redis.Client { return client }

func SetClient(newClient *redis.Client) {
	client = newClient
}

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}
	return client
}

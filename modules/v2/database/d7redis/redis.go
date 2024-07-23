package d7redis

import "github.com/redis/go-redis/v9"

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(addr string, password string) *RedisClient {
	return &RedisClient{
		Client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       0,
		}),
	}
}

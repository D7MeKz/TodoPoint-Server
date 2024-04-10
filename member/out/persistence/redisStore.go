package persistence

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
	"todopoint/common/d7redis"
)

type RedisStore struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisStore() *RedisStore {
	client := d7redis.NewRedisClient()
	return &RedisStore{
		client: client,
		ctx:    context.Background(),
	}
}

func (s *RedisStore) Create(key string, value any, expires int64) error {
	rt := time.Unix(expires, 0) // Convert TO UTC
	err := s.client.Set(s.ctx, key, value, rt.Sub(time.Now())).Err()
	if err != nil {
		return err
	}
	return nil
}

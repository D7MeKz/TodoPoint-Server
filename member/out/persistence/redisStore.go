package persistence

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
	"todopoint/common2/d7redis"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore() *RedisStore {
	return &RedisStore{
		client: d7redis.GetClient(),
	}


func (s *RedisStore) Create(ctx *gin.Context, key string, value string, expires int64) error {
	rt := time.Unix(expires, 0) // Convert TO UTC
	err := s.client.Set(ctx.Request.Context(), key, value, rt.Sub(time.Now())).Err()
	if err != nil {
		logrus.Warn(err)
		panic(err)
	}
	return nil
}

func (s *RedisStore) Find(ctx *gin.Context, key string) (int, error) {
	memId, err := s.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return 0, err
	} else if err != nil {
		panic(err)
	} else {
		converted, err2 := strconv.Atoi(memId)
		if err2 != nil {
			logrus.Error(err2)
			return -1, err2
		}
		return converted, nil
	}
}

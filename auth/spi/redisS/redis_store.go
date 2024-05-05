package redisS

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"time"
	"todopoint/auth/data/params"
	"todopoint/common/database/d7redis"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore() *RedisStore {
	return &RedisStore{
		client: d7redis.NewRedisClient().Client,
	}
}

func (s *RedisStore) Create(ctx *gin.Context, data interface{}) error {
	params, ok := data.(params.RedisParams)
	if !ok {
		return errors.New("Invalid data type")
	}

	rt := time.Unix(params.Expires, 0) // Convert TO UTC
	err := s.client.Set(ctx.Request.Context(), params.Key, params.Value, rt.Sub(time.Now())).Err()
	if err != nil {
		logrus.Warn(err)
		panic(err)
	}
	return nil
}

func (s *RedisStore) IsExist(ctx *gin.Context, data interface{}) (bool, error) {
	key, ok := data.(string)
	if !ok {
		return false, errors.New("Invalid data type")
	}

	_, err := s.client.Get(ctx.Request.Context(), key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, err
		}
		return false, err
	}
	return true, nil
}

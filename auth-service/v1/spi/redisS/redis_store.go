package redisS

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"modules/d7redis"
	"time"
	"todopoint/auth/data/params"
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
		return errors.New("Invalid httpdata type")
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
		return false, errors.New("Invalid httpdata type")
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

// Modify modifies refresh token in redis through the user id.
func (s *RedisStore) Modify(ctx *gin.Context, data interface{}) error {
	// check the httpdata type
	d, ok := data.(params.RedisParams)
	if !ok {
		return errors.New("Invalid httpdata type")
	}

	// modify
	err := s.client.Set(ctx, d.Key, d.Value, getExpiredTime()).Err()
	if err != nil {
		return err
	}
	return nil
}

func getExpiredTime() time.Duration {
	rt := time.Unix(time.Now().Add(time.Hour*24*7).Unix(), 0) // Convert TO UTC
	return rt.Sub(time.Now())
}

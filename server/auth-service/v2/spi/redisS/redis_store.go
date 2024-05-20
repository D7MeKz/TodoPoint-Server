package redisS

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"time"
	"todopoint/auth/v2/data"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(client *redis.Client) *RedisStore {
	return &RedisStore{
		client: client,
	}
}

func (s *RedisStore) Create(ctx *gin.Context, d interface{}) error {
	params, ok := d.(data.RedisParams)
	if !ok {
		return errors.New("Invalid httpdata type")
	}

	rt := time.Unix(params.Expires, 0) // Convert TO UTC
	err := s.client.Set(ctx.Request.Context(), params.Key, params.Value, rt.Sub(time.Now())).Err()
	if err != nil {
		logrus.Warn(err)
		return err
	}
	return nil
}

func (s *RedisStore) IsExist(ctx *gin.Context, d interface{}) (bool, error) {
	key, ok := d.(string)
	if !ok {
		return false, errors.New("invalid http data type")
	}

	_, err := s.client.Get(ctx.Request.Context(), key).Result()
	if errors.Is(err, redis.Nil) {
		return false, errors.New("key not found in redis")
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// Modify modifies refresh token in redis through the user id.
func (s *RedisStore) Modify(ctx *gin.Context, d interface{}) error {
	// check the httpdata type
	param, ok := d.(data.RedisParams)
	if !ok {
		return errors.New("Invalid httpdata type")
	}

	// modify
	value, err := s.client.Set(ctx, param.Key, param.Value, getExpiredTime()).Result()
	logrus.Debugf("Modify : %v", value)
	if errors.Is(err, redis.Nil) {
		return errors.New("key not found in redis")
	} else if err != nil {
		return err
	}
	return nil
}

func getExpiredTime() time.Duration {
	rt := time.Unix(time.Now().Add(time.Hour*24*7).Unix(), 0) // Convert TO UTC
	return rt.Sub(time.Now())
}

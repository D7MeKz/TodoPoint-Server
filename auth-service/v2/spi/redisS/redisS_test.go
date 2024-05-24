package redisS_test

import (
	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"modules/common/testutils"
	"testing"
	"todopoint/auth/v2/data"
	"todopoint/auth/v2/spi/redisS"
)

func TestRedisStore_Create(t *testing.T) {
	// setup
	db, _ := redismock.NewClientMock()
	client := redisS.NewRedisStore(db)
	ctx := testutils.GetTestGinContext()

	tests := []struct {
		description string
		input       interface{}
		isError     bool
	}{
		{
			description: "Success : Create a new redis key",
			input:       data.RedisParams{Key: "key", Value: "value", Expires: 100},
			isError:     false,
		},
		{
			description: "Fail#1 : Empty Key",
			input:       data.RedisParams{Key: "", Value: "value", Expires: 100},
			isError:     true,
		},
		{
			description: "Fail#2 : Empty Value",
			input:       data.RedisParams{Key: "key", Value: "", Expires: 100},
			isError:     true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			err := client.Create(ctx, tc.input)
			if tc.isError {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestRedisStore_IsExist(t *testing.T) {
	// setup
	db, mock := redismock.NewClientMock()
	client := redisS.NewRedisStore(db)
	defer func(db *redis.Client) {
		err := db.Close()
		if err != nil {
			t.Log(err)
		}
	}(db)
	ctx := testutils.GetTestGinContext()

	tests := []struct {
		description string
		input       interface{}
		isExist     bool
	}{
		{
			description: "Success : Key Exist",
			input:       "key",
			isExist:     true,
		},
		{
			description: "Fail#1 : Invalid Key",
			input:       "key1",
			isExist:     false,
		},
		{
			description: "Fail#2 : Empty Key",
			input:       "",
			isExist:     false,
		},
	}

	// expected
	mock.ExpectGet("key").SetVal("value")

	// Testing
	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			_, _ = client.IsExist(ctx, tc.input)
			mock.ExpectGet(tc.input.(string))
		})
	}
}

func TestRedisStore_Modify(t *testing.T) {
	// setup
	db, mock := redismock.NewClientMock()
	client := redisS.NewRedisStore(db)
	ctx := testutils.GetTestGinContext()

	tests := []struct {
		description string
		input       interface{}
	}{
		{
			description: "Success : Modify a new redis key",
			input:       data.RedisParams{Key: "key", Value: "modified", Expires: 100},
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			_ = client.Modify(ctx, tc.input)
			mock.ExpectSet("key", "value", 0).SetVal("modified")
		})
	}

}

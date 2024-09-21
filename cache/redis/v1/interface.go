package redis

import (
	"time"

	"golang.org/x/net/context"
)

type RedisClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
	Remove(ctx context.Context, key string) error
	RemoveAll(ctx context.Context, keys ...string) error
	GetInt(ctx context.Context, key string) (int64, error)
	Expire(ctx context.Context, key string, expiration time.Duration) error
	Incr(ctx context.Context, key string) (int64, error)
	Exist(ctx context.Context, key string) (int64, error)
	IncrBy(ctx context.Context, key string, val int64) (int64, error)
	HSet(ctx context.Context, key string, field string, value interface{}) error
	HGet(ctx context.Context, key string, field string) (string, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	HDel(ctx context.Context, key string, fields ...string) error
	HExists(ctx context.Context, key string, field string) (bool, error)
	HLen(ctx context.Context, key string) (int64, error)
	HIncrBy(ctx context.Context, key string, field string, value int64) (int64, error)
	HIncrByFloat(ctx context.Context, key string, field string, value float64) (float64, error)
	HKeys(ctx context.Context, key string) ([]string, error)
	HVals(ctx context.Context, key string) ([]string, error)
	HSetNX(ctx context.Context, key string, field string, value interface{}) (bool, error)
	HMSet(ctx context.Context, key string, fields map[string]interface{}) error
	RPush(ctx context.Context, key string, value interface{}) error
	LRange(ctx context.Context, key string, start, stop int64) ([]string, error)
}

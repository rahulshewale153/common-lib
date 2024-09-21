package redis

import (
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

type redisConnector struct {
	connection *redis.Client
}

func NewRedisClient(host string, port int, password string, db, poolSize int) (RedisClient, error) {
	opt := new(redis.Options)
	opt.Addr = fmt.Sprintf("%s:%d", host, port)
	opt.Password = password
	opt.DB = db
	opt.PoolSize = poolSize
	c := redis.NewClient(opt)
	redisConn := redisConnector{
		connection: c,
	}
	if _, err := c.Ping(context.Background()).Result(); nil != err {
		return &redisConn, err
	}

	return &redisConn, nil
}

func (r *redisConnector) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.connection.Set(ctx, key, value, expiration).Err()
}

func (r *redisConnector) Get(ctx context.Context, key string) (string, error) {
	return r.connection.Get(ctx, key).Result()
}

func (r *redisConnector) Delete(ctx context.Context, key string) error {
	return r.connection.Del(ctx, key).Err()
}

func (r *redisConnector) Remove(ctx context.Context, key string) error {
	return r.connection.Del(ctx, key).Err()
}

func (r *redisConnector) RemoveAll(ctx context.Context, keys ...string) error {
	return r.connection.Del(ctx, keys...).Err()
}

func (r *redisConnector) GetInt(ctx context.Context, key string) (int64, error) {
	str, err := r.connection.Get(ctx, key).Result()
	if nil != err {
		return 0, err
	}
	return strconv.ParseInt(str, 10, 64)
}

func (r *redisConnector) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return r.connection.Expire(ctx, key, expiration).Err()
}

func (r *redisConnector) Incr(ctx context.Context, key string) (int64, error) {
	return r.connection.Incr(ctx, key).Result()
}

func (r *redisConnector) Exist(ctx context.Context, key string) (int64, error) {
	return r.connection.Exists(ctx, key).Result()
}

func (r *redisConnector) IncrBy(ctx context.Context, key string, val int64) (int64, error) {
	return r.connection.IncrBy(ctx, key, val).Result()
}

// Redis hash functions
func (r *redisConnector) HSet(ctx context.Context, key string, field string, value interface{}) error {
	return r.connection.HSet(ctx, key, field, value).Err()
}

func (r *redisConnector) HGet(ctx context.Context, key string, field string) (string, error) {
	return r.connection.HGet(ctx, key, field).Result()
}

func (r *redisConnector) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return r.connection.HGetAll(ctx, key).Result()
}

func (r *redisConnector) HDel(ctx context.Context, key string, fields ...string) error {
	return r.connection.HDel(ctx, key, fields...).Err()
}

func (r *redisConnector) HExists(ctx context.Context, key string, field string) (bool, error) {
	return r.connection.HExists(ctx, key, field).Result()
}

func (r *redisConnector) HLen(ctx context.Context, key string) (int64, error) {
	return r.connection.HLen(ctx, key).Result()
}

func (r *redisConnector) HIncrBy(ctx context.Context, key string, field string, value int64) (int64, error) {
	return r.connection.HIncrBy(ctx, key, field, value).Result()
}

func (r *redisConnector) HIncrByFloat(ctx context.Context, key string, field string, value float64) (float64, error) {
	return r.connection.HIncrByFloat(ctx, key, field, value).Result()
}

func (r *redisConnector) HKeys(ctx context.Context, key string) ([]string, error) {
	return r.connection.HKeys(ctx, key).Result()
}

func (r *redisConnector) HVals(ctx context.Context, key string) ([]string, error) {
	return r.connection.HVals(ctx, key).Result()
}

func (r *redisConnector) HSetNX(ctx context.Context, key string, field string, value interface{}) (bool, error) {
	return r.connection.HSetNX(ctx, key, field, value).Result()
}

func (r *redisConnector) HMSet(ctx context.Context, key string, fields map[string]interface{}) error {
	return r.connection.HMSet(ctx, key, fields).Err()
}

func (r *redisConnector) RPush(ctx context.Context, key string, value interface{}) error {
	return r.connection.RPush(ctx, key, value).Err()
}

func (r *redisConnector) LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return r.connection.LRange(ctx, key, start, stop).Result()
}

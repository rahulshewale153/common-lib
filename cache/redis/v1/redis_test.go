package redis

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

var redisOpt RedisClient
var rc *redis.Client

func setup(t *testing.T) {
	redisServer := mockRedis(t)
	rc = redis.NewClient(&redis.Options{
		Addr: redisServer.Addr(),
	})
	redisOpt = &redisConnector{connection: rc}
}

func teardown() {
	rc.Close()
}

// 1 and 2
func mockRedis(t *testing.T) *miniredis.Miniredis {
	s, err := miniredis.Run()
	assert.NoError(t, err)
	return s
}

// testing NewRedisClient func in redis.go
func TestNewRedisClient(t *testing.T) {
	host := ""
	port := 0
	password := ""
	db := 0
	poolSize := 10

	_, err := NewRedisClient(host, port, password, db, poolSize)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "connection refused")
}

func TestRedisClientSetAndGet(t *testing.T) {
	setup(t)
	defer teardown()
	errSet := redisOpt.Set(context.Background(), "test_key", "test_value", 10*time.Minute)
	assert.NoError(t, errSet)

	value, errGet := redisOpt.Get(context.Background(), "test_key")
	assert.NoError(t, errGet, value)

}

func TestRedisDelete(t *testing.T) {
	setup(t)
	defer teardown()

	key1 := "test_key2"
	value1 := "test_value2"
	expiration := time.Minute

	err := redisOpt.Set(context.Background(), key1, value1, expiration)
	assert.NoError(t, err)

	err1 := redisOpt.Delete(context.Background(), key1)
	assert.NoError(t, err1)

	value, err2 := redisOpt.Get(context.Background(), key1)
	assert.Error(t, err2)
	assert.EqualError(t, err2, "redis: nil")
	assert.Empty(t, value)

}

func TestRedisRemove(t *testing.T) {
	setup(t)
	defer teardown()

	err := redisOpt.Set(context.Background(), "test_key", "test_value", 10*time.Minute)
	assert.NoError(t, err)

	err = redisOpt.Remove(context.Background(), "test_key")
	assert.NoError(t, err)

	value, err := redisOpt.Get(context.Background(), "test_key")
	assert.Error(t, err)
	assert.EqualError(t, err, "redis: nil")
	assert.Empty(t, value)
}

func TestRedisRemoveAll(t *testing.T) {
	setup(t)
	defer teardown()

	err := redisOpt.Set(context.Background(), "key1", "value1", 10*time.Minute)
	assert.NoError(t, err)
	err = redisOpt.Set(context.Background(), "key2", "value2", 10*time.Minute)
	assert.NoError(t, err)

	err = redisOpt.RemoveAll(context.Background(), "key1", "key2")
	assert.NoError(t, err)

	_, err = redisOpt.Get(context.Background(), "key1")
	assert.Error(t, err)
	assert.EqualError(t, err, "redis: nil")

	_, err = redisOpt.Get(context.Background(), "key2")
	assert.Error(t, err)
	assert.EqualError(t, err, "redis: nil")
}

func TestRedisGetInt(t *testing.T) {
	setup(t)
	defer teardown()

	// Test case 1: Valid integer value
	err := redisOpt.Set(context.Background(), "int_key", "123", time.Minute)
	assert.NoError(t, err)

	val, err := redisOpt.GetInt(context.Background(), "int_key")
	assert.NoError(t, err)
	assert.Equal(t, int64(123), val)

	// Test case 2: Non-existent key
	_, err = redisOpt.GetInt(context.Background(), "non_existing_key")
	assert.Error(t, err)
	assert.EqualError(t, err, "redis: nil")

	// Test case 3: Invalid integer value
	err = redisOpt.Set(context.Background(), "invalid_int_key", "abc", time.Minute)
	assert.NoError(t, err)

	_, err = redisOpt.GetInt(context.Background(), "invalid_int_key")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "strconv.ParseInt")

	// Test case 4: Empty string value
	err = redisOpt.Set(context.Background(), "empty_key", "", time.Minute)
	assert.NoError(t, err)

	_, err = redisOpt.GetInt(context.Background(), "empty_key")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "strconv.ParseInt: parsing \"\": invalid syntax")
}

func TestRedisExpire(t *testing.T) {
	setup(t)
	defer teardown()

	err := redisOpt.Set(context.Background(), "exp_key", "value", time.Minute)
	assert.NoError(t, err)

	err = redisOpt.Expire(context.Background(), "exp_key", 1*time.Hour)
	assert.NoError(t, err)

	ttl, err := rc.TTL(context.Background(), "exp_key").Result()
	assert.NoError(t, err)
	assert.True(t, ttl > 0 && ttl <= 1*time.Hour)
}

func TestRedisIncr(t *testing.T) {
	setup(t)
	defer teardown()

	err := redisOpt.Set(context.Background(), "count", "1", 10*time.Minute)
	assert.NoError(t, err)

	val, err := redisOpt.Incr(context.Background(), "count")
	assert.NoError(t, err)
	assert.Equal(t, int64(2), val)

	val, err = redisOpt.Incr(context.Background(), "count")
	assert.NoError(t, err)
	assert.Equal(t, int64(3), val)
}

func TestRedisExist(t *testing.T) {
	setup(t)
	defer teardown()

	err := redisOpt.Set(context.Background(), "check_key", "value", 10*time.Minute)
	assert.NoError(t, err)

	exists, err := redisOpt.Exist(context.Background(), "check_key")
	assert.NoError(t, err)
	assert.Equal(t, int64(1), exists)

	exists, err = redisOpt.Exist(context.Background(), "non_existing_key")
	assert.NoError(t, err)
	assert.Equal(t, int64(0), exists)
}

func TestRedisIncrBy(t *testing.T) {
	setup(t)
	defer teardown()

	err := redisOpt.Set(context.Background(), "count", "1", 10*time.Minute)
	assert.NoError(t, err)

	val, err := redisOpt.IncrBy(context.Background(), "count", 5)
	assert.NoError(t, err)
	assert.Equal(t, int64(6), val)

	val, err = redisOpt.IncrBy(context.Background(), "count", -2)
	assert.NoError(t, err)
	assert.Equal(t, int64(4), val)
}

func TestRedisHSetAndGet(t *testing.T) {
	setup(t)
	defer teardown()

	err := redisOpt.HSet(context.Background(), "hash_key", "field1", "value1")
	assert.NoError(t, err)

	val, err := redisOpt.HGet(context.Background(), "hash_key", "field1")
	assert.NoError(t, err)
	assert.Equal(t, "value1", val)
}

func TestRedisHGetAll(t *testing.T) {
	setup(t)
	defer teardown()

	fields := map[string]interface{}{
		"field1": "value1",
		"field2": "value2",
	}
	err := redisOpt.HMSet(context.Background(), "hash_key", fields)
	assert.NoError(t, err)

	allValues, err := redisOpt.HGetAll(context.Background(), "hash_key")
	assert.NoError(t, err)
	assert.Len(t, allValues, 2)
	assert.Equal(t, "value1", allValues["field1"])
	assert.Equal(t, "value2", allValues["field2"])
}

func TestRedisHDel(t *testing.T) {
	setup(t)
	defer teardown()

	err := redisOpt.HSet(context.Background(), "hash_key", "field1", "value1")
	assert.NoError(t, err)

	err = redisOpt.HSet(context.Background(), "hash_key", "field2", "value2")
	assert.NoError(t, err)

	err = redisOpt.HDel(context.Background(), "hash_key", "field1", "field2")
	assert.NoError(t, err)

	val, err := redisOpt.HGet(context.Background(), "hash_key", "field1")
	assert.Error(t, err) // Expecting an error since field1 should be deleted
	assert.EqualError(t, err, "redis: nil")
	assert.Empty(t, val)

	val, err = redisOpt.HGet(context.Background(), "hash_key", "field2")
	assert.Error(t, err) // Expecting an error since field2 should be deleted
	assert.EqualError(t, err, "redis: nil")
	assert.Empty(t, val)
}

func TestRedisHExists(t *testing.T) {
	setup(t)
	defer teardown()

	err := redisOpt.HSet(context.Background(), "hash_key", "field1", "value1")
	assert.NoError(t, err)

	exists, err := redisOpt.HExists(context.Background(), "hash_key", "field1")
	assert.NoError(t, err)
	assert.True(t, exists)

	exists, err = redisOpt.HExists(context.Background(), "hash_key", "field2")
	assert.NoError(t, err)
	assert.False(t, exists)
}

func TestRedisHLen(t *testing.T) {
	setup(t)
	defer teardown()

	fields := map[string]interface{}{
		"field1": "value1",
		"field2": "value2",
	}
	err := redisOpt.HMSet(context.Background(), "hash_key", fields)
	assert.NoError(t, err)

	len, err := redisOpt.HLen(context.Background(), "hash_key")
	assert.NoError(t, err)
	assert.Equal(t, int64(2), len)
}

func TestRedisHIncrBy(t *testing.T) {
	setup(t)
	defer teardown()

	err := redisOpt.HSet(context.Background(), "hash_key", "field1", "10")
	assert.NoError(t, err)

	val, err := redisOpt.HIncrBy(context.Background(), "hash_key", "field1", 5)
	assert.NoError(t, err)
	assert.Equal(t, int64(15), val)
}

func TestRedisHIncrByFloat(t *testing.T) {
	setup(t)
	defer teardown()

	err := redisOpt.HSet(context.Background(), "hash_key", "field1", "10.5")
	assert.NoError(t, err)

	val, err := redisOpt.HIncrByFloat(context.Background(), "hash_key", "field1", 1.5)
	assert.NoError(t, err)
	assert.Equal(t, float64(12.0), val)
}
func TestRedisHKeys(t *testing.T) {
	setup(t)
	defer teardown()

	fields := map[string]interface{}{
		"field1": "value1",
		"field2": "value2",
	}
	err := redisOpt.HMSet(context.Background(), "hash_key", fields)
	assert.NoError(t, err)

	keys, err := redisOpt.HKeys(context.Background(), "hash_key")
	assert.NoError(t, err)
	assert.ElementsMatch(t, []string{"field1", "field2"}, keys)
}

func TestRedisHVals(t *testing.T) {
	setup(t)
	defer teardown()

	fields := map[string]interface{}{
		"field1": "value1",
		"field2": "value2",
	}
	err := redisOpt.HMSet(context.Background(), "hash_key", fields)
	assert.NoError(t, err)

	vals, err := redisOpt.HVals(context.Background(), "hash_key")
	assert.NoError(t, err)
	assert.ElementsMatch(t, []string{"value1", "value2"}, vals)
}

func TestRedisHSetNX(t *testing.T) {
	setup(t)
	defer teardown()

	ok, err := redisOpt.HSetNX(context.Background(), "hash_key", "field1", "value1")
	assert.NoError(t, err)
	assert.True(t, ok)

	ok, err = redisOpt.HSetNX(context.Background(), "hash_key", "field1", "new_value")
	assert.NoError(t, err)
	assert.False(t, ok)

	val, err := redisOpt.HGet(context.Background(), "hash_key", "field1")
	assert.NoError(t, err)
	assert.Equal(t, "value1", val)
}

func TestRedisHMSet(t *testing.T) {
	setup(t)
	defer teardown()

	fields := map[string]interface{}{
		"field1": "value1",
		"field2": "value2",
	}
	err := redisOpt.HMSet(context.Background(), "hash_key", fields)
	assert.NoError(t, err)

	val1, err := redisOpt.HGet(context.Background(), "hash_key", "field1")
	assert.NoError(t, err)
	assert.Equal(t, "value1", val1)

	val2, err := redisOpt.HGet(context.Background(), "hash_key", "field2")
	assert.NoError(t, err)
	assert.Equal(t, "value2", val2)
}

func TestRedisRPushAndLRange(t *testing.T) {
	setup(t)
	defer teardown()

	err := redisOpt.RPush(context.Background(), "list_key", "value1")
	assert.NoError(t, err)

	err = redisOpt.RPush(context.Background(), "list_key", "value2")
	assert.NoError(t, err)

	values, err := redisOpt.LRange(context.Background(), "list_key", 0, -1)
	assert.NoError(t, err)
	assert.Equal(t, []string{"value1", "value2"}, values)
}

func TestRedisSetWithExpiration(t *testing.T) {
	setup(t)
	defer teardown()

	err := redisOpt.Set(context.Background(), "exp_key", "value", time.Minute)
	assert.NoError(t, err)

	ttl, err := rc.TTL(context.Background(), "exp_key").Result()
	assert.NoError(t, err)
	assert.True(t, ttl > 0 && ttl <= time.Minute)
}

func TestRedisHGetInvalidField(t *testing.T) {
	setup(t)
	defer teardown()

	_, err := redisOpt.HGet(context.Background(), "hash_key", "non_existing_field")
	assert.Error(t, err)
	assert.EqualError(t, err, "redis: nil")
}

package main

import (
	"context"
	"fmt"
	"time"

	gocache "github.com/rahulshewale153/common-lib/cache/gocache"
	redis "github.com/rahulshewale153/common-lib/cache/redis/v1"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	// initialise the cache of gocache and redis
	gocache.NewGoCache(1*time.Second, 1*time.Second)

	gocache.Set("foo", "bar", 2*time.Second)
	value, found := gocache.Get("foo")
	if found {
		fmt.Println(value)
	} else {
		fmt.Println("Key not found")
	}

	redisClient, err := redis.NewRedisClient("localhost", 6379, "", 0, 0)
	if err != nil {
		fmt.Printf("Error creating Redis client: %v\n", err)
		return
	}

	err = redisClient.Set(context.Background(), "key2", "value2", time.Minute)
	if err != nil {
		fmt.Printf("Error setting value in Redis: %v\n", err)
		return
	}

	val, err := redisClient.Get(context.Background(), "key2")
	if err != nil {
		fmt.Printf("Error getting value from Redis: %v\n", err)
		return
	}
	fmt.Println(val, "this is val")
}

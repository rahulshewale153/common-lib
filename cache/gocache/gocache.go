package gocache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	gc *goCacheWrapper
)

// NewGoCache initializes the GoCache instance.
func NewGoCache(defaultExpiration, cleanupInterval time.Duration) {
	gc = &goCacheWrapper{cache.New(defaultExpiration, cleanupInterval)}
}

type goCacheWrapper struct {
	cache *cache.Cache
}

func (gcw *goCacheWrapper) Set(key string, value interface{}, expiration time.Duration) {
	gcw.cache.Set(key, value, expiration)
}

func (gcw *goCacheWrapper) Get(key string) (interface{}, bool) {
	return gcw.cache.Get(key)
}

func (gcw *goCacheWrapper) Delete(key string) {
	gcw.cache.Delete(key)
}

func Set(key string, value interface{}, expiration time.Duration) {
	gc.Set(key, value, expiration)
}

func Get(key string) (interface{}, bool) {
	return gc.Get(key)
}

func Delete(key string) {
	gc.Delete(key)
}

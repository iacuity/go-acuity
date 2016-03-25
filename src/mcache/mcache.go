package mcache

import (
	"errors"
	"log"
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	CACHE_EXPIRY_ROUTINE_RUN_INTERVAL = 1 * time.Minute // default eviction algorithm run interval
)

var (
	gCache *cache.Cache // private cache object
)

// defaultExpiry: default cache item expiration time in second
func Init(defaultExpiry int) error {

	log.Println("defaultExpiry", time.Duration(defaultExpiry)*time.Second)

	gCache = cache.New(time.Duration(defaultExpiry)*time.Second,
		CACHE_EXPIRY_ROUTINE_RUN_INTERVAL)

	if nil == gCache {
		return errors.New("Error while initializing cache")
	}

	return nil
}

// set object in cache with provided expiry time
// Override the default expiry time with provided expiry time
func SetWithExpiry(key string, value interface{}, expiry time.Duration) {
	gCache.Set(key, value, expiry)
}

// set object in cache with default expiry
func Set(key string, value interface{}) {
	gCache.Set(key, value, cache.DefaultExpiration)
}

// fetch object from cache
func Get(key string) (interface{}, bool) {
	return gCache.Get(key)
}

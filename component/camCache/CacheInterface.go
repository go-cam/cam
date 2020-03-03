package camCache

import "time"

// cache driver interface
type CacheInterface interface {
	// init cache driver. connect server or create cache system
	Init() error
	// put cache
	Set(key string, value interface{}, duration time.Duration) error
	// get cache
	Get(key string) interface{}
	// delete cache
	Del(keys ...string) error
	// get GC interval
	GetGCInterval() time.Duration
	// garbage collection
	GC() error
	// clear all cache
	Flush() error
}

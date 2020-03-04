package camCache

import (
	"github.com/go-redis/redis/v7"
	"time"
)

// redis cache engine
type RedisCache struct {
	CacheInterface
	Addr     string
	Password string
	DB       int

	client *redis.Client
}

var _ CacheInterface = new(RedisCache)

// new redis engine
func NewRedisEngine() *RedisCache {
	engine := new(RedisCache)
	engine.Addr = "127.0.0.1:6379"
	engine.Password = "" // no password set
	engine.DB = 0        // use default DB
	engine.client = nil
	return engine
}

// init engine
func (cache *RedisCache) Init() error {
	cache.client = redis.NewClient(&redis.Options{
		Addr:     cache.Addr,
		Password: cache.Password,
		DB:       cache.DB,
	})
	return nil
}

// put key-value to engine
func (cache *RedisCache) Set(key string, value interface{}, duration time.Duration) error {
	return cache.client.Set(key, value, duration).Err()
}

// get value
func (cache *RedisCache) Get(key string) interface{} {
	value, err := cache.client.Get(key).Result()
	if err != nil {
		return nil
	}
	return value
}

// delete value form engine
func (cache *RedisCache) Del(keys ...string) error {
	return cache.client.Del(keys...).Err()
}

// get GC interval
func (cache *RedisCache) GetGCInterval() time.Duration {
	return -1
}

// garbage collection
func (cache *RedisCache) GC() error {
	return nil
}

// clear all cache
func (cache *RedisCache) Flush() error {
	return cache.client.FlushAll().Err()
}

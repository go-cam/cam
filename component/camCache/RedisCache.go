package camCache

import (
	"encoding/base64"
	"github.com/go-cam/cam/base/camUtils"
	"github.com/go-redis/redis/v7"
	"time"
)

// redis cache engine
type RedisCache struct {
	CacheInterface
	Addr           string
	Password       string
	DB             int
	encryptHandler func(value interface{}) interface{}
	decryptHandler func(value interface{}) interface{}
	client         *redis.Client
}

// new redis engine
func NewRedisEngine() *RedisCache {
	cache := new(RedisCache)
	cache.Addr = "127.0.0.1:6379"
	cache.Password = "" // no password set
	cache.DB = 0        // use default DB
	cache.client = nil
	cache.encryptHandler = cache.defaultCryptHandler
	cache.decryptHandler = cache.defaultCryptHandler
	return cache
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
	value = cache.encryptHandler(value)
	return cache.client.Set(key, value, duration).Err()
}

// get value
func (cache *RedisCache) Get(key string) interface{} {
	value, err := cache.client.Get(key).Result()
	if err != nil {
		return nil
	}
	return cache.decryptHandler(value)
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

// default encrypt and decrypt handler
func (cache *RedisCache) defaultCryptHandler(value interface{}) interface{} {
	return value
}

// set custom encrypt and decrypt handler
func (cache *RedisCache) SetCustomCrypt(encryptHandler func(interface{}) interface{}, dectypeHandler func(interface{}) interface{}) {
	cache.encryptHandler = encryptHandler
	cache.decryptHandler = dectypeHandler
}

// use base64's encrypt and decrypt handler
func (cache *RedisCache) SetBase64Crypt() {
	cache.encryptHandler = cache.base64EncryptHandler
	cache.decryptHandler = cache.base64DecryptHandler
}

// base64 encrypt handler
func (cache *RedisCache) base64EncryptHandler(value interface{}) interface{} {
	ao := NewCacheAo(value)
	bytes := camUtils.Json.Encode(ao)
	return base64.StdEncoding.EncodeToString(bytes)
}

// base64 decrypt handler
func (cache *RedisCache) base64DecryptHandler(value interface{}) interface{} {
	str, ok := value.(string)
	if !ok {
		panic("value was not type of string")
	}
	bytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic(err)
	}
	ao := new(CacheAo)
	camUtils.Json.DecodeToObj(bytes, ao)
	return ao.Value
}

package camCache

import (
	"github.com/go-cam/cam/base/camBase"
	"time"
)

// cache component config
type CacheComponentConfig struct {
	camBase.ComponentConfig

	Engine          CacheInterface
	DefaultDuration time.Duration
}

// new cache config
func NewCacheConfig() *CacheComponentConfig {
	config := new(CacheComponentConfig)
	config.Component = &CacheComponent{}
	config.Engine = NewFileCache()
	config.DefaultDuration = 7 * 24 * time.Hour
	return config
}

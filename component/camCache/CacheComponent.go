package camCache

import (
	"github.com/go-cam/cam/base/camBase"
	"time"
)

// cache component
type CacheComponent struct {
	camBase.Component
	camBase.CacheComponentInterface

	config *CacheComponentConfig
}

// init config
func (comp *CacheComponent) Init(configI camBase.ComponentConfigInterface) {
	comp.Component.Init(configI)

	var ok bool
	comp.config, ok = configI.(*CacheComponentConfig)
	if !ok {
		camBase.App.Fatal("CacheComponent", "invalid config")
		return
	}

	err := comp.config.Engine.Init()
	if err != nil {
		camBase.App.Error("CacheComponent.Init", err.Error())
	}
}

// start
func (comp *CacheComponent) Start() {
	comp.Component.Start()

	if comp.config.Engine.GetGCInterval() > 0 {
		go comp.gcLoop()
	}
}

// stop
func (comp *CacheComponent) Stop() {
	comp.Component.Stop()
}

// set cache
func (comp *CacheComponent) Set(key string, value interface{}) error {
	return comp.config.Engine.Set(key, value, comp.config.DefaultDuration)
}

// set cache with duration
func (comp *CacheComponent) SetDuration(key string, value interface{}, duration time.Duration) error {
	return comp.config.Engine.Set(key, value, duration)
}

// whether the key exists
func (comp *CacheComponent) Exists(key string) bool {
	value := comp.config.Engine.Get(key)
	return value != nil
}

// get value by key
func (comp *CacheComponent) Get(key string) interface{} {
	return comp.config.Engine.Get(key)
}

// delete cache
func (comp *CacheComponent) Del(keys ...string) error {
	return comp.config.Engine.Del(keys...)
}

// delete all cache
func (comp *CacheComponent) Flush() error {
	return comp.config.Engine.Flush()
}

// gc loop
func (comp *CacheComponent) gcLoop() {
	interval := comp.config.Engine.GetGCInterval()
	for {
		camBase.App.Debug("CacheComponent.gcLoop", "run")
		err := comp.config.Engine.GC()
		if err != nil {
			camBase.App.Error("CacheComponent.gcLoop", err.Error())
		}

		time.Sleep(interval)
	}
}

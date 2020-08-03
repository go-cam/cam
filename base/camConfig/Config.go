package camConfig

import (
	"github.com/go-cam/cam/base/camStatics"
)

// config struct
// config can merge in application
type Config struct {
	camStatics.AppConfigInterface
	// Application config
	AppConfig *AppConfig
	// Params.
	// You can get value by cam.App.Param(key string)
	Params map[string]interface{}
	// Components's config
	ComponentDict map[string]camStatics.ComponentConfigInterface
}

// new config
func NewConfig() *Config {
	config := new(Config)
	config.AppConfig = nil
	config.ComponentDict = map[string]camStatics.ComponentConfigInterface{}
	config.Params = map[string]interface{}{}
	return config
}

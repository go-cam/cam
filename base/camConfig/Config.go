package camConfig

import (
	"github.com/go-cam/cam/base/camBase"
)

// config struct
// config can merge in application
type Config struct {
	camBase.AppConfigInterface
	// Application config
	AppConfig *AppConfig
	// Params.
	// You can get value by cam.App.Param(key string)
	Params map[string]interface{}
	// Components's config
	ComponentDict map[string]camBase.ComponentConfigInterface
}

// new config
func NewConfig() *Config {
	config := new(Config)
	config.AppConfig = nil
	config.ComponentDict = map[string]camBase.ComponentConfigInterface{}
	config.Params = map[string]interface{}{}
	return config
}

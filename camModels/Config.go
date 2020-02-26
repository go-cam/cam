package camModels

import (
	"github.com/go-cam/cam/camBase"
)

// config struct
// config can merge in application
type Config struct {
	BaseModel
	// Application config
	AppConfig *AppConfig
	// Params required by the business logic
	Params map[string]interface{}
	// Components's config
	ComponentDict map[string]camBase.ConfigComponentInterface
}

// init params
func (config *Config) Init() {
	config.AppConfig = nil
	config.ComponentDict = map[string]camBase.ConfigComponentInterface{}
	config.Params = map[string]interface{}{}
}

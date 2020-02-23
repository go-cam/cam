package camConfigs

import (
	"github.com/go-cam/cam/core/camBase"
)

// base config
type BaseConfig struct {
	camBase.ConfigComponentInterface
	Component camBase.ComponentInterface // Instance of corresponding component
}

// get component instance
func (config *BaseConfig) GetComponent() camBase.ComponentInterface {
	return config.Component
}

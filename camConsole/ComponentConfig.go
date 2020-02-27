package camConsole

import (
	"github.com/go-cam/cam/camConfigs"
	"github.com/go-cam/cam/camPluginRouter"
)

// console config
type ComponentConfig struct {
	camConfigs.BaseConfig
	camPluginRouter.PluginConfig
}

// new console config
func NewComponentConfig() *ComponentConfig {
	config := new(ComponentConfig)
	config.Component = new(Component)
	config.PluginConfig.Init()
	return config
}

package camConsole

import (
	"github.com/go-cam/cam/camConfigs"
	"github.com/go-cam/cam/camPluginRouter"
)

// console config
type ConsoleComponentConfig struct {
	camConfigs.BaseConfig
	camPluginRouter.RouterPluginConfig
}

// new console config
func NewConsoleComponentConfig() *ConsoleComponentConfig {
	config := new(ConsoleComponentConfig)
	config.Component = &ConsoleComponent{}
	config.RouterPluginConfig.Init()
	return config
}

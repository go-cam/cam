package camConsole

import (
	"github.com/go-cam/cam/camBase"
	"github.com/go-cam/cam/camPluginRouter"
)

// console config
type ConsoleComponentConfig struct {
	camBase.ComponentConfig
	camPluginRouter.RouterPluginConfig
}

// new console config
func NewConsoleComponentConfig() *ConsoleComponentConfig {
	config := new(ConsoleComponentConfig)
	config.Component = &ConsoleComponent{}
	config.RouterPluginConfig.Init()

	config.registerFrameworkController()
	return config
}

// register controller in the framework
func (config *ConsoleComponentConfig) registerFrameworkController() {
	config.RouterPluginConfig.Register(&MigrateController{})
	config.RouterPluginConfig.Register(&XormController{})
}

package camConsole

import (
	"github.com/go-cam/cam/camBase"
	"github.com/go-cam/cam/camPluginRouter"
)

// console config
type ConsoleComponentConfig struct {
	camBase.Config
	camPluginRouter.RouterPluginConfig
}

// new console config
func NewConsoleComponentConfig() *ConsoleComponentConfig {
	config := new(ConsoleComponentConfig)
	config.Component = &ConsoleComponent{}
	config.RouterPluginConfig.Init()
	config.RouterPluginConfig.Register(&MigrateController{})
	config.RouterPluginConfig.Register(&XormController{})
	return config
}

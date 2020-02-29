package cam

import (
	"github.com/go-cam/cam/camModels"
)

// new RouterPluginConfig
func NewConfig() *camModels.Config {
	config := new(camModels.Config)
	config.Init()
	return config
}

package cam

import (
	"github.com/go-cam/cam/camModels"
)

// module config status
type Config struct {
	camModels.Config
}

// new Config
func NewConfig() *Config {
	config := new(Config)
	configModel := new(camModels.Config)
	configModel.Init()
	config.Config = *configModel
	return config
}

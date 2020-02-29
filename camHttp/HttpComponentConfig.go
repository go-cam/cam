package camHttp

import (
	"github.com/go-cam/cam/camBase"
	"github.com/go-cam/cam/camPluginContext"
	"github.com/go-cam/cam/camPluginRouter"
	"github.com/go-cam/cam/camPluginSsl"
)

// http server config
type HttpComponentConfig struct {
	camBase.ComponentConfig
	camPluginRouter.RouterPluginConfig
	camPluginContext.ContextPluginConfig
	camPluginSsl.SslPluginConfig

	Port        uint16
	SessionName string
	SessionKey  string
}

// new config
func NewHttpComponentConfig(port uint16) *HttpComponentConfig {
	config := new(HttpComponentConfig)
	config.Component = &HttpComponent{}
	config.Port = port
	config.SessionName = "cam"
	config.SessionKey = "cam"
	config.RouterPluginConfig.Init()
	config.ContextPluginConfig.Init()
	config.ContextPluginConfig.Init()
	return config
}

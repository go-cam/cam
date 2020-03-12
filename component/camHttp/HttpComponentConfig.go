package camHttp

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/component"
	"github.com/go-cam/cam/plugin/camPluginContext"
	"github.com/go-cam/cam/plugin/camPluginRouter"
	"github.com/go-cam/cam/plugin/camPluginSsl"
)

// http server config
type HttpComponentConfig struct {
	component.ComponentConfig
	camPluginRouter.RouterPluginConfig
	camPluginContext.ContextPluginConfig
	camPluginSsl.SslPluginConfig

	Port        uint16
	SessionName string
	SessionKey  string

	routeHandlerDict map[string]camBase.HttpRouteHandler
}

// new config
func NewHttpComponentConfig(port uint16) *HttpComponentConfig {
	config := new(HttpComponentConfig)
	config.Component = &HttpComponent{}
	config.Port = port
	config.SessionName = "cam"
	config.SessionKey = "cam"
	config.routeHandlerDict = map[string]camBase.HttpRouteHandler{}
	config.RouterPluginConfig.Init()
	config.ContextPluginConfig.Init()
	config.ContextPluginConfig.Init()
	return config
}

// add custom route handler.
// its priority is higher than the controller.
func (config *HttpComponentConfig) AddRoute(route string, handler camBase.HttpRouteHandler) {
	config.routeHandlerDict[route] = handler
}

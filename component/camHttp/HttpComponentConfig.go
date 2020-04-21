package camHttp

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/component"
	"github.com/go-cam/cam/plugin/camContext"
	"github.com/go-cam/cam/plugin/camMiddleware"
	"github.com/go-cam/cam/plugin/camRouter"
	"github.com/go-cam/cam/plugin/camSsl"
)

// http server config
type HttpComponentConfig struct {
	component.ComponentConfig
	camRouter.RouterPluginConfig
	camContext.ContextPluginConfig
	camSsl.SslPluginConfig
	camMiddleware.MiddlewarePluginConfig

	Port        uint16
	SessionName string
	SessionKey  string

	// Deprecated: remove on v0.5.0
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
	config.SslPluginConfig.Init()
	config.MiddlewarePluginConfig.Init()
	config.SetContextStruct(&HttpContext{})
	return config
}

// add custom route handler.
// its priority is higher than the controller.
// Deprecated: remove on v0.5.0  It's not support middleware
// Instead: HttpComponentConfig.MiddlewarePluginConfig.AddRoute()
func (config *HttpComponentConfig) AddRoute(route string, handler camBase.HttpRouteHandler) {
	config.routeHandlerDict[route] = handler
}

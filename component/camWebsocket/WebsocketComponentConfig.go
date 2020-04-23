package camWebsocket

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/component"
	"github.com/go-cam/cam/plugin/camContext"
	"github.com/go-cam/cam/plugin/camMiddleware"
	"github.com/go-cam/cam/plugin/camRouter"
	"github.com/go-cam/cam/plugin/camSsl"
)

// websocket server 所需的配置
type WebsocketComponentConfig struct {
	component.ComponentConfig
	camSsl.SslPluginConfig
	camRouter.RouterPluginConfig
	camContext.ContextPluginConfig
	camMiddleware.MiddlewarePluginConfig

	Port uint16 // server port
	// message parse handler
	//
	// message: Messages sent by clients
	//
	// controllerName:
	// actionName:
	// values: send data, just like post form data
	MessageParseHandler camBase.MessageParseHandler

	// Deprecated: remove on v0.5.0  it's not support Middleware
	routeHandlerDict map[string]camBase.WebsocketRouteHandler
}

// new websocket component
func NewWebsocketComponentConfig(port uint16) *WebsocketComponentConfig {
	config := new(WebsocketComponentConfig)
	config.Component = &WebsocketComponent{}
	config.Port = port
	config.routeHandlerDict = map[string]camBase.WebsocketRouteHandler{}
	config.SslPluginConfig.Init()
	config.RouterPluginConfig.Init()
	config.ContextPluginConfig.Init()
	config.MiddlewarePluginConfig.Init()
	config.ContextPluginConfig.SetContextStruct(&WebsocketContext{})
	return config
}

// add custom route handler.
// its priority is higher than the controller.
// Deprecated: remove on v0.5.0  it's not support Middleware
// Instead: WebsocketComponentConfig.RouterPluginConfig.AddRoute()
func (config *WebsocketComponentConfig) AddRoute(route string, handler camBase.WebsocketRouteHandler) {
	config.routeHandlerDict[route] = handler
}

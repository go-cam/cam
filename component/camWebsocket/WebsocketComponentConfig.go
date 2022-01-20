package camWebsocket

import (
	"github.com/go-cam/cam/component"
	"github.com/go-cam/cam/plugin/camContext"
	"github.com/go-cam/cam/plugin/camMiddleware"
	"github.com/go-cam/cam/plugin/camRouter"
	"github.com/go-cam/cam/plugin/camTls"
)

// websocket server 所需的配置
type WebsocketComponentConfig struct {
	component.ComponentConfig
	camTls.TlsPluginConfig
	camRouter.RouterPluginConfig
	camContext.ContextPluginConfig
	camMiddleware.MiddlewarePluginConfig

	Port uint16 // server port
}

// new websocket component
func NewWebsocketComponentConfig(port uint16) *WebsocketComponentConfig {
	config := new(WebsocketComponentConfig)
	config.Component = &WebsocketComponent{}
	config.Port = port
	config.TlsPluginConfig.Init()
	config.RouterPluginConfig.Init()
	config.ContextPluginConfig.Init()
	config.MiddlewarePluginConfig.Init()
	config.ContextPluginConfig.SetContextStruct(&WebsocketContext{})
	return config
}

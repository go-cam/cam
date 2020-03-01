package camWebsocket

import (
	"github.com/go-cam/cam/camBase"
	"github.com/go-cam/cam/camPluginContext"
	"github.com/go-cam/cam/camPluginRouter"
	"github.com/go-cam/cam/camPluginSsl"
)

// websocket server 所需的配置
type WebsocketComponentConfig struct {
	camBase.ComponentConfig
	camPluginSsl.SslPluginConfig
	camPluginRouter.RouterPluginConfig
	camPluginContext.ContextPluginConfig

	Port uint16 // server port

	// message parse handler
	//
	// message: Messages sent by clients
	//
	// controllerName:
	// actionName:
	// values: send data, just like post form data
	MessageParseHandler camBase.WebsocketComponentMessageParseHandler
}

// new websocket component
func NewWebsocketComponentConfig(port uint16) *WebsocketComponentConfig {
	config := new(WebsocketComponentConfig)
	config.Component = &WebsocketComponent{}
	config.Port = port
	config.SslPluginConfig.Init()
	config.RouterPluginConfig.Init()
	config.ContextPluginConfig.Init()
	return config
}
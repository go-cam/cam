package camConfigs

import "github.com/go-cam/cam/camBase"

// websocket server 所需的配置
type WebsocketServer struct {
	camBase.ComponentConfig
	RouterPlugin
	SslPlugin
	ContextPlugin
	Port uint16 // server port

	// message parse handler
	//
	// message: Messages sent by clients
	//
	// controllerName:
	// actionName:
	// values: send data, just like post form data
	MessageParseHandler func(message []byte) (controllerName string, actionName string, values map[string]interface{})
}

// listen SslPlugin
func (config *WebsocketServer) ListenSsl(port uint16, certFile string, keyFile string) *WebsocketServer {
	config.SslPlugin.ListenSsl(port, certFile, keyFile)
	return config
}

// only SSl mode.
func (config *WebsocketServer) SslOnly() *WebsocketServer {
	config.SslPlugin.SslOnly()
	return config
}

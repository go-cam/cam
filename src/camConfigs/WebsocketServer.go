package camConfigs

// websocket server 所需的配置
type WebsocketServer struct {
	BaseConfig
	PluginRouter
	PluginSsl
	PluginContext
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

// listen PluginSsl
func (config *WebsocketServer) ListenSsl(port uint16, certFile string, keyFile string) *WebsocketServer {
	config.PluginSsl.ListenSsl(port, certFile, keyFile)
	return config
}

// only SSl mode.
func (config *WebsocketServer) SslOnly() *WebsocketServer {
	config.PluginSsl.SslOnly()
	return config
}

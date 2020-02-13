package camConfigs

// websocket server 所需的配置
type WebsocketServer struct {
	BaseConfig
	PluginRouter
	PluginSsl
	Port uint16 // 服务器端口

	// 传输消息解析器
	// message: 客户端发送过来的消息
	// controllerName: 控制器名字
	// actionName: 控制器方法名字
	// values: 传输的参数
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

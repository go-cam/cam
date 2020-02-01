package camConfigs

// websocket server 所需的配置
type WebsocketServer struct {
	BaseConfig
	PluginRouter
	Port uint16 // 服务器端口

	// 传输消息解析器
	// message: 客户端发送过来的消息
	// controllerName: 控制器名字
	// actionName: 控制器方法名字
	// values: 传输的参数
	MessageParseHandler func(message []byte) (controllerName string, actionName string, values map[string]interface{})
}

package configs

import "cin/src/base"

// websocket server 所需的配置
type WebsocketServer struct {
	BaseConfig
	Port uint16                   // 服务器端口
}

// 新建 websocket server 配置
func NewWebsocketServer(component base.ComponentInterface, port uint16) *WebsocketServer {
	config := new(WebsocketServer)
	config.Port = port
	config.Component = component
	return config
}
package configs

import "cin/src/alias"

// websocket server 所需的配置
type WebsocketServer struct {
	ComponentInterface
	Port uint16                    // 服务器端口
	Mode alias.WebsocketServerMode // 运行模式
}

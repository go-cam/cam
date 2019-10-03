package configs

import (
	"controllers"
	"models"
)

// 路由插件。添加这个插件后的配置可获取路由的参数
type PluginRouter struct {
	HandlerList               []controllers.HandlerInterface
	OnWebsocketMessageHandler func(conn *models.WebsocketSession, recvMessage []byte)
}
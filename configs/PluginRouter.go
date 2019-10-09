package configs

import (
	"cin/base"
	"cin/models"
)

// 路由插件。添加这个插件后的配置可获取路由的参数
type PluginRouter struct {
	ControllerList            []base.ControllerInterface
	OnWebsocketMessageHandler func(conn *models.Context, recvMessage []byte)
}
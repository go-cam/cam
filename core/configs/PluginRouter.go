package configs

import (
	"github.com/cinling/cam/core/base"
	"github.com/cinling/cam/core/models"
)

// 路由插件。添加这个插件后的配置可获取路由的参数
type PluginRouter struct {
	// http 或 websocket 控制器列表
	ControllerList []base.ControllerInterface
	// 控制台 控制器列表
	ConsoleControllerList []base.ControllerInterface
	// websocket 接收到参数执行的方法
	OnWebsocketMessageHandler func(conn *models.Context, recvMessage []byte)
}

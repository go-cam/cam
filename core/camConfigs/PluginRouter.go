package camConfigs

import (
	"github.com/go-cam/cam/core/camBase"
	"github.com/go-cam/cam/core/camModels"
)

// 路由插件。添加这个插件后的配置可获取路由的参数
type PluginRouter struct {
	// http 或 websocket 控制器列表
	ControllerList []camBase.ControllerInterface
	// 控制台 控制器列表
	ConsoleControllerList []camBase.ControllerInterface
	// websocket 接收到参数执行的方法
	OnWebsocketMessageHandler func(conn *camModels.Context, recvMessage []byte)
}

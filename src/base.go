package cin

import (
	"cin/src/constants"
	"cin/src/controllers"
	"cin/src/models"
)

const (
	ComponentWebsocketServer = constants.ComponentWebsocketServer

	// 应用状态：初始化
	ApplicationStatusInit = constants.ApplicationStatusInit
	// 应用状态：开始（启动中）
	ApplicationStatusStart = constants.ApplicationStatusStart
	// 应用状态：运行中
	ApplicationStatusRun = constants.ApplicationStatusRun
	// 应用状态：即将停止
	ApplicationStatusStop = constants.ApplicationStatusStop
	// 应用状态：销毁
	ApplicationStatusDestroy = constants.ApplicationStatusDestroy

	// websocket server 运行模式：自动处理【推荐】
	// 使用框架内规定的 Handler 或 Controller 自动匹配对应的方法。发送数据必须是规范的数据。
	// 使用该模式依然可以使用 OnMessage 接收数据。但是不能发送数据
	WebsocketServerModeAutoHandler = constants.WebsocketServerModeAutoHandler
	// websocket server 运行模式：自定义处理。
	// 自定义 OnMessage 回调方法发送数据。根据实际需求自定义数据返回
	WebsocketServerModeCustom = constants.WebsocketServerModeCustom

	// 控制器类型：websocket
	ControllerTypeWebsocket = constants.ControllerTypeWebsocket
	// 控制器类型：http
	ControllerTypeHttp = constants.ControllerTypeHttp
)

// 基础 websocket 处理器
type Controller struct {
	controllers.BaseController
}

// 配置类
type Config struct {
	*models.Config
}
// 新建配置
func NewConfig() *Config {
	config := new(Config)
	config.Config = new(models.Config)
	config.Config.Init()
	return config
}

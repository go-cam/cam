package cam

import (
	"github.com/go-cam/cam/core/camConstants"
	"github.com/go-cam/cam/core/camModels"
)

const (
	ComponentWebsocketServer = camConstants.ComponentWebsocketServer

	// 应用状态：初始化
	ApplicationStatusInit = camConstants.ApplicationStatusInit
	// 应用状态：开始（启动中）
	ApplicationStatusStart = camConstants.ApplicationStatusStart
	// 应用状态：即将停止
	ApplicationStatusStop = camConstants.ApplicationStatusStop

	// websocket server 运行模式：自动处理【推荐】
	// 使用框架内规定的 Handler 或 BaseController 自动匹配对应的方法。发送数据必须是规范的数据。
	// 使用该模式依然可以使用 OnMessage 接收数据。但是不能发送数据
	WebsocketServerModeAutoHandler = camConstants.WebsocketServerModeAutoHandler
	// websocket server 运行模式：自定义处理。
	// 自定义 OnMessage 回调方法发送数据。根据实际需求自定义数据返回
	WebsocketServerModeCustom = camConstants.WebsocketServerModeCustom

	// 控制器类型：websocket
	ControllerTypeWebsocket = camConstants.ControllerTypeWebsocket
	// 控制器类型：http
	ControllerTypeHttp = camConstants.ControllerTypeHttp
)

// 配置类
type Config struct {
	camModels.Config
}

// 新建配置
func NewConfig() *Config {
	config := new(Config)
	configModel := new(camModels.Config)
	configModel.Init()
	config.Config = *configModel
	return config
}

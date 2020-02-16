package cam

import (
	consoleController "github.com/go-cam/cam/console/camConsoleControllers"
	"github.com/go-cam/cam/core/camBase"
)

// 路由器
type router struct {
	controllerList            []camBase.ControllerInterface
	consoleControllerList     []camBase.ControllerInterface
	onWebsocketMessageHandler func(conn camBase.ContextInterface, recvMessage []byte)
}

// 新建路由器
func newRouter() *router {
	r := new(router)
	r.controllerList = []camBase.ControllerInterface{}
	r.onWebsocketMessageHandler = nil
	r.consoleControllerList = []camBase.ControllerInterface{}
	r.registerDefaultConsoleController()
	return r
}

// 写入默认的命令行控制器
func (r *router) registerDefaultConsoleController() {
	r.RegisterConsole(new(consoleController.MigrateController))
	r.RegisterConsole(new(consoleController.XormController))
}

// 注册控制器
func (r *router) Register(controller camBase.ControllerInterface) {
	r.controllerList = append(r.controllerList, controller)
}

// 注册命令行控制器
func (r *router) RegisterConsole(controller camBase.ControllerInterface) {
	r.consoleControllerList = append(r.consoleControllerList, controller)
}

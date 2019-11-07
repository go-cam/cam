package cin

import (
	"github.com/cinling/cin/base"
	consoleController "github.com/cinling/cin/console/controllers"
	"github.com/cinling/cin/models"
)

// 路由器
type router struct {
	controllerList            []base.ControllerInterface
	consoleControllerList     []base.ControllerInterface
	onWebsocketMessageHandler func(conn *models.Context, recvMessage []byte)
}

// 新建路由器
func newRouter() *router {
	r := new(router)
	r.controllerList = []base.ControllerInterface{}
	r.onWebsocketMessageHandler = nil
	r.consoleControllerList = []base.ControllerInterface{}
	r.registerDefaultConsoleController()
	return r
}

// 写入默认的命令行控制器
func (r *router) registerDefaultConsoleController() {
	r.RegisterConsole(new(consoleController.MigrateController))
}

// 注册控制器
func (r *router) Register(controller base.ControllerInterface) {
	r.controllerList = append(r.controllerList, controller)
}

// 注册命令行控制器
func (r *router) RegisterConsole(controller base.ControllerInterface) {
	r.consoleControllerList = append(r.consoleControllerList, controller)
}

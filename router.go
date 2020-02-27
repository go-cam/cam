package cam

import (
	"github.com/go-cam/cam/camBase"
)

// router
// Deprecated: remove on v0.3.0
type router struct {
	controllerList            []camBase.ControllerBakInterface
	consoleControllerList     []camBase.ControllerBakInterface
	onWebsocketMessageHandler func(conn camBase.ContextInterface, recvMessage []byte)
}

// new router
// Deprecated: remove on v0.3.0
func newRouter() *router {
	r := new(router)
	r.controllerList = []camBase.ControllerBakInterface{}
	r.onWebsocketMessageHandler = nil
	r.consoleControllerList = []camBase.ControllerBakInterface{}
	r.registerDefaultConsoleController()
	return r
}

// Deprecated: remove on v0.3.0
func (r *router) registerDefaultConsoleController() {
	//r.RegisterConsole(new(camConsole.MigrateController))
	//r.RegisterConsole(new(camConsole.XormController))
}

// Deprecated: remove on v0.3.0
func (r *router) Register(controller camBase.ControllerBakInterface) {
	r.controllerList = append(r.controllerList, controller)
}

// Deprecated: remove on v0.3.0
func (r *router) RegisterConsole(controller camBase.ControllerBakInterface) {
	r.consoleControllerList = append(r.consoleControllerList, controller)
}

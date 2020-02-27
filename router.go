package cam

import (
	"github.com/go-cam/cam/camBase"
	"github.com/go-cam/cam/camConsole"
)

// router
type router struct {
	controllerList            []camBase.ControllerBakInterface
	consoleControllerList     []camBase.ControllerBakInterface
	onWebsocketMessageHandler func(conn camBase.ContextInterface, recvMessage []byte)
}

// new router
func newRouter() *router {
	r := new(router)
	r.controllerList = []camBase.ControllerBakInterface{}
	r.onWebsocketMessageHandler = nil
	r.consoleControllerList = []camBase.ControllerBakInterface{}
	r.registerDefaultConsoleController()
	return r
}

//
func (r *router) registerDefaultConsoleController() {
	r.RegisterConsole(new(camConsole.MigrateController))
	r.RegisterConsole(new(camConsole.XormController))
}

//
func (r *router) Register(controller camBase.ControllerBakInterface) {
	r.controllerList = append(r.controllerList, controller)
}

//
func (r *router) RegisterConsole(controller camBase.ControllerBakInterface) {
	r.consoleControllerList = append(r.consoleControllerList, controller)
}

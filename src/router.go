package cam

import (
	"github.com/go-cam/cam/camBase"
	consoleController "github.com/go-cam/cam/camConsoleControllers"
)

// router
type router struct {
	controllerList            []camBase.ControllerInterface
	consoleControllerList     []camBase.ControllerInterface
	onWebsocketMessageHandler func(conn camBase.ContextInterface, recvMessage []byte)
}

// new router
func newRouter() *router {
	r := new(router)
	r.controllerList = []camBase.ControllerInterface{}
	r.onWebsocketMessageHandler = nil
	r.consoleControllerList = []camBase.ControllerInterface{}
	r.registerDefaultConsoleController()
	return r
}

//
func (r *router) registerDefaultConsoleController() {
	r.RegisterConsole(new(consoleController.MigrateController))
	r.RegisterConsole(new(consoleController.XormController))
}

//
func (r *router) Register(controller camBase.ControllerInterface) {
	r.controllerList = append(r.controllerList, controller)
}

//
func (r *router) RegisterConsole(controller camBase.ControllerInterface) {
	r.consoleControllerList = append(r.consoleControllerList, controller)
}

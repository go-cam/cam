package camRouter

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/plugin"
)

// router plugin.
// it can inject component if it need to
type RouterPluginConfig struct {
	camBase.PluginConfigInterface

	// recover route
	recoverRoute string
	// controller list
	controllerList []camBase.ControllerInterface
	// custom handler
	customHandlerDict map[string]camBase.RouteHandler
	// receive message parse handler
	recvMessageParseHandler plugin.RecvMessageParseHandler
}

// init
func (config *RouterPluginConfig) Init() {
	config.controllerList = []camBase.ControllerInterface{}
	config.customHandlerDict = map[string]camBase.RouteHandler{}
	config.recvMessageParseHandler = nil
}

// register controller
// controller: Inout &Controller{} OR new(Controller) OR Controller{} is ok
func (config *RouterPluginConfig) Register(controller camBase.ControllerInterface) {
	config.controllerList = append(config.controllerList, controller)
}

// recover route.
// when panic on controller, Is will catch error and redirect to this route
func (config *RouterPluginConfig) RecoverRoute(route string) {
	config.recoverRoute = route
}

// add router handler.
// its priority is higher than the controller.
func (config *RouterPluginConfig) AddRoute(route string, handler camBase.RouteHandler) {
	config.customHandlerDict[route] = handler
}

func (config *RouterPluginConfig) SetRecvMessageParseHandler(handler plugin.RecvMessageParseHandler) {
	config.recvMessageParseHandler = handler
}

func (config *RouterPluginConfig) GetRecvMessageParseHandler() plugin.RecvMessageParseHandler {
	if config.recvMessageParseHandler == nil {
		return plugin.DefaultRecvToMessageHandler
	}
	return config.recvMessageParseHandler
}

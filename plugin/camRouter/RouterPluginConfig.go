package camRouter

import (
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/plugin"
)

// public options
type RouterOption struct {
	// Default route
	DefaultRoute string
	// Recover route
	RecoverRoute string
	// Receive message parse handler
	RecvMessageParseHandler plugin.RecvMessageParseHandler
}

// router plugin.
// it can inject component if it need to
type RouterPluginConfig struct {
	camStatics.PluginConfigInterface

	option *RouterOption

	// controller list
	controllerList []camStatics.ControllerInterface
	// custom handler
	customHandlerDict map[string]camStatics.RouteHandler
}

// init
func (conf *RouterPluginConfig) Init() {
	conf.option = &RouterOption{}
	conf.controllerList = []camStatics.ControllerInterface{}
	conf.customHandlerDict = map[string]camStatics.RouteHandler{}
}

// register controller
// controller: Inout &Controller{} OR new(Controller) OR Controller{} is ok
func (conf *RouterPluginConfig) Register(controller camStatics.ControllerInterface) {
	conf.controllerList = append(conf.controllerList, controller)
}

// Deprecated: remove since v0.6.0. instead by RouterOption()
// recover route.
// when panic on controller, Is will catch error and redirect to this route
func (conf *RouterPluginConfig) RecoverRoute(route string) {
	conf.option.RecoverRoute = route
}

// add router handler.
// its priority is higher than the controller.
func (conf *RouterPluginConfig) AddRoute(route string, handler camStatics.RouteHandler) {
	conf.customHandlerDict[route] = handler
}

// Deprecated: remove since v0.6.0. instead by RouterOption()
func (conf *RouterPluginConfig) SetRecvMessageParseHandler(handler plugin.RecvMessageParseHandler) {
	conf.option.RecvMessageParseHandler = handler
}

func (conf *RouterPluginConfig) GetRecvMessageParseHandler() plugin.RecvMessageParseHandler {
	if conf.option.RecvMessageParseHandler == nil {
		return plugin.DefaultRecvToMessageHandler
	}
	return conf.option.RecvMessageParseHandler
}

// Set router option
func (conf *RouterPluginConfig) RouterOption(opt *RouterOption) {
	conf.option = opt
}

// Get default route
func (conf *RouterPluginConfig) DefaultRoute() string {
	return conf.option.DefaultRoute
}

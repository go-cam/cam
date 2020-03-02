package cam

import (
	"github.com/go-cam/cam/component/camConsole"
	"github.com/go-cam/cam/plugin/camPluginContext"
	"github.com/go-cam/cam/plugin/camPluginRouter"
)

type Controller struct {
	camPluginRouter.Controller
}

type ConstantController struct {
	camConsole.ConsoleController
}

type ControllerAction struct {
	camPluginRouter.ControllerAction
}

type Context struct {
	camPluginContext.Context
}

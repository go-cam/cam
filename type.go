package cam

import (
	"github.com/go-cam/cam/camConsole"
	"github.com/go-cam/cam/camPluginContext"
	"github.com/go-cam/cam/camPluginRouter"
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

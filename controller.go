package cam

import "github.com/go-cam/cam/core/camModels"

// base controller [websocket/http/console]
type BaseController struct {
	camModels.BaseController
}

// get application instance
func (controller *BaseController) GetApp() *application {
	return controller.GetAppInterface().(*application)
}

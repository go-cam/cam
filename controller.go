package cam

import "github.com/go-cam/cam/camModels"

// base controller [websocket/http/console]
// Deprecated: instead by camModels.Controller. Remove on v0.3.0
type BaseController struct {
	camModels.BaseController
}

// get Application instance
// Deprecated: instead by camModels.Controller. Remove on v0.3.0
func (controller *BaseController) GetApp() *Application {
	return controller.GetAppInterface().(*Application)
}

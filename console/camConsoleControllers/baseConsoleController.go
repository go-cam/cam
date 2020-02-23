package camConsoleControllers

import (
	"github.com/go-cam/cam/core/camComponents"
	"github.com/go-cam/cam/core/camModels"
	"os"
)

//
type baseConsoleController struct {
	camModels.BaseController
}

// get database component instance
func (controller *baseConsoleController) GetDatabaseComponent() *camComponents.Database {
	ins := controller.GetAppInterface().GetComponentByName("db")
	if ins == nil {
		return nil
	}
	return ins.(*camComponents.Database)
}

// get params
// key start is 0
// example:
// 		/path/to/main.exe console/run argv0 argv1
//		controller.GetArgv(0) => "argv0"
//		controller.GetArgv(1) => "argv1"
//		controller.GetArgv(2) => ""
//		controller.GetArgv(-1) => ""
func (controller *baseConsoleController) GetArgv(key int) string {
	key += 2
	if key < 2 {
		return ""
	}
	if len(os.Args) <= key {
		return ""
	}
	return os.Args[key]
}

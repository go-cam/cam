package camConsole

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/plugin/camPluginRouter"
	"os"
)

//
type ConsoleController struct {
	camPluginRouter.Controller
}

// get params
// key start is 0
// example:
// 		/path/to/main.exe console/run argv0 argv1
//		controller.GetArgv(0) => "argv0"
//		controller.GetArgv(1) => "argv1"
//		controller.GetArgv(2) => ""
//		controller.GetArgv(-1) => ""
func (controller *ConsoleController) GetArgv(key int) string {
	key += 2
	if key < 2 {
		return ""
	}
	if len(os.Args) <= key {
		return ""
	}
	return os.Args[key]
}

// get console component
func (controller *ConsoleController) GetConsoleComponent() *ConsoleComponent {
	componentI := camBase.App.GetComponent(&ConsoleComponent{})
	console, ok := componentI.(*ConsoleComponent)
	if !ok {
		panic("console component is not enable")
	}
	return console
}

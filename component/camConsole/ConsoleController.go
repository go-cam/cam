package camConsole

import (
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/plugin/camRouter"
	"os"
)

//
type ConsoleController struct {
	camRouter.Controller
}

// get params
// key start is 0
// example:
// 		/path/to/main.exe console/run argv0 argv1
//		controller.GetArgv(0) => "argv0"
//		controller.GetArgv(1) => "argv1"
//		controller.GetArgv(2) => ""
//		controller.GetArgv(-1) => ""
func (ctrl *ConsoleController) GetArgv(key int) string {
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
func (ctrl *ConsoleController) GetConsoleComponent() *ConsoleComponent {
	componentI := camStatics.App.GetComponent(&ConsoleComponent{})
	console, ok := componentI.(*ConsoleComponent)
	if !ok {
		panic("console component is not enable")
	}
	return console
}

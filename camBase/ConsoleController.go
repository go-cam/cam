package camBase

import (
	"os"
)

//
type ConsoleController struct {
	Controller
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

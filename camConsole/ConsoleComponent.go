package camConsole

import (
	"fmt"
	"github.com/go-cam/cam/camBase"
	"github.com/go-cam/cam/camPluginRouter"
	"os"
	"reflect"
)

// command component
type ConsoleComponent struct {
	camBase.Component
	camPluginRouter.RouterPlugin

	config *ConsoleComponentConfig
}

// init
func (component *ConsoleComponent) Init(configInterface camBase.ComponentConfigInterface) {
	component.Component.Init(configInterface)

	configValue := reflect.ValueOf(configInterface)
	var config *ConsoleComponentConfig
	if configValue.Kind() == reflect.Ptr {
		config = configValue.Interface().(*ConsoleComponentConfig)
	} else if configValue.Kind() == reflect.Struct {
		configStruct := configValue.Interface().(ConsoleComponentConfig)
		config = &configStruct
	} else {
		panic("illegal config")
	}
	component.config = config

	// register controller
	//component.controllerDict, component.controllerActionDict = camComponents.Common.GetControllerDict(config.ControllerList)

	// init router plugin
	component.RouterPlugin.Init(&config.RouterPluginConfig)
}

// run command
// Example:
//	# go build cam.go
//	# ./cam controllerName/actionName param1 param2
func (component *ConsoleComponent) RunAction() {
	if len(os.Args) < 2 {
		fmt.Println("please input route")
		return
	}

	route := os.Args[1]
	controller, action, _, _ := component.GetControllerAction(route)
	if controller == nil || action == nil {
		panic("route not found: " + route)
	}
	controller.Init()
	controller.SetApp(component.App)
	if !controller.BeforeAction(action) {
		panic("invalid call")
		return
	}
	action.Call()
	response := controller.AfterAction(action, controller.GetResponse())
	fmt.Println(string(response))
}

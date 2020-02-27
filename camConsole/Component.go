package camConsole

import (
	"fmt"
	"github.com/go-cam/cam/camBase"
	"github.com/go-cam/cam/camComponents"
	"github.com/go-cam/cam/camPluginRouter"
	"os"
	"reflect"
)

// command component
type Component struct {
	camComponents.Base
	camPluginRouter.Plugin

	config *ComponentConfig
}

// init
func (component *Component) Init(configInterface camBase.ConfigComponentInterface) {
	component.Base.Init(configInterface)

	configValue := reflect.ValueOf(configInterface)
	var config *ComponentConfig
	if configValue.Kind() == reflect.Ptr {
		config = configValue.Interface().(*ComponentConfig)
	} else if configValue.Kind() == reflect.Struct {
		configStruct := configValue.Interface().(ComponentConfig)
		config = &configStruct
	} else {
		panic("illegal config")
	}
	component.config = config

	// register controller
	//component.controllerDict, component.controllerActionDict = camComponents.Common.GetControllerDict(config.ControllerList)

	// init router plugin
	component.Plugin.Init(&config.PluginConfig)
}

// run command
// Example:
//	# go build cam.go
//	# ./cam controllerName/actionName param1 param2
func (component *Component) RunAction() {
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

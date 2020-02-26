package camComponents

import (
	"fmt"
	"github.com/go-cam/cam/camBase"
	"github.com/go-cam/cam/camConfigs"
	"github.com/go-cam/cam/camUtils"
	"os"
	"reflect"
	"strings"
)

// command component
type Console struct {
	Base

	config *camConfigs.Console

	controllerDict       map[string]reflect.Type    // controller reflect.Type dict
	controllerActionDict map[string]map[string]bool // map[controllerName]map[actionName]
}

// init
func (component *Console) Init(configInterface camBase.ConfigComponentInterface) {
	component.Base.Init(configInterface)

	configValue := reflect.ValueOf(configInterface)
	var config *camConfigs.Console
	if configValue.Kind() == reflect.Ptr {
		config = configValue.Interface().(*camConfigs.Console)
	} else if configValue.Kind() == reflect.Struct {
		configStruct := configValue.Interface().(camConfigs.Console)
		config = &configStruct
	} else {
		panic("illegal config")
	}
	component.config = config

	// register controller
	component.controllerDict, component.controllerActionDict = common.getControllerDict(config.ConsoleControllerList)
}

// run command
// Example:
//	# go build cam.go
//	# ./cam controllerName/actionName param1 param2
func (component *Console) RunAction() {
	if len(os.Args) < 2 {
		fmt.Println("please input route")
		return
	}
	route := os.Args[1]
	tmpArr := strings.Split(route, "/")
	controllerName := camUtils.Url.UrlToHump(tmpArr[0])
	actionName := ""
	if len(tmpArr) >= 2 {
		actionName = camUtils.Url.UrlToHump(tmpArr[1])
	}

	controllerType := component.controllerDict[controllerName]
	controllerValue := reflect.New(controllerType.Elem())
	controllerInterface := controllerValue.Interface().(camBase.ControllerInterface)

	// init controller
	controllerInterface.Init()
	controllerInterface.SetApp(component.app)

	controllerInterface.BeforeAction(actionName)

	action := controllerValue.MethodByName(actionName)
	_ = action.Call([]reflect.Value{})

	controllerInterface.AfterAction(actionName, nil)
}

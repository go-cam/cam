package camPluginRouter

import (
	"github.com/go-cam/cam/camBase"
	"github.com/go-cam/cam/camUtils"
	"reflect"
	"strings"
)

type RouterPlugin struct {
	camBase.PluginInterface

	config *RouterPluginConfig // plugin config

	controllerDict       map[string]reflect.Type    // controller reflect.Type dict
	controllerActionDict map[string]map[string]bool // map[controllerName]map[actionName]
}

func (plugin *RouterPlugin) Init(config *RouterPluginConfig) {
	plugin.config = config
	plugin.controllerDict = map[string]reflect.Type{}
	plugin.controllerActionDict = map[string]map[string]bool{}

	plugin.parseController()
}

// parse controller List
func (plugin *RouterPlugin) parseController() {
	excludeMethodNameDict := plugin.getExcludeActionDict()

	for _, controllerInterface := range plugin.config.ControllerList {
		controllerType := reflect.TypeOf(controllerInterface)
		if controllerType.Kind() == reflect.Ptr {
			controllerType = controllerType.Elem()
		}
		controllerName := controllerType.Name()
		controllerName = strings.TrimSuffix(controllerName, "Controller")
		plugin.controllerDict[controllerName] = controllerType

		controllerValue := reflect.New(controllerType)
		controllerInterface := controllerValue.Interface().(camBase.ControllerInterface)
		if controllerInterface == nil {
			panic(controllerName + " must be implement base.ControllerBakInterface")
		}

		// save all action of controller
		plugin.controllerActionDict[controllerName] = map[string]bool{}
		methodLen := controllerType.NumMethod()
		for i := 0; i < methodLen; i++ {
			method := controllerType.Method(i)
			methodName := method.Name

			if _, exclude := excludeMethodNameDict[methodName]; exclude {
				// exclude action
				continue
			}

			plugin.controllerActionDict[controllerName][methodName] = true
		}
	}
}

// Get controller and action by route
//
// route: 				controller and action name.  Example: "user/register-and-login"
//
// controller: 			the controller interface
//						return nil if controller not exists
// action: 				the method, controller's action. You can call by `action.Call()` is ok
//						return nil if action not exists
// controllerName: 		Example: "User"
// actionName: 			Example: "RegisterAndLogin"
func (plugin *RouterPlugin) GetControllerAction(route string) (controller camBase.ControllerInterface, action camBase.ControllerActionInterface, controllerName string, actionName string) {
	tmpArr := strings.Split(route, "/")

	controllerName = camUtils.Url.UrlToHump(tmpArr[0])
	controllerType, has := plugin.controllerDict[controllerName]
	if !has {
		return nil, nil, controllerName, ""
	}

	controllerValue := reflect.New(controllerType)
	controller = controllerValue.Interface().(camBase.ControllerInterface)

	actionName = ""

	tmpArrLen := len(tmpArr)
	if tmpArrLen >= 2 {
		actionName = camUtils.Url.UrlToHump(tmpArr[1])
	} else if tmpArrLen == 1 {
		actionName = controller.GetDefaultActionName()
	}

	actionValue := controllerValue.MethodByName(actionName)
	if !actionValue.IsValid() { // method not exists
		return controller, nil, controllerName, actionName
	}
	action = camBase.NewControllerAction(route, &actionValue)

	return controller, action, controllerName, actionName
}

// exclude the camModels.BaseController method, this is not a user callable action
func (plugin *RouterPlugin) getExcludeActionDict() map[string]bool {
	excludeDict := map[string]bool{}

	t := reflect.TypeOf(new(camBase.Controller))
	methodLen := t.NumMethod()
	for i := 0; i < methodLen; i++ {
		method := t.Method(i)
		methodName := method.Name
		excludeDict[methodName] = true
	}

	return excludeDict
}

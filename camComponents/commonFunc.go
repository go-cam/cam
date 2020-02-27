package camComponents

import (
	"github.com/go-cam/cam/camBase"
	"github.com/go-cam/cam/camModels"
	"reflect"
	"strings"
)

// component Common function
// Deprecated: instead by camPlugins.RouterPlugin
var Common = newCommonFunc()

// component share function
// Deprecated: instead by camPlugins.RouterPlugin
type commonFunc struct {
	excludeDict map[string]bool // exclude controller names
}

// new instance
func newCommonFunc() *commonFunc {
	cf := new(commonFunc)
	cf.excludeDict = nil
	return cf
}

// get controller and action dict
// Deprecated: instead by camPlugins.RouterPlugin
func (cf *commonFunc) GetControllerDict(controllerList []camBase.ControllerBakInterface) (map[string]reflect.Type, map[string]map[string]bool) {
	controllerDict := map[string]reflect.Type{}
	controllerActionDict := map[string]map[string]bool{}

	excludeMethodNameDict := cf.GetControllerExcludeMethodNameDict()

	for _, controllerInterface := range controllerList {
		t := reflect.TypeOf(controllerInterface)
		controllerType := t.Elem()
		controllerName := controllerType.Name()
		controllerName = strings.TrimSuffix(controllerName, "Controller")
		controllerDict[controllerName] = t

		controllerValue := reflect.New(controllerType)
		controllerInterface := controllerValue.Interface().(camBase.ControllerBakInterface)
		if controllerInterface == nil {
			panic(controllerName + " must be implement base.ControllerBakInterface")
		}

		// save all action of controller
		controllerActionDict[controllerName] = map[string]bool{}
		methodLen := t.NumMethod()
		for i := 0; i < methodLen; i++ {
			method := t.Method(i)
			methodName := method.Name

			if _, exclude := excludeMethodNameDict[methodName]; exclude {
				// exclude action
				continue
			}

			controllerActionDict[controllerName][methodName] = true
		}
	}

	return controllerDict, controllerActionDict
}

// get controller exclude action dict
// Deprecated: instead by camPlugins.RouterPlugin
func (cf *commonFunc) GetControllerExcludeMethodNameDict() map[string]bool {
	if cf.excludeDict == nil {
		cf.excludeDict = map[string]bool{}

		t := reflect.TypeOf(new(camModels.BaseController))
		methodLen := t.NumMethod()
		for i := 0; i < methodLen; i++ {
			method := t.Method(i)
			methodName := method.Name
			cf.excludeDict[methodName] = true
		}
	}

	return cf.excludeDict
}

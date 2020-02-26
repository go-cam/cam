package camComponents

import (
	"github.com/go-cam/cam/camBase"
	"github.com/go-cam/cam/camModels"
	"reflect"
	"strings"
)

// component common function
var common = newCommonFunc()

// component share function
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
func (cf *commonFunc) getControllerDict(controllerList []camBase.ControllerInterface) (map[string]reflect.Type, map[string]map[string]bool) {
	controllerDict := map[string]reflect.Type{}
	controllerActionDict := map[string]map[string]bool{}

	excludeMethodNameDict := cf.getControllerExcludeMethodNameDict()

	for _, controllerInterface := range controllerList {
		t := reflect.TypeOf(controllerInterface)
		controllerType := t.Elem()
		controllerName := controllerType.Name()
		controllerName = strings.TrimSuffix(controllerName, "Controller")
		controllerDict[controllerName] = t

		controllerValue := reflect.New(controllerType)
		controllerInterface := controllerValue.Interface().(camBase.ControllerInterface)
		if controllerInterface == nil {
			panic(controllerName + " must be implement base.ControllerInterface")
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
func (cf *commonFunc) getControllerExcludeMethodNameDict() map[string]bool {
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

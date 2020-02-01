package camComponents

import (
	"github.com/go-cam/cam/core/camBase"
	"github.com/go-cam/cam/core/camModels"
	"reflect"
	"strings"
)

// 组件内通用的方法
var common = newCommonFunc()

// 设置app
func SetApplication(app camBase.ApplicationInterface) {
	common.app = app
}

// 组件内通用的方法封装
type commonFunc struct {
	excludeDict map[string]bool
	app         camBase.ApplicationInterface
}

//
func newCommonFunc() *commonFunc {
	cf := new(commonFunc)
	cf.excludeDict = nil
	return cf
}

// 获取控制器记录map
func (cf *commonFunc) getControllerDict(controllerList []camBase.ControllerInterface) (map[string]reflect.Type, map[string]map[string]bool) {
	controllerDict := map[string]reflect.Type{}
	controllerActionDict := map[string]map[string]bool{}

	excludeMethodNameDict := cf.getControllerExcludeMethodNameDict()

	for _, controllerInterface := range controllerList {
		t := reflect.TypeOf(controllerInterface)
		controllerType := t.Elem() // 获取实体
		controllerName := controllerType.Name()
		controllerName = strings.TrimSuffix(controllerName, "Controller")
		controllerDict[controllerName] = t

		controllerValue := reflect.New(controllerType)
		controllerInterface := controllerValue.Interface().(camBase.ControllerInterface)
		if controllerInterface == nil {
			panic(controllerName + " must be implement base.ControllerInterface")
		}

		// 保存控制器啊所有方法名字
		controllerActionDict[controllerName] = map[string]bool{}
		methodLen := t.NumMethod()
		for i := 0; i < methodLen; i++ {
			method := t.Method(i)
			methodName := method.Name

			// 判断是否是排除的方法名字
			if _, exclude := excludeMethodNameDict[methodName]; exclude {
				continue
			}

			controllerActionDict[controllerName][methodName] = true
		}
	}

	return controllerDict, controllerActionDict
}

// 获取控制器排除的方法名字
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

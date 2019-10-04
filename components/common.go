package components

import (
	"cin/controllers"
	"reflect"
	"strings"
)

// 组件内通用的方法封装
type commonFunc struct {
}

// 组件内通用的方法
var common = new(commonFunc)

// 获取控制器记录map
func (cf *commonFunc) getControllerDict(controllerList []controllers.HandlerInterface) (map[string]reflect.Type, map[string]map[string]bool) {
	handlerDict := map[string]reflect.Type{}
	handlerActionDict := map[string]map[string]bool{}

	excludeMethodNameDict := cf.getControllerExcludeMethodNameDict()

	for _, handlerInterface := range controllerList {
		t := reflect.TypeOf(handlerInterface)
		handlerType := t.Elem() // 获取实体
		handlerName := handlerType.Name()
		handlerName = strings.TrimSuffix(handlerName, "Handler")
		handlerName = strings.TrimSuffix(handlerName, "Controller")
		handlerName = strings.ToLower(handlerName)
		handlerDict[handlerName] = t

		// 保存控制器啊所有方法名字
		handlerActionDict[handlerName] = map[string]bool{}
		methodLen := t.NumMethod()
		for i := 0; i < methodLen; i++ {
			method := t.Method(i)
			methodName := method.Name

			// 判断是否是排除的方法名字
			if _, exclude := excludeMethodNameDict[methodName]; exclude {
				continue
			}

			handlerActionDict[handlerName][methodName] = true
		}
	}

	return handlerDict, handlerActionDict
}

// 获取控制器排除的方法名字
func (cf *commonFunc) getControllerExcludeMethodNameDict() map[string]bool {
	excludeDict := map[string]bool{}

	t := reflect.TypeOf(new(controllers.BaseController))
	methodLen := t.NumMethod()
	for i := 0; i < methodLen; i++ {
		method := t.Method(i)
		methodName := method.Name
		excludeDict[methodName] = true
	}

	return excludeDict
}

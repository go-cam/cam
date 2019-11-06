package components

import (
	"cin/base"
	"cin/configs"
	"cin/utils"
	"fmt"
	"os"
	"reflect"
	"strings"
)

// 命令行组件
type Console struct {
	Base

	config *configs.Console

	controllerDict       map[string]reflect.Type    // 控制器反射map
	controllerActionDict map[string]map[string]bool // 控制器 => 方法 => 是否存在（注册时记录）
}

// 初始化方法
func (component *Console) Init(configInterface base.ConfigComponentInterface) {
	component.Base.Init(configInterface)

	configValue := reflect.ValueOf(configInterface)
	var config *configs.Console
	if configValue.Kind() == reflect.Ptr {
		config = configValue.Interface().(*configs.Console)
	} else if configValue.Kind() == reflect.Struct {
		configStruct := configValue.Interface().(configs.Console)
		config = &configStruct
	} else {
		panic("illegal config")
	}
	component.config = config

	// 注册处理器（控制器）
	component.controllerDict, component.controllerActionDict = common.getControllerDict(config.ConsoleControllerList)
}

// 运行控制台命令
func (component *Console) RunAction() {
	if len(os.Args) < 2 {
		fmt.Println("please input route")
		return
	}
	route := os.Args[1]
	tmpArr := strings.Split(route, "/")
	controllerName := tmpArr[0]
	actionName := ""
	if len(tmpArr) >= 2 {
		actionName = tmpArr[1]
	}

	controllerType := component.controllerDict[controllerName]
	controllerValue := reflect.New(controllerType.Elem())
	controllerInterface := controllerValue.Interface().(base.ControllerInterface)
	controllerInterface.Init()
	component.injectControllerValues(controllerInterface)

	controllerInterface.BeforeAction(actionName)

	action := controllerValue.MethodByName(utils.Url.UrlToHump(actionName))
	_ = action.Call([]reflect.Value{})

	controllerInterface.AfterAction(actionName, nil)
}

// 注入控制器参数
func (component *Console) injectControllerValues(controllerIns base.ControllerInterface) {
	controllerName := utils.Reflect.GetClassName(controllerIns)
	if controllerName == "MigrateController" {
		databaseComponentIns := common.app.GetComponent(new(Database))
		if databaseComponentIns == nil {
			return
		}
		databaseComponent := databaseComponentIns.(*Database)

		controllerIns.AddValue("migrateDir", databaseComponent.config.MigrateDir)
	}
}
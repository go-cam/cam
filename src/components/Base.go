package components

import (
	"cin/src/base"
	"fmt"
	"reflect"
)

// 所有组件的基类
type Base struct {
	name string // 组件名字
	isPrintLog bool
	base.ComponentInterface
}

// 使用配置初始化数据
func (component *Base) Init(configInterface base.ConfigComponentInterface) {
	component.name = component.getComponentName(configInterface.GetComponent())
	component.isPrintLog = true
}

func (component *Base) Start(configDict map[string]interface{}) {

}

func (component *Base) Run(configDict map[string]interface{}) {

}

func (component *Base) Stop(configDict map[string]interface{}) {

}

func (component *Base) Destroy(configDict map[string]interface{}) {

}

// 打印日志
func (component *Base) Log(message string) {
	if !component.isPrintLog {
		return
	}

	fmt.Println(message)
}

// 获取组件名字
func (component *Base) getComponentName(componentInterface base.ComponentInterface) string {
	t := reflect.TypeOf(componentInterface)
	return t.Elem().Name()
}
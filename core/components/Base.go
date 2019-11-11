package components

import (
	base2 "github.com/cinling/cin/core/base"
	"reflect"
)

// 所有组件的基类
type Base struct {
	base2.ComponentInterface

	// component name
	name string
	// app instance
	app base2.ApplicationInterface
}

// set app instance
func (component *Base) SetApp(app base2.ApplicationInterface) {
	component.app = app
}

// 使用配置初始化数据
func (component *Base) Init(configInterface base2.ConfigComponentInterface) {
	component.name = component.getComponentName(configInterface.GetComponent())
}

// 开始
func (component *Base) Start() {

}

// 结束
func (component *Base) Stop() {

}

// 获取组件名字
func (component *Base) getComponentName(componentInterface base2.ComponentInterface) string {
	t := reflect.TypeOf(componentInterface)
	return t.Elem().Name()
}

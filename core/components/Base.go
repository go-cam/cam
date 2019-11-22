package components

import (
	"github.com/go-cam/cam/core/base"
	"reflect"
)

// 所有组件的基类
type Base struct {
	base.ComponentInterface

	// component name
	name string
	// app instance
	app base.ApplicationInterface
}

// set app instance
func (component *Base) SetApp(app base.ApplicationInterface) {
	component.app = app
}

// 使用配置初始化数据
func (component *Base) Init(configInterface base.ConfigComponentInterface) {
	component.name = component.getComponentName(configInterface.GetComponent())
}

// 开始
func (component *Base) Start() {

}

// 结束
func (component *Base) Stop() {

}

// 获取组件名字
func (component *Base) getComponentName(componentInterface base.ComponentInterface) string {
	t := reflect.TypeOf(componentInterface)
	return t.Elem().Name()
}

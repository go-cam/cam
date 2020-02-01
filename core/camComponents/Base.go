package camComponents

import (
	"github.com/go-cam/cam/core/camBase"
	"reflect"
)

// 所有组件的基类
type Base struct {
	camBase.ComponentInterface

	// component name
	name string
	// app instance
	app camBase.ApplicationInterface
}

// set app instance
func (component *Base) SetApp(app camBase.ApplicationInterface) {
	component.app = app
}

// 使用配置初始化数据
func (component *Base) Init(configInterface camBase.ConfigComponentInterface) {
	component.name = component.getComponentName(configInterface.GetComponent())
}

// 开始
func (component *Base) Start() {

}

// 结束
func (component *Base) Stop() {

}

// 获取组件名字
func (component *Base) getComponentName(componentInterface camBase.ComponentInterface) string {
	t := reflect.TypeOf(componentInterface)
	return t.Elem().Name()
}

package components

import (
	"github.com/cinling/cin/base"
	"reflect"
)

// 所有组件的基类
type Base struct {
	name string // 组件名字
	base.ComponentInterface
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

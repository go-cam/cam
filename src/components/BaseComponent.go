package components

import "cin/src/base"

// 所有组件的基类
type BaseComponent struct {
	base.ComponentInterface
}

// 使用配置初始化数据
func (component *BaseComponent) Init(configInterface base.ConfigComponentInterface) {

}

func (component *BaseComponent) Start() {

}

func (component *BaseComponent) Run() {

}

func (component *BaseComponent) Stop() {

}

func (component *BaseComponent) Destroy() {

}
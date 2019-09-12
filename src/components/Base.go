package components

import "cin/src/base"

// 所有组件的基类
type Base struct {
	base.ComponentInterface
}

// 使用配置初始化数据
func (component *Base) Init(configInterface base.ConfigComponentInterface) {

}

func (component *Base) Start(configDict map[string]interface{}) {

}

func (component *Base) Run(configDict map[string]interface{}) {

}

func (component *Base) Stop(configDict map[string]interface{}) {

}

func (component *Base) Destroy(configDict map[string]interface{}) {

}
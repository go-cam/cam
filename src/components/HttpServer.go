package components

import "base"

// http服务
type HttpServer struct {
	Base

	port uint16 // http 端口
}

// 使用配置初始化数据
func (component *HttpServer) Init(configInterface base.ConfigComponentInterface) {
	component.Base.Init(configInterface)

	component.name = component.getComponentName(configInterface.GetComponent())
}

// 启动
func (component *HttpServer) Start() {
	component.Base.Start()


}
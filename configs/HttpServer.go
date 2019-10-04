package configs

import "cin/base"

// http 服务
type HttpServer struct {
	BaseConfig
	PluginRouter
	Port uint16
}

//新建http服务
func NewHttpServer(component base.ComponentInterface, port uint16) *HttpServer {
	config := new(HttpServer)
	config.Component = component
	config.Port = port
	return config
}

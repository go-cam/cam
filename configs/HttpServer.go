package configs

import "cin/base"

// http 服务
type HttpServer struct {
	BaseConfig
	PluginRouter
	Port uint16
	SessionName string
	SessionKey string
}

//新建http服务
func NewHttpServer(component base.ComponentInterface, port uint16) *HttpServer {
	config := new(HttpServer)
	config.Component = component
	config.Port = port
	config.SessionKey = "cin-key"
	config.SessionName = "cin"
	return config
}

// 设置 session key
func (config *HttpServer) SetSessionKey(sessionKey string) *HttpServer {
	config.SessionKey = sessionKey
	return config
}

// 设置session name
func (config *HttpServer) SetSessionName(sessionName string) *HttpServer {
	config.SessionName = sessionName
	return config
}
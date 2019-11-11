package configs

// http 服务
type HttpServer struct {
	BaseConfig
	PluginRouter
	Port        uint16
	SessionName string
	SessionKey  string
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

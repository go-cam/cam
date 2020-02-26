package camConfigs

// http server config
type HttpServer struct {
	BaseConfig
	RouterPlugin
	SslPlugin
	ContextPlugin
	Port        uint16
	SessionName string
	SessionKey  string
}

// set session key
func (config *HttpServer) SetSessionKey(sessionKey string) *HttpServer {
	config.SessionKey = sessionKey
	return config
}

// set session name
func (config *HttpServer) SetSessionName(sessionName string) *HttpServer {
	config.SessionName = sessionName
	return config
}

// listen SslPlugin
func (config *HttpServer) ListenSsl(port uint16, certFile string, keyFile string) *HttpServer {
	config.SslPlugin.ListenSsl(port, certFile, keyFile)
	return config
}

// only SSl mode.
func (config *HttpServer) SslOnly() *HttpServer {
	config.SslPlugin.SslOnly()
	return config
}

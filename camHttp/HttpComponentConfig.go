package camHttp

import (
	"github.com/go-cam/cam/camConfigs"
	"github.com/go-cam/cam/camPluginContext"
	"github.com/go-cam/cam/camPluginRouter"
)

// http server config
type HttpComponentConfig struct {
	camConfigs.BaseConfig
	camConfigs.RouterPlugin
	camConfigs.SslPlugin
	camConfigs.ContextPlugin
	camPluginRouter.RouterPluginConfig
	camPluginContext.ContextPluginConfig

	Port        uint16
	SessionName string
	SessionKey  string
}

// new config
func NewHttpComponentConfig(port uint16) *HttpComponentConfig {
	config := new(HttpComponentConfig)
	config.Component = &HttpComponent{}
	config.Port = port
	config.SessionName = "cam"
	config.SessionKey = "cam"

	config.RouterPluginConfig.Init()
	config.ContextPluginConfig.Init()
	return config
}

// set session key
// Deprecated: remove on v0.3.0
func (config *HttpComponentConfig) SetSessionKey(sessionKey string) *HttpComponentConfig {
	config.SessionKey = sessionKey
	return config
}

// set session name
// Deprecated: remove on v0.3.0
func (config *HttpComponentConfig) SetSessionName(sessionName string) *HttpComponentConfig {
	config.SessionName = sessionName
	return config
}

// listen SslPlugin
// Deprecated: remove on v0.3.0
func (config *HttpComponentConfig) ListenSsl(port uint16, certFile string, keyFile string) *HttpComponentConfig {
	config.SslPlugin.ListenSsl(port, certFile, keyFile)
	return config
}

// only SSl mode.
// Deprecated: remove on v0.3.0
func (config *HttpComponentConfig) SslOnly() *HttpComponentConfig {
	config.SslPlugin.SslOnly()
	return config
}

package camHttp

import (
	"github.com/go-cam/cam/component"
	"github.com/go-cam/cam/plugin/camContext"
	"github.com/go-cam/cam/plugin/camMiddleware"
	"github.com/go-cam/cam/plugin/camRouter"
	"github.com/go-cam/cam/plugin/camTls"
)

// http server config
type HttpComponentConfig struct {
	component.ComponentConfig
	camRouter.RouterPluginConfig
	camContext.ContextPluginConfig
	camTls.TlsPluginConfig
	camMiddleware.MiddlewarePluginConfig

	Port uint16

	sessionStore  Store
	sessionOption *SessionOption
}

// new config
func NewHttpComponentConfig(port uint16) *HttpComponentConfig {
	config := new(HttpComponentConfig)
	config.Component = &HttpComponent{}
	config.Port = port
	config.RouterPluginConfig.Init()
	config.ContextPluginConfig.Init()
	config.TlsPluginConfig.Init()
	config.MiddlewarePluginConfig.Init()
	config.sessionStore = nil
	config.sessionOption = nil
	config.SetContextStruct(&HttpContext{})
	return config
}

func (conf *HttpComponentConfig) SetSessionStore(store Store) {
	conf.sessionStore = store
}

func (conf *HttpComponentConfig) SetSessionOption(opt *SessionOption) {
	conf.sessionOption = opt
}

func (conf *HttpComponentConfig) getSessionStore() Store {
	if conf.sessionStore == nil {
		return conf.defaultSessionStore()
	}
	return conf.sessionStore
}

func (conf *HttpComponentConfig) defaultSessionStore() Store {
	return NewCacheStore("SESSION")
}

func (conf *HttpComponentConfig) getSessionOption() *SessionOption {
	if conf.sessionOption == nil {
		return new(SessionOption)
	}
	return conf.sessionOption
}

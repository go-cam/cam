package camHttp

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/component"
	"github.com/go-cam/cam/plugin/camContext"
	"github.com/go-cam/cam/plugin/camMiddleware"
	"github.com/go-cam/cam/plugin/camRouter"
	"github.com/go-cam/cam/plugin/camSsl"
)

// http server config
type HttpComponentConfig struct {
	component.ComponentConfig
	camRouter.RouterPluginConfig
	camContext.ContextPluginConfig
	camSsl.SslPluginConfig
	camMiddleware.MiddlewarePluginConfig

	Port uint16
	// Deprecated: remove on v0.6.0
	SessionName string
	// Deprecated: remove on v0.6.0
	SessionKey string

	sessionStore  Store
	sessionOption *SessionOption

	// Deprecated: remove on v0.5.0
	routeHandlerDict map[string]camBase.HttpRouteHandler
}

// new config
func NewHttpComponentConfig(port uint16) *HttpComponentConfig {
	config := new(HttpComponentConfig)
	config.Component = &HttpComponent{}
	config.Port = port
	config.SessionName = "cam"
	config.SessionKey = "cam"
	config.routeHandlerDict = map[string]camBase.HttpRouteHandler{}
	config.RouterPluginConfig.Init()
	config.ContextPluginConfig.Init()
	config.SslPluginConfig.Init()
	config.MiddlewarePluginConfig.Init()
	config.sessionStore = nil
	config.sessionOption = nil
	config.SetContextStruct(&HttpContext{})
	return config
}

// add custom route handler.
// its priority is higher than the controller.
// Deprecated: remove on v0.5.0  It's not support middleware
// Instead: HttpComponentConfig.MiddlewarePluginConfig.AddRoute()
func (conf *HttpComponentConfig) AddRoute(route string, handler camBase.HttpRouteHandler) {
	conf.routeHandlerDict[route] = handler
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

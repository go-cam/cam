package camSocket

import (
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/component"
	"github.com/go-cam/cam/plugin/camContext"
	"github.com/go-cam/cam/plugin/camMiddleware"
	"github.com/go-cam/cam/plugin/camRouter"
	"time"
)

type SocketComponentConfig struct {
	component.ComponentConfig
	camRouter.RouterPluginConfig
	camContext.ContextPluginConfig
	camMiddleware.MiddlewarePluginConfig

	// tcp listen port
	Port uint16

	// =========== Default conn handler params ===========

	// Block transfer end.
	// Default: '\x17'
	//
	// Example:
	//	recvMessage := "17\x17receive a message"
	//	recvMessage can be divided into three parts: [len], [etb flat], [content] := "17", "\x17", "receive a message"
	//		[len]:          the content length
	//		[etb flat]:     block transfer end
	//		[content]:      actual received data
	Etb byte
	// Max number of bytes in a single receive message
	// Default: 1MB
	RecvMaxLen uint64
	// Max number of bytes in a single send message
	// Default: 128MB
	SendMaxLen uint64
	// Read recv message timeout
	RecvTimeout time.Duration
	// Write send message timeout
	SendTimeout time.Duration

	// socket conn handler
	ConnHandler camStatics.SocketConnHandler
	// trace recv and send message
	Trace bool
}

func NewSocketComponentConfig(port uint16) *SocketComponentConfig {
	config := new(SocketComponentConfig)
	config.Component = &SocketComponent{}
	config.Port = port
	config.Etb = '\x17'
	config.RecvMaxLen = camStatics.MB
	config.SendMaxLen = 128 * camStatics.MB
	config.RecvTimeout = 15 * time.Second
	config.SendTimeout = 60 * time.Second
	config.ConnHandler = nil
	config.Trace = false
	config.RouterPluginConfig.Init()
	config.ContextPluginConfig.Init()
	config.MiddlewarePluginConfig.Init()
	config.SetContextStruct(&SocketContext{})
	return config
}

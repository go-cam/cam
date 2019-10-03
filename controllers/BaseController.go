package controllers

import (
	"github.com/cinling/cin/base"
)

// http 处理器接口
type HttpHandlerInterface interface {
}

// websocket 和 http 通用控制器接口
type ControllerInterface interface {
	SetControllerType(controllerType base.ControllerType)
}

// 控制器（所有处理器的整合类）
type BaseController struct {
	ControllerInterface
	BaseWebsocketHandler

	controllerType base.ControllerType
}

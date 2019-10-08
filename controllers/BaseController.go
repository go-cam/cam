package controllers

import (
	"cin/base"
)

// 控制器（所有处理器的整合类）
type BaseController struct {
	BaseWebsocketHandler

	controllerType base.ControllerType
}

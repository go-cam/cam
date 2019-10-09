package cin

import (
	"cin/base"
	"cin/models"
)

// 路由器
type router struct {
	controllerList            []base.ControllerInterface
	onWebsocketMessageHandler func(conn *models.WebsocketSession, recvMessage []byte)
}

// 新建路由器
func newRouter() *router {
	r := new(router)
	r.controllerList = []base.ControllerInterface{}
	r.onWebsocketMessageHandler = nil
	return r
}

// 注册控制器
func (r *router) Register(controller base.ControllerInterface) {
	r.controllerList = append(r.controllerList, controller)
}
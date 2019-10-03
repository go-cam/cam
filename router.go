package cin

import (
	"cin/controllers"
	"cin/models"
)

// 路由器
type router struct {
	handlerList               []controllers.HandlerInterface
	onWebsocketMessageHandler func(conn *models.WebsocketSession, recvMessage []byte)
}

// 新建路由器
func newRouter() *router {
	r := new(router)
	r.handlerList = []controllers.HandlerInterface{}
	r.onWebsocketMessageHandler = nil
	return r
}

// 注册控制器
func (r *router) Register(controller controllers.HandlerInterface) {
	r.handlerList = append(r.handlerList, controller)
}

//
func (r *router) OnWebsocketMessage(handler func(conn *models.WebsocketSession, recvMessage []byte)) {
	r.onWebsocketMessageHandler = handler
}
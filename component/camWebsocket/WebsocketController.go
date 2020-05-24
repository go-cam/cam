package camWebsocket

import (
	"github.com/go-cam/cam/plugin/camRouter"
	"github.com/gorilla/websocket"
)

type WebsocketController struct {
	camRouter.Controller
}

// get *websocket.Conn
func (ctrl *WebsocketController) GetConn() *websocket.Conn {
	return ctrl.GetWebsocketContext().GetConn()
}

// get WebsocketContextInterface
func (ctrl *WebsocketController) GetWebsocketContext() WebsocketContextInterface {
	ctx, ok := ctrl.GetContext().(WebsocketContextInterface)
	if !ok {
		panic("context was not implement WebsocketContextInterface")
	}
	return ctx
}

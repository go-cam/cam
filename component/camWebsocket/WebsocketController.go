package camWebsocket

import (
	"github.com/go-cam/cam/plugin/camRouter"
	"github.com/gorilla/websocket"
)

type WebsocketController struct {
	camRouter.Controller
}

// get *websocket.Conn
func (controller *WebsocketController) GetConn() *websocket.Conn {
	session, ok := controller.GetSession().(*WebsocketSession)
	if !ok {
		return nil
	}
	return session.GetConn()
}

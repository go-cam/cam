package models

import (
	"github.com/gorilla/websocket"
	"net"
)

// websocket 使用的 session 。没有发送的功能。必须依赖 WebsocketServer 进行发送
// Deprecated: 需要实现 SessionInterface 接口
type WebsocketSession struct {
	BaseModel
	conn *websocket.Conn
	sendMessage []byte
}

// 新建websocket session
func NewWebsocketSession(conn *websocket.Conn) *WebsocketSession {
	model := new(WebsocketSession)
	model.conn = conn
	model.sendMessage = nil
	return model
}

// 关闭连接
func (model *WebsocketSession) Close() error {
	var err error
	// 获取关闭的 handler
	var closeHandler = model.conn.CloseHandler()
	if closeHandler != nil {
		err = closeHandler(0, "")
		if err != nil {
			return nil
		}
	}
	// 关闭连接
	err = model.conn.Close()
	if err != nil {
		return err
	}
	return nil
}

// 发送消息
func (model *WebsocketSession) Send(message []byte) {
	model.sendMessage = message
}

// 获取发送消息
func (model *WebsocketSession) GetSendMessage() []byte {
	return model.sendMessage
}

// 获取远端地址
func (model *WebsocketSession) RemoveAddr() net.Addr {
	return model.conn.RemoteAddr()
}

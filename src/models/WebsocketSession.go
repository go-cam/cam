package models

import "github.com/gorilla/websocket"

// websocket 使用的 session
type WebsocketSession struct {
	conn *websocket.Conn
}

// 新建websocket session
func NewWebsocketSession(conn *websocket.Conn) *WebsocketSession {
	model := new(WebsocketSession)
	model.conn = conn
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
func (model *WebsocketSession) Send(bytes []byte) error {
	return model.conn.WriteMessage(websocket.TextMessage, bytes)
}
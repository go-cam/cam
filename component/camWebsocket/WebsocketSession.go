package camWebsocket

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camUtils"
	"github.com/gorilla/websocket"
)

// websocket session
type WebsocketSession struct {
	camBase.SessionInterface

	conn      *websocket.Conn             // websocket connection
	sessionId string                      // sessionId
	values    map[interface{}]interface{} // save session value
}

func NewWebsocketSession(conn *websocket.Conn) *WebsocketSession {
	model := new(WebsocketSession)
	model.conn = conn
	model.sessionId = camUtils.String.UUID()
	model.values = map[interface{}]interface{}{}
	return model
}

// 获取 sessionId
func (session *WebsocketSession) GetSessionId() string {
	return session.sessionId
}

// 设置值
func (session *WebsocketSession) Set(key interface{}, value interface{}) {
	session.values[key] = value
}

// 获取值
func (session *WebsocketSession) Get(key interface{}) interface{} {
	value, has := session.values[key]
	if !has {
		return nil
	}
	return value
}

// 销毁session 清空 session 所有数据
func (session *WebsocketSession) Destroy() {
	_ = session.conn.Close()
	session.sessionId = ""
	session.values = map[interface{}]interface{}{}
}

// get client connection
func (session *WebsocketSession) GetConn() *websocket.Conn {
	return session.conn
}

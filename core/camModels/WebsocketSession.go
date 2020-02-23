package camModels

import (
	"github.com/go-cam/cam/core/camUtils"
	"github.com/gorilla/websocket"
)

// Deprecated:
// websocket 使用的 session 。没有发送的功能。必须依赖 WebsocketServer 进行发送
type WebsocketSession struct {
	BaseModel

	conn      *websocket.Conn             // websocket 链接
	sessionId string                      // sessionId 用于记录记录链接的sessionId
	values    map[interface{}]interface{} // session 存储的 key value 数据
}

func NewWebsocketSession(conn *websocket.Conn) *WebsocketSession {
	model := new(WebsocketSession)
	model.conn = conn
	model.sessionId = camUtils.String.UUID()
	model.values = map[interface{}]interface{}{}
	return model
}

// 获取 sessionId
func (model *WebsocketSession) GetSessionId() string {
	return model.sessionId
}

// 设置值
func (model *WebsocketSession) Set(key interface{}, value interface{}) {
	model.values[key] = value
}

// 获取值
func (model *WebsocketSession) Get(key interface{}) interface{} {
	value, has := model.values[key]
	if !has {
		return nil
	}
	return value
}

// 销毁session 清空 session 所有数据
func (model *WebsocketSession) Destroy() {
	_ = model.conn.Close()
	model.sessionId = ""
	model.values = map[interface{}]interface{}{}
}

// get client connection
func (model *WebsocketSession) GetConn() *websocket.Conn {
	return model.conn
}

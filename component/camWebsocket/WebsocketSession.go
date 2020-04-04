package camWebsocket

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camUtils"
	"github.com/gorilla/websocket"
	"sync"
)

// websocket session
type WebsocketSession struct {
	camBase.SessionInterface

	conn      *websocket.Conn // websocket connection
	sessionId string          // sessionId
	values    sync.Map        // save session value
}

// new websocket session
func NewWebsocketSession(conn *websocket.Conn) *WebsocketSession {
	sess := new(WebsocketSession)
	sess.conn = conn
	sess.sessionId = camUtils.String.UUID()
	sess.values = sync.Map{}
	return sess
}

// get sessionId
func (sess *WebsocketSession) GetSessionId() string {
	return sess.sessionId
}

// set session value
func (sess *WebsocketSession) Set(key interface{}, value interface{}) {
	sess.values.Store(key, value)
}

// get session value
func (sess *WebsocketSession) Get(key interface{}) interface{} {
	value, ok := sess.values.Load(key)
	if !ok {
		return nil
	}
	return value
}

// delete key
func (sess *WebsocketSession) Del(key interface{}) {
	sess.values.Delete(key)
}

// destroy session
func (sess *WebsocketSession) Destroy() {
	_ = sess.conn.Close()
	sess.sessionId = ""
	sess.values.Range(func(key, value interface{}) bool {
		sess.values.Delete(key)
		return true
	})
}

// get client connection
func (sess *WebsocketSession) GetConn() *websocket.Conn {
	return sess.conn
}

package camWebsocket

import (
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/base/camUtils"
	"sync"
)

// websocket session
type WebsocketSession struct {
	camStatics.SessionInterface

	sessionId string   // sessionId
	values    sync.Map // save session value
}

// new websocket session
func NewWebsocketSession() *WebsocketSession {
	sess := new(WebsocketSession)
	sess.sessionId = camUtils.String.UUID()
	sess.values = sync.Map{}
	return sess
}

// set sessionId
func (sess *WebsocketSession) SetSessionId(sessId string) {
	sess.sessionId = sessId
}

// get sessionId
func (sess *WebsocketSession) GetSessionId() string {
	return sess.sessionId
}

// set session value
func (sess *WebsocketSession) Set(key string, value interface{}) {
	sess.values.Store(key, value)
}

// get session value
func (sess *WebsocketSession) Get(key string) interface{} {
	value, ok := sess.values.Load(key)
	if !ok {
		return nil
	}
	return value
}

// delete key
func (sess *WebsocketSession) Del(key string) {
	sess.values.Delete(key)
}

// destroy session
func (sess *WebsocketSession) Destroy() {
	sess.sessionId = ""
	sess.values.Range(func(key, value interface{}) bool {
		sess.values.Delete(key)
		return true
	})
}

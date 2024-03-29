package camSocket

import (
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/base/camUtils"
	"sync"
)

type SocketSession struct {
	camStatics.SessionInterface

	sessionId string   // session id. Generate when new instance
	values    sync.Map // values
}

func NewSocketSession() *SocketSession {
	sess := new(SocketSession)
	sess.sessionId = camUtils.String.UUID()
	sess.values = sync.Map{}
	return sess
}

func (sess *SocketSession) GetSessionId() string {
	return sess.sessionId
}

// set session value
func (sess *SocketSession) Set(key string, value interface{}) {
	sess.values.Store(key, value)
}

// get session value
func (sess *SocketSession) Get(key string) interface{} {
	value, ok := sess.values.Load(key)
	if !ok {
		return nil
	}
	return value
}

// delete key
func (sess *SocketSession) Del(key string) {
	sess.values.Delete(key)
}

// destroy session
func (sess *SocketSession) Destroy() {
	sess.sessionId = ""
	sess.values.Range(func(key, value interface{}) bool {
		sess.values.Delete(key)
		return true
	})
}

package camSocket

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camUtils"
	"net"
	"sync"
)

type SocketSession struct {
	camBase.SessionInterface

	// Deprecated: remove on v0.5.0
	conn      net.Conn // socket connection
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
func (sess *SocketSession) Set(key interface{}, value interface{}) {
	sess.values.Store(key, value)
}

// get session value
func (sess *SocketSession) Get(key interface{}) interface{} {
	value, ok := sess.values.Load(key)
	if !ok {
		return nil
	}
	return value
}

// delete key
func (sess *SocketSession) Del(key interface{}) {
	sess.values.Delete(key)
}

// Deprecated: remove on v0.5.0
func (sess *SocketSession) SetConn(conn net.Conn) {
	sess.conn = conn
}

// Deprecated: remove on v0.5.0
func (sess *SocketSession) GetConn() net.Conn {
	return sess.conn
}

// destroy session
func (sess *SocketSession) Destroy() {
	_ = sess.conn.Close()
	sess.sessionId = ""
	sess.values.Range(func(key, value interface{}) bool {
		sess.values.Delete(key)
		return true
	})
}

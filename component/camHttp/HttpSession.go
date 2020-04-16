package camHttp

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/gorilla/sessions"
)

// session
type HttpSession struct {
	camBase.SessionInterface

	session *sessions.Session
}

func NewHttpSession(storeSession *sessions.Session) *HttpSession {
	session := new(HttpSession)
	session.session = storeSession
	return session
}

// set sessionId
func (sess *HttpSession) GetSessionId() string {
	return sess.session.ID
}

// set key-value
func (sess *HttpSession) Set(key interface{}, value interface{}) {
	sess.session.Values[key] = value
}

// get value by key
func (sess *HttpSession) Get(key interface{}) interface{} {
	value, has := sess.session.Values[key]
	if !has {
		return nil
	}
	return value
}

// del value by key
func (sess *HttpSession) Del(key interface{}) {
	sess.Set(key, nil)
}

// destroy session
func (sess *HttpSession) Destroy() {
	sess.session.Values = map[interface{}]interface{}{}
}

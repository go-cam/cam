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
func (session *HttpSession) GetSessionId() string {
	return session.session.ID
}

// set key-value
func (session *HttpSession) Set(key interface{}, value interface{}) {
	session.session.Values[key] = value
}

// get value by key
func (session *HttpSession) Get(key interface{}) interface{} {
	value, has := session.session.Values[key]
	if !has {
		return nil
	}
	return value
}

// destroy session
func (session *HttpSession) Destroy() {
	session.session.Values = map[interface{}]interface{}{}
}

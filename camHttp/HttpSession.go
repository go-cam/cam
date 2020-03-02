package camHttp

import (
	"github.com/go-cam/cam/camBase"
	"github.com/gorilla/sessions"
)

// session
type HttpSession struct {
	camBase.SessionInterface

	session *sessions.Session
}

func NewHttpSession(session *sessions.Session) *HttpSession {
	model := new(HttpSession)
	model.session = session
	return model
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

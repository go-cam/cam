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
func (model *HttpSession) GetSessionId() string {
	return model.session.ID
}

// set key-value
func (model *HttpSession) Set(key interface{}, value interface{}) {
	model.session.Values[key] = value
}

// get value by key
func (model *HttpSession) Get(key interface{}) interface{} {
	value, has := model.session.Values[key]
	if !has {
		return nil
	}
	return value
}

// destroy session
func (model *HttpSession) Destroy() {
	model.session.Values = map[interface{}]interface{}{}
}

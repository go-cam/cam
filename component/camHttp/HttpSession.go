package camHttp

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/gorilla/sessions"
)

// session
type HttpSession struct {
	camBase.SessionInterface
	// Deprecated
	session *sessions.Session

	id     string
	values map[string]interface{}
}

type HttpSessionInterface interface {
	// get sessionId
	GetSessionId() string
	// set key-value in session
	Set(key string, value interface{})
	// get value by key
	Get(key string) interface{}
	// delete value by key
	Del(key string)
	// get all values in session
	Values() map[string]interface{}
	// destroy session
	Destroy()
}

// Deprecated
func NewHttpSession(storeSession *sessions.Session) *HttpSession {
	session := new(HttpSession)
	session.session = storeSession
	return session
}

func newHttpSession(sessId string, values map[string]interface{}) *HttpSession {
	sess := new(HttpSession)
	sess.id = sessId
	sess.values = values
	return sess
}

// set sessionId
func (sess *HttpSession) GetSessionId() string {
	return sess.id
}

// set key-value
func (sess *HttpSession) Set(key string, value interface{}) {
	sess.values[key] = value
}

// get value by key
func (sess *HttpSession) Get(key string) interface{} {
	value, has := sess.values[key]
	if !has {
		return nil
	}
	return value
}

// del value by key
func (sess *HttpSession) Del(key string) {
	delete(sess.values, key)
}

func (sess *HttpSession) Values() map[string]interface{} {
	return sess.values
}

// destroy session
// TODO fix it
func (sess *HttpSession) Destroy() {
	//sess.session.Values = map[string]interface{}{}
}

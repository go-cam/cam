package camHttp

import (
	"github.com/go-cam/cam/base/camStatics"
)

// session
type HttpSession struct {
	camStatics.SessionInterface

	id             string
	values         map[string]interface{}
	destroyHandler func()
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

func newHttpSession(sessId string, values map[string]interface{}) *HttpSession {
	sess := new(HttpSession)
	sess.id = sessId
	sess.values = values
	sess.destroyHandler = func() {}
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
func (sess *HttpSession) Destroy() {
	sess.destroyHandler()
}

func (sess *HttpSession) SetDestroyHandler(handler func()) {
	sess.destroyHandler = handler
}

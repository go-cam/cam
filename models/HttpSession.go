package models

import (
	"github.com/cinling/cin/base"
	"github.com/gorilla/sessions"
)

// session
type HttpSession struct {
	base.SessionInterface

	session *sessions.Session
}

func NewHttpSession(session *sessions.Session) *HttpSession {
	model := new(HttpSession)
	model.session = session
	return model
}

// 获取 sessionId
func (model *HttpSession) GetSessionId() string {
	return model.session.ID
}

// 设置值
func (model *HttpSession) Set(key interface{}, value interface{}) {
	model.session.Values[key] = value
}

// 获取值
func (model *HttpSession) Get(key interface{}) interface{} {
	value, has := model.session.Values[key]
	if !has {
		return nil
	}
	return value
}

// 销毁session 清空 session 所有数据
func (model *HttpSession) Destroy() {
	model.session.Values = map[interface{}]interface{}{}
}

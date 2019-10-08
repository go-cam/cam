package models

import "cin/base"

// http 上下文对象
type HttpContext struct {
	BaseContext
	session *HttpSession
}

// 新建 http 上下文对象
func NewHttpContext(session *HttpSession) *HttpContext {
	model := new(HttpContext)
	model.session = session
	return model
}

// 获取session
func (model *HttpContext) GetSession() base.SessionInterface {
	return model.session
}
package models

import "cin/base"

// http 上下文对象
type Context struct {
	BaseContext
	session base.SessionInterface
}

// 新建 http 上下文对象
func NewContext(session base.SessionInterface) *Context {
	model := new(Context)
	model.session = session
	return model
}

// 获取session
func (model *Context) GetSession() base.SessionInterface {
	return model.session
}
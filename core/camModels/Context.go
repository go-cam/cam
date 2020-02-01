package camModels

import "github.com/go-cam/cam/core/camBase"

// http 上下文对象
type Context struct {
	BaseContext
	session camBase.SessionInterface
}

// 新建 http 上下文对象
func NewContext(session camBase.SessionInterface) *Context {
	model := new(Context)
	model.session = session
	return model
}

// 获取session
func (model *Context) GetSession() camBase.SessionInterface {
	return model.session
}

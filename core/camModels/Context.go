package camModels

import "github.com/go-cam/cam/core/camBase"

// context
type Context struct {
	BaseContext
	session camBase.SessionInterface
}

func NewContext(session camBase.SessionInterface) *Context {
	model := new(Context)
	model.session = session
	return model
}

// get session
func (model *Context) GetSession() camBase.SessionInterface {
	return model.session
}

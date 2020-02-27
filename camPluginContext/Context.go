package camPluginContext

import "github.com/go-cam/cam/camBase"

type Context struct {
	camBase.ContextInterface

	session camBase.SessionInterface
}

func (model *Context) SetSession(session camBase.SessionInterface) {
	model.session = session
}

func (model *Context) GetSession() camBase.SessionInterface {
	return model.session
}

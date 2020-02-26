package camModels

import "github.com/go-cam/cam/camBase"

type BaseContext struct {
	camBase.ContextInterface

	session camBase.SessionInterface
}

func (model *BaseContext) SetSession(session camBase.SessionInterface) {
	model.session = session
}

func (model *BaseContext) GetSession() camBase.SessionInterface {
	return model.session
}

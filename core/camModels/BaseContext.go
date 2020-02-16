package camModels

import "github.com/go-cam/cam/core/camBase"

// 上下文对象
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

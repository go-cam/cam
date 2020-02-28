package camBase

type Context struct {
	ContextInterface

	session SessionInterface
}

func (model *Context) SetSession(session SessionInterface) {
	model.session = session
}

func (model *Context) GetSession() SessionInterface {
	return model.session
}

package camContext

import "github.com/go-cam/cam/base/camBase"

type Context struct {
	sess camBase.SessionInterface
	res  []byte
	rec  camBase.RecoverInterface
}

func (ctx *Context) SetSession(session camBase.SessionInterface) {
	ctx.sess = session
}

func (ctx *Context) GetSession() camBase.SessionInterface {
	return ctx.sess
}

func (ctx *Context) Write(res []byte) {
	ctx.res = res
}

func (ctx *Context) Read() []byte {
	return ctx.res
}

func (ctx *Context) SetRecover(rec camBase.RecoverInterface) {
	ctx.rec = rec
}

func (ctx *Context) GetRecover() camBase.RecoverInterface {
	return ctx.rec
}

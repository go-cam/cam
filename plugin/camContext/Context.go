package camContext

import (
	"github.com/go-cam/cam/base/camStatics"
)

type Context struct {
	sess  camStatics.SessionInterface
	res   []byte
	rec   camStatics.RecoverInterface
	route string
}

func (ctx *Context) SetSession(session camStatics.SessionInterface) {
	ctx.sess = session
}

func (ctx *Context) GetSession() camStatics.SessionInterface {
	return ctx.sess
}

func (ctx *Context) Write(res []byte) {
	ctx.res = res
}

func (ctx *Context) Read() []byte {
	return ctx.res
}

func (ctx *Context) SetRecover(rec camStatics.RecoverInterface) {
	ctx.rec = rec
}

func (ctx *Context) GetRecover() camStatics.RecoverInterface {
	return ctx.rec
}

func (ctx *Context) SetRoute(route string) {
	ctx.route = route
}

func (ctx *Context) GetRoute() string {
	return ctx.route
}

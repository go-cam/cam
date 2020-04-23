package camHttp

import (
	"github.com/go-cam/cam/plugin/camContext"
	"net/http"
)

type HttpContext struct {
	camContext.Context
	responseWriter http.ResponseWriter
	request        *http.Request
	closeHandlers  []func()
}

func (ctx *HttpContext) SetHttpResponseWriter(responseWriter http.ResponseWriter) {
	ctx.responseWriter = responseWriter
}

func (ctx *HttpContext) GetHttpResponseWriter() http.ResponseWriter {
	return ctx.responseWriter
}

func (ctx *HttpContext) SetHttpRequest(request *http.Request) {
	ctx.request = request
}

func (ctx *HttpContext) GetHttpRequest() *http.Request {
	return ctx.request
}

func (ctx *HttpContext) CloseHandler(handler func()) {
	ctx.closeHandlers = append(ctx.closeHandlers, handler)
}

func (ctx *HttpContext) Close() {
	for _, handler := range ctx.closeHandlers {
		handler()
	}
}

package camHttp

import (
	"github.com/go-cam/cam/plugin/camContext"
	"net/http"
	"time"
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

func (ctx *HttpContext) GetCookie(name string) *http.Cookie {
	cookie, err := ctx.request.Cookie(name)
	if err != nil && err != http.ErrNoCookie {
		panic(err)
	}
	return cookie
}

func (ctx *HttpContext) SetCookie(cookie *http.Cookie) {
	http.SetCookie(ctx.responseWriter, cookie)
}

func (ctx *HttpContext) DelCookie(name string) {
	cookie := ctx.GetCookie(name)
	if cookie == nil {
		return
	}
	cookie.Expires = time.Now().Add(-time.Second)
}

func (ctx *HttpContext) SetCookieValue(name string, value string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	ctx.SetCookie(cookie)
}

func (ctx *HttpContext) GetCookieValue(name string) string {
	cookie := ctx.GetCookie(name)
	if cookie == nil {
		return ""
	}
	return cookie.Value
}

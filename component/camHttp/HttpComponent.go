package camHttp

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camUtils"
	"github.com/go-cam/cam/component"
	"github.com/go-cam/cam/plugin/camContext"
	"github.com/go-cam/cam/plugin/camMiddleware"
	"github.com/go-cam/cam/plugin/camRouter"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// http server component
type HttpComponent struct {
	component.Component
	camRouter.RouterPlugin
	camContext.ContextPlugin
	camMiddleware.MiddlewarePlugin

	config *HttpComponentConfig

	sessStore *SessionStoreManager
}

// init
func (comp *HttpComponent) Init(configI camBase.ComponentConfigInterface) {
	comp.Component.Init(configI)

	var ok bool
	comp.config, ok = configI.(*HttpComponentConfig)
	if !ok {
		camBase.App.Fatal("HttpComponent", "invalid config")
		return
	}
	comp.RouterPlugin.Init(&comp.config.RouterPluginConfig)
	comp.ContextPlugin.Init(&comp.config.ContextPluginConfig)
	comp.MiddlewarePlugin.Init(&comp.config.MiddlewarePluginConfig)
	comp.sessStore = NewSessionStoreManager(comp.config.getSessionStore(), comp.config.getSessionOption())
}

// start
func (comp *HttpComponent) Start() {
	comp.Component.Start()

	if !comp.config.TlsOnly {
		camBase.App.Trace("HttpComponent", "listen http://localhost:"+strconv.FormatUint(uint64(comp.config.Port), 10))
		go comp.listenAndServe()
	}
	if comp.config.IsTlsOn {
		camBase.App.Trace("HttpComponent", "listen https://localhost:"+strconv.FormatUint(uint64(comp.config.TlsPort), 10))
		go comp.listenAndServeTLS()
	}
}

// stop
func (comp *HttpComponent) Stop() {
	defer comp.Component.Stop()
}

// Receive http request, Call controller action, Send http response
func (comp *HttpComponent) handlerFunc(rw http.ResponseWriter, r *http.Request) {
	defer func() {
		if rec := recover(); rec != nil {
			comp.tryRecover(rw, r, rec)
		}
	}()

	route := comp.getRoute(r)

	// Deprecated: remove this block in v0.5.0  It's not support middleware
	// =========== START ===========
	handler := comp.getCustomRoute(route)
	if handler != nil {
		handler(rw, r)
		return
	}
	// =========== END ===========

	ctx := comp.newHttpContext(r, rw)
	defer func() { ctx.Close() }()
	comp.routeHandler(ctx, route)
}

// Handle route and set httpResponse
func (comp *HttpComponent) routeHandler(ctx camBase.HttpContextInterface, route string) {
	next := func() []byte {
		return comp.callNext(ctx, route)
	}
	res := comp.CallWithMiddleware(ctx, route, next)
	_, err := ctx.GetHttpResponseWriter().Write(res)
	if err != nil {
		panic(err)
	}
}

// call controller or custom route handler
func (comp *HttpComponent) callNext(ctx camBase.HttpContextInterface, route string) []byte {
	// Use custom handler first
	handler := comp.GetCustomHandler(route)
	if handler != nil {
		return handler(ctx)
	}

	// Use Controller and Action to handle route
	ctrl, action := comp.GetControllerAction(route)
	if ctrl == nil || action == nil {
		rw := ctx.GetHttpResponseWriter()
		camBase.App.Warn("HttpComponent", "404. Not found route: "+route)
		rw.WriteHeader(404)
		return nil
	}

	return comp.callNextControllerAction(ctx, ctrl, action)
}

// Encapsulate the flow of call control
func (comp *HttpComponent) callNextControllerAction(ctx camBase.HttpContextInterface, ctrl camBase.ControllerInterface, action camBase.ControllerActionInterface) []byte {
	rw := ctx.GetHttpResponseWriter()
	r := ctx.GetHttpRequest()
	var err error

	// Compatible. Remove on v0.5.0
	if httpCtrlI, ok := ctrl.(HttpControllerInterface); ok {
		httpCtrlI.setResponseWriterAndRequest(&rw, r)
	}

	values := comp.getRequestValues(ctx.GetHttpRequest())
	ctrl.SetContext(ctx)
	ctrl.SetSession(ctx.GetSession()) // Compatible. Remove on v0.5.0
	ctrl.SetValues(values)

	if !ctrl.BeforeAction(action) {
		rw.WriteHeader(400)
		_, err = rw.Write([]byte("invalid request"))
		if err != nil {
			panic(err)
		}
		return nil
	}
	action.Call()
	res := ctrl.AfterAction(action, ctx.Read())
	return res
}

// inject *http.Request and http.ResponseWriter into context
func (comp *HttpComponent) injectHttpInContext(ctx camBase.ContextInterface, responseWriter http.ResponseWriter, request *http.Request) {
	ctxHttp, ok := ctx.(camBase.HttpContextInterface)
	if !ok {
		return
	}
	ctxHttp.SetHttpResponseWriter(responseWriter)
	ctxHttp.SetHttpRequest(request)
}

// try to recover panic
func (comp *HttpComponent) tryRecover(rw http.ResponseWriter, r *http.Request, v interface{}) {
	rec, ok := v.(camBase.RecoverInterface)
	if !ok {
		comp.Recover(v)
		return
	}

	recoverRoute := comp.GetRecoverRoute()
	ctx := comp.newHttpContext(r, rw)
	ctx.SetRecover(rec)
	defer func() { ctx.Close() }()
	comp.routeHandler(ctx, recoverRoute)
}

// enable server
func (comp *HttpComponent) listenAndServe() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", comp.handlerFunc)
	server := &http.Server{
		Addr:    ":" + strconv.FormatUint(uint64(comp.config.Port), 10),
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// enable server with SSl
func (comp *HttpComponent) listenAndServeTLS() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", comp.handlerFunc)
	server := &http.Server{
		Addr:    ":" + strconv.FormatUint(uint64(comp.config.TlsPort), 10),
		Handler: mux,
	}
	err := server.ListenAndServeTLS(comp.config.TlsCertFile, comp.config.TlsKeyFile)
	if err != nil {
		panic(err)
	}
}

// get request params
func (comp *HttpComponent) getRequestValues(request *http.Request) map[string]interface{} {
	values := comp.getRequestValuesByUrl(request)

	contentType := request.Header.Get("Content-Type")
	if strings.Contains(contentType, "multipart/form-data") {
		newValues := comp.getRequestValuesByFormData(request)
		for key, value := range newValues {
			values[key] = value
		}
	}
	if strings.Contains(contentType, "application/json") {
		newValues := comp.getRequestValuesByJson(request)
		for key, value := range newValues {
			values[key] = value
		}
	}

	return values
}

// parse params from request url
func (comp *HttpComponent) getRequestValuesByUrl(request *http.Request) map[string]interface{} {
	values := map[string]interface{}{}
	_ = request.ParseForm()
	for key, value := range request.Form {
		values[key] = value
	}
	return values
}

// parse params from form data
func (comp *HttpComponent) getRequestValuesByFormData(request *http.Request) map[string]interface{} {
	values := map[string]interface{}{}
	_ = request.ParseMultipartForm(128 * 1024 * 1024)
	for key, stringList := range request.MultipartForm.Value {
		if len(stringList) == 1 {
			values[key] = stringList[0]
		} else {
			values[key] = stringList
		}
	}

	return values
}

// parse params from json
func (comp *HttpComponent) getRequestValuesByJson(request *http.Request) map[string]interface{} {
	var values map[string]interface{}
	bytes, _ := ioutil.ReadAll(request.Body)
	camUtils.Json.DecodeToObj(bytes, &values)
	return values
}

// get custom route handler
// Deprecated: remove on v0.5.0
func (comp *HttpComponent) getCustomRoute(route string) camBase.HttpRouteHandler {
	handler, has := comp.config.routeHandlerDict[route]
	if !has {
		return nil
	}
	return handler
}

// new HttpContext
func (comp *HttpComponent) newHttpContext(r *http.Request, rw http.ResponseWriter) camBase.HttpContextInterface {
	ctx := comp.NewContext()
	httpCtx, ok := ctx.(camBase.HttpContextInterface)
	if !ok {
		panic("invalid HttpContext struct. Must implements camBase.ContextHttpInterface")
	}
	httpCtx.SetHttpRequest(r)
	httpCtx.SetHttpResponseWriter(rw)
	sessI, err := comp.sessStore.GetSession(httpCtx)
	if err != nil {
		panic(err)
	}
	httpCtx.SetSession(sessI)
	httpCtx.SetRoute(comp.getRoute(r))
	httpCtx.CloseHandler(func() {
		err := comp.sessStore.store.Save(sessI)
		if err != nil {
			panic(err)
		}
	})

	return httpCtx
}

// Get route by request
func (comp *HttpComponent) getRoute(r *http.Request) string {
	route := ""
	url := r.URL.String()
	dirs := camUtils.Url.SplitUrl(url)
	dirLen := len(dirs)
	if dirLen == 1 {
		route = dirs[0]
	} else {
		route = dirs[0] + "/" + dirs[1]
	}
	if route == "" {
		return comp.config.DefaultRoute()
	}

	return route
}

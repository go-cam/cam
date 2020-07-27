package camWebsocket

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camUtils"
	"github.com/go-cam/cam/component"
	"github.com/go-cam/cam/plugin"
	"github.com/go-cam/cam/plugin/camContext"
	"github.com/go-cam/cam/plugin/camMiddleware"
	"github.com/go-cam/cam/plugin/camRouter"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

type WebsocketComponent struct {
	component.Component
	camRouter.RouterPlugin
	camContext.ContextPlugin
	camMiddleware.MiddlewarePlugin

	config *WebsocketComponentConfig
	// struct of http upgrade to websocket
	upgrader websocket.Upgrader
	// receive message parse handler
	recvMessageParseHandler plugin.RecvMessageParseHandler
}

// init
func (comp *WebsocketComponent) Init(configI camBase.ComponentConfigInterface) {
	comp.Component.Init(configI)

	var ok bool
	comp.config, ok = configI.(*WebsocketComponentConfig)
	if !ok {
		camBase.App.Fatal("WebsocketComponent", "invalid config")
		return
	}
	comp.upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	comp.recvMessageParseHandler = comp.config.GetRecvMessageParseHandler()
	comp.RouterPlugin.Init(&comp.config.RouterPluginConfig)
	comp.ContextPlugin.Init(&comp.config.ContextPluginConfig)
	comp.MiddlewarePlugin.Init(&comp.config.MiddlewarePluginConfig)
}

// start
func (comp *WebsocketComponent) Start() {
	comp.Component.Start()

	if !comp.config.TlsOnly {
		camBase.App.Trace("WebsocketComponent", "listen ws://:"+camUtils.C.Uint16ToString(comp.config.Port))
		go comp.listenAndServe()
	}
	if comp.config.IsTlsOn {
		camBase.App.Trace("WebsocketComponent", "listen wss://:"+camUtils.C.Uint16ToString(comp.config.TlsPort))
		go comp.listenAndServeTLS()
	}
}

// new connection
func (comp *WebsocketComponent) handlerFunc(w http.ResponseWriter, r *http.Request) {
	conn, err := comp.upgrader.Upgrade(w, r, nil)
	if conn == nil || err != nil {
		return
	}

	sess := NewWebsocketSession()
	sess.SetConn(conn)

	defer func() {
		sess.Destroy()
	}()

	for {
		var recv []byte
		var msgType int
		msgType, recv, err = conn.ReadMessage()
		if err != nil {
			break
		}

		if msgType == websocket.TextMessage || msgType == websocket.BinaryMessage {
			ctx := comp.newWebsocketContext(conn, recv, sess)
			msg := comp.recvMessageParseHandler(recv)
			ctx.SetMessage(msg)
			route := msg.Route
			if route == "" {
				route = comp.config.DefaultRoute()
			}
			comp.routeHandler(ctx, msg.Route, msg.Data)
		}
	}
}

// Handle route and set sendMessage
func (comp *WebsocketComponent) routeHandler(ctx WebsocketContextInterface, route string, values map[string]interface{}) {
	defer func() {
		if rec := recover(); rec != nil {
			comp.tryRecover(ctx, rec)
		}
	}()

	next := func() []byte {
		return comp.callNext(ctx, route, values)
	}
	res := comp.CallWithMiddleware(ctx, route, next)
	if err := ctx.GetConn().WriteMessage(websocket.TextMessage, res); err != nil {
		panic(err)
	}
}

// call Controller-Action or Custom-Handler
func (comp *WebsocketComponent) callNext(ctx WebsocketContextInterface, route string, values map[string]interface{}) []byte {
	handler := comp.getCustomHandler(route)
	if handler != nil {
		websocketSession, ok := ctx.GetSession().(*WebsocketSession)
		if !ok {
			panic("session cannot convert to *WebsocketSession")
		}
		return handler(websocketSession.GetConn())
	}

	customHandler := comp.GetCustomHandler(route)
	if customHandler != nil {
		return customHandler(ctx)
	}

	ctrl, action := comp.GetControllerAction(route)
	if ctrl == nil || action == nil {
		camBase.App.Warn("WebsocketComponent", "404. not found route: "+route)
		return nil
	}
	// init ctrl
	ctrl.Init()
	ctrl.SetContext(ctx)
	ctrl.SetSession(ctx.GetSession())
	ctrl.SetValues(values)
	if !ctrl.BeforeAction(action) {
		return []byte("illegal request")
	}
	action.Call()
	response := ctrl.AfterAction(action, ctx.Read())

	return response
}

func (comp *WebsocketComponent) tryRecover(oldCtx WebsocketContextInterface, v interface{}) {
	rec, ok := v.(camBase.RecoverInterface)
	if !ok {
		comp.Recover(v)
		return
	}

	recoverRoute := comp.GetRecoverRoute()
	ctx := comp.newWebsocketContext(oldCtx.GetConn(), nil, oldCtx.GetSession().(*WebsocketSession))
	ctx.SetMessage(oldCtx.GetMessage())
	ctx.SetRecover(rec)
	comp.routeHandler(ctx, recoverRoute, nil)
}

// enable server
func (comp *WebsocketComponent) listenAndServe() {
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
func (comp *WebsocketComponent) listenAndServeTLS() {
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

// get custom route handler
// Deprecated: remove on v0.5.0  it's not support Middleware
func (comp *WebsocketComponent) getCustomHandler(route string) camBase.WebsocketRouteHandler {
	handler, has := comp.config.routeHandlerDict[route]
	if !has {
		return nil
	}
	return handler
}

// new websocket context
func (comp *WebsocketComponent) newWebsocketContext(conn *websocket.Conn, recv []byte, sess *WebsocketSession) WebsocketContextInterface {
	ctxI := comp.NewContext()
	wsCtxI, ok := ctxI.(WebsocketContextInterface)
	if !ok {
		panic("invalid WebsocketContext struct. Must implements camWebsocket.WebsocketContextInterface")
	}
	wsCtxI.SetSession(sess)
	wsCtxI.SetConn(conn)
	wsCtxI.SetRecv(recv)
	return wsCtxI
}

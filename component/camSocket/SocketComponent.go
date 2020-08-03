package camSocket

import (
	"bufio"
	"fmt"
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/base/camUtils"
	"github.com/go-cam/cam/component"
	"github.com/go-cam/cam/plugin"
	"github.com/go-cam/cam/plugin/camContext"
	"github.com/go-cam/cam/plugin/camMiddleware"
	"github.com/go-cam/cam/plugin/camRouter"
	"net"
	"time"
)

type SocketComponent struct {
	component.Component
	camRouter.RouterPlugin
	camContext.ContextPlugin
	camMiddleware.MiddlewarePlugin

	config *SocketComponentConfig

	connHandler camStatics.SocketConnHandler
	// receive message parse handler
	recvMessageParseHandler plugin.RecvMessageParseHandler
}

// init config
func (comp *SocketComponent) Init(configI camStatics.ComponentConfigInterface) {
	comp.Component.Init(configI)

	var ok bool
	comp.config, ok = configI.(*SocketComponentConfig)
	if !ok {
		camStatics.App.Fatal("SocketComponent", "invalid config")
		return
	}

	comp.RouterPlugin.Init(&comp.config.RouterPluginConfig)
	comp.ContextPlugin.Init(&comp.config.ContextPluginConfig)
	comp.MiddlewarePlugin.Init(&comp.config.MiddlewarePluginConfig)

	comp.connHandler = comp.defaultConnHandler
	if comp.config.ConnHandler != nil {
		comp.connHandler = comp.config.ConnHandler
	}
	comp.recvMessageParseHandler = plugin.DefaultRecvToMessageHandler
}

// start
func (comp *SocketComponent) Start() {
	comp.Component.Start()
	camStatics.App.Trace("SocketComponent", "listen tcp://:"+camUtils.C.Uint16ToString(comp.config.Port))
	go comp.listenAndServe()
}

// stop
func (comp *SocketComponent) Stop() {
	defer comp.Component.Stop()
}

func (comp *SocketComponent) listenAndServe() {
	listener, err := net.Listen("tcp", ":"+camUtils.C.Uint16ToString(comp.config.Port))
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go comp.connHandler(conn)
	}
}

// get config
func (comp *SocketComponent) GetConfig() *SocketComponentConfig {
	return comp.config
}

// Default conn handler.
// Handling incoming links
func (comp *SocketComponent) defaultConnHandler(conn net.Conn) {
	defer func() {
		_ = conn.Close()
		if rec := recover(); rec != nil {
			camStatics.App.Fatal("SocketComponent.defaultConnHandler", fmt.Sprint(rec))
		}
	}()

	if comp.config.Trace {
		camStatics.App.Trace("SocketComponent.defaultConnHandler", "new connection: "+conn.RemoteAddr().String())
	}

	sess := NewSocketSession()
	sess.SetConn(conn)
	defer sess.Destroy()

	for {
		recv := comp.recv(conn)
		ctx := comp.newSocketContext(conn, recv, sess)
		msg := comp.recvMessageParseHandler(recv)
		ctx.SetMessage(msg)
		route := msg.Route
		if route == "" {
			route = comp.config.DefaultRoute()
		}
		comp.routeHandler(ctx, msg.Route, msg.Data)
	}
}

func (comp *SocketComponent) routeHandler(ctx SocketContextInterface, route string, values map[string]interface{}) {
	defer func() {
		if rec := recover(); rec != nil {
			comp.tryRecover(ctx, rec)
		}
	}()

	next := func() []byte {
		return comp.callNext(ctx, route, values)
	}
	res := comp.CallWithMiddleware(ctx, route, next)
	comp.send(ctx.GetConn(), res)
}

func (comp *SocketComponent) callNext(ctx SocketContextInterface, route string, values map[string]interface{}) []byte {
	customHandler := comp.GetCustomHandler(route)
	if customHandler != nil {
		return customHandler(ctx)
	}

	ctrl, action := comp.GetControllerAction(route)
	if ctrl == nil || action == nil {
		camStatics.App.Warn("SocketComponent", "404. not found route: "+route)
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

// get recv message
func (comp *SocketComponent) recv(conn net.Conn) []byte {
	var err error

	// set read deadline
	if err = conn.SetReadDeadline(time.Now().Add(comp.config.RecvTimeout)); err != nil {
		panic(err)
	}

	var recv []byte
	if recv, err = bufio.NewReader(conn).ReadBytes(comp.config.Etb); err != nil {
		panic(err)
	}
	recv = recv[:len(recv)-1]

	if uint64(len(recv)) > comp.config.RecvMaxLen {
		_ = conn.Close()
		panic("Receive data too length. Max is " + camUtils.C.Uint64ToString(comp.config.RecvMaxLen) + " bytes.")
	}

	// unset read deadline
	if err = conn.SetReadDeadline(time.Time{}); err != nil {
		panic(err)
	}

	return recv
}

// send message to client
func (comp *SocketComponent) send(conn net.Conn, send []byte) {
	var err error

	// check send size
	if uint64(len(send)) > comp.config.SendMaxLen {
		panic("Send data too length. Max is " + camUtils.C.Uint64ToString(comp.config.SendMaxLen) + " bytes.")
	}

	// set send deadline
	if err = conn.SetWriteDeadline(time.Now().Add(comp.config.SendTimeout)); err != nil {
		panic(err)
	}

	// send
	if _, err = conn.Write(append(send, comp.config.Etb)); err != nil {
		panic(err)
	}

	// unset send deadline
	if err = conn.SetWriteDeadline(time.Time{}); err != nil {
		panic(err)
	}
}

// try to recover
func (comp *SocketComponent) tryRecover(oldCtx SocketContextInterface, v interface{}) {
	rec, ok := v.(camStatics.RecoverInterface)
	if !ok {
		comp.Recover(v)
		return
	}

	recoverRoute := comp.GetRecoverRoute()
	ctx := comp.newSocketContext(oldCtx.GetConn(), nil, oldCtx.GetSession().(*SocketSession))
	ctx.SetMessage(oldCtx.GetMessage())
	ctx.SetRecover(rec)
	comp.routeHandler(ctx, recoverRoute, nil)
}

// new SocketContext
func (comp *SocketComponent) newSocketContext(conn net.Conn, recv []byte, sess *SocketSession) SocketContextInterface {
	ctx := comp.NewContext()
	socketCtxI, ok := ctx.(SocketContextInterface)
	if !ok {
		panic("invalid SocketContext struct. Must implements camSocket.SocketContextInterface")
	}
	socketCtxI.SetConn(conn)
	socketCtxI.SetRecv(recv)
	socketCtxI.SetSession(sess)
	return socketCtxI
}

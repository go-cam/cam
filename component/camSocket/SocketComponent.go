package camSocket

import (
	"bufio"
	"fmt"
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camUtils"
	"github.com/go-cam/cam/component"
	"github.com/go-cam/cam/plugin"
	"github.com/go-cam/cam/plugin/camPluginContext"
	"github.com/go-cam/cam/plugin/camPluginRouter"
	"net"
	"time"
)

type SocketComponent struct {
	component.Component
	camPluginRouter.RouterPlugin
	camPluginContext.ContextPlugin

	config *SocketComponentConfig

	connHandler         camBase.SocketConnHandler
	messageParseHandler camBase.MessageParseHandler
}

// init config
func (comp *SocketComponent) Init(configI camBase.ComponentConfigInterface) {
	comp.Component.Init(configI)

	var ok bool
	comp.config, ok = configI.(*SocketComponentConfig)
	if !ok {
		camBase.App.Fatal("SocketComponent", "invalid config")
		return
	}

	comp.connHandler = comp.defaultConnHandler
	if comp.config.ConnHandler != nil {
		comp.connHandler = comp.config.ConnHandler
	}
	comp.messageParseHandler = plugin.DefaultMessageParseHandler
	if comp.config.MessageParseHandler != nil {
		comp.messageParseHandler = comp.config.MessageParseHandler
	}
}

// start
func (comp *SocketComponent) Start() {
	comp.Component.Start()
	camBase.App.Trace("SocketComponent", "listen tcp://:"+camUtils.C.Uint16ToString(comp.config.Port))
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
			camBase.App.Fatal("SocketComponent.defaultConnHandler", fmt.Sprint(rec))
		}
	}()

	if comp.config.Trace {
		camBase.App.Trace("SocketComponent.defaultConnHandler", "new connection: "+conn.RemoteAddr().String())
	}

	sess := NewSocketSession(conn)
	defer sess.Destroy()

	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		panic("net.Conn can't trans to *net.TCPConn")
	}

	for {
		comp.recvAndSend(tcpConn, sess)
	}
}

// handling recv message and send message. Distribute message to route handler
func (comp *SocketComponent) recvAndSend(conn *net.TCPConn, sess *SocketSession) {
	defer func() {
		comp.Recover(recover())
	}()

	recv := comp.recv(conn)
	if comp.config.Trace {
		camBase.App.Trace(comp.Name()+" recv", string(recv))
	}

	controllerName, actionName, values := comp.messageParseHandler(recv)
	route := camUtils.Url.HumpToUrl(controllerName) + "/" + camUtils.Url.HumpToUrl(actionName)

	controller, action := comp.GetControllerAction(route)
	if controller == nil || action == nil {
		camBase.App.Warn("WebsocketComponent", "404. not found route: "+route)
		return
	}

	ctx := comp.NewContext()

	controller.Init()
	controller.SetContext(ctx)
	controller.SetSession(sess)
	controller.SetValues(values)

	if !controller.BeforeAction(action) {
		camBase.App.Warn("WebsocketComponent", "BeforeAction: invalid request.")
		return
	}
	action.Call()
	send := controller.AfterAction(action, controller.GetResponse())

	if comp.config.Trace {
		camBase.App.Trace(comp.Name()+" send", string(send))
	}
	comp.send(conn, send)
}

// get recv message
func (comp *SocketComponent) recv(conn *net.TCPConn) []byte {
	var err error

	// set read deadline
	if err = conn.SetReadDeadline(time.Now().Add(comp.config.RecvTimeout)); err != nil {
		panic(err)
	}

	var recv []byte
	if recv, err = bufio.NewReader(conn).ReadBytes(comp.config.Etb); err != nil {
		panic(err)
	}

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
func (comp *SocketComponent) send(conn *net.TCPConn, send []byte) {
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
	if _, err = conn.Write(send); err != nil {
		panic(err)
	}

	// unset send deadline
	if err = conn.SetReadDeadline(time.Time{}); err != nil {
		panic(err)
	}
}

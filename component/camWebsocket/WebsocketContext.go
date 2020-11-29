package camWebsocket

import (
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/base/camStructs"
	"github.com/go-cam/cam/plugin/camContext"
	"github.com/gorilla/websocket"
)

// websocket context
type WebsocketContextInterface interface {
	camStatics.ContextInterface
	camContext.MessageContextInterface
	SetConn(conn *websocket.Conn)
	GetConn() *websocket.Conn
	SetRecv(recv []byte)
	GetRecv() []byte
}

type WebsocketContext struct {
	camContext.Context
	camContext.MessageContext
	conn *websocket.Conn
	recv []byte
	msg  *camStructs.RecvMessage
}

func (ctx *WebsocketContext) SetConn(conn *websocket.Conn) {
	ctx.conn = conn
}

func (ctx *WebsocketContext) GetConn() *websocket.Conn {
	return ctx.conn
}

func (ctx *WebsocketContext) SetRecv(recv []byte) {
	ctx.recv = recv
}

func (ctx *WebsocketContext) GetRecv() []byte {
	return ctx.recv
}

func (ctx *WebsocketContext) SetMessage(msg *camStructs.RecvMessage) {
	ctx.msg = msg
}

func (ctx *WebsocketContext) GetMessage() *camStructs.RecvMessage {
	return ctx.msg
}

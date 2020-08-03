package camSocket

import (
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/plugin/camContext"
	"net"
)

type SocketContextInterface interface {
	camStatics.ContextInterface
	camContext.MessageContextInterface
	SetConn(conn net.Conn)
	GetConn() net.Conn
	SetRecv(recv []byte)
	GetRecv() []byte
}

type SocketContext struct {
	camContext.Context
	camContext.MessageContext
	conn net.Conn
	recv []byte
}

func (ctx *SocketContext) SetConn(conn net.Conn) {
	ctx.conn = conn
}

func (ctx *SocketContext) GetConn() net.Conn {
	return ctx.conn
}

func (ctx *SocketContext) SetRecv(recv []byte) {
	ctx.recv = recv
}

func (ctx *SocketContext) GetRecv() []byte {
	return ctx.recv
}

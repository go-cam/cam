package camContext

import "github.com/go-cam/cam/base/camStructs"

type MessageContextInterface interface {
	// return by plugin.RecvMessageParseHandler
	// SEE: WebsocketComponent.recvMessageParseHandler
	SetMessage(msg *camStructs.RecvMessage)
	GetMessage() *camStructs.RecvMessage
}

type MessageContext struct {
	msg *camStructs.RecvMessage
}

func (ctx *MessageContext) SetMessage(msg *camStructs.RecvMessage) {
	ctx.msg = msg
}

func (ctx *MessageContext) GetMessage() *camStructs.RecvMessage {
	return ctx.msg
}

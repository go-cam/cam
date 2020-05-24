package camContext

import "github.com/go-cam/cam/base/camStructs"

type MessageContextInterface interface {
	// return by plugin.RecvMessageParseHandler
	// SEE: WebsocketComponent.recvMessageParseHandler
	SetMessage(msg *camStructs.Message)
	GetMessage() *camStructs.Message
}

type MessageContext struct {
	msg *camStructs.Message
}

func (ctx *MessageContext) SetMessage(msg *camStructs.Message) {
	ctx.msg = msg
}

func (ctx *MessageContext) GetMessage() *camStructs.Message {
	return ctx.msg
}

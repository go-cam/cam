package plugin

import (
	"github.com/go-cam/cam/base/camStructs"
	"github.com/go-cam/cam/base/camUtils"
)

// recvMessage parse handler
type RecvMessageParseHandler func(recv []byte) *camStructs.Message

// default RecvMessageParseHandler
func DefaultRecvMessageParseHandler(bytes []byte) *camStructs.Message {
	if bytes == nil {
		return nil
	}
	msgStruct := new(camStructs.Message)
	camUtils.Json.DecodeToObj(bytes, msgStruct)
	return msgStruct
}

// sendMessage parse handler
type SendMessageParseHandler func(msg *camStructs.Message, response []byte) []byte

// default SendMessageParseHandler
func DefaultSendMessageParseHandler(msg *camStructs.Message, response []byte) []byte {
	camUtils.Json.DecodeToObj(response, msg.Data)
	return camUtils.Json.Encode(msg)
}

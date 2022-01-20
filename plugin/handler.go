package plugin

import (
	"github.com/go-cam/cam/base/camStructs"
	"github.com/go-cam/cam/base/camUtils"
)

// recvMessage parse handler
type RecvMessageParseHandler func(recv []byte) *camStructs.RecvMessage

// default RecvMessageParseHandler
func DefaultRecvMessageParseHandler(bytes []byte) *camStructs.RecvMessage {
	if bytes == nil {
		return nil
	}
	msgStruct := new(camStructs.RecvMessage)
	camUtils.Json.DecodeToObj(bytes, msgStruct)
	return msgStruct
}

// sendMessage parse handler
type SendMessageParseHandler func(msg *camStructs.SendMessage) []byte

// default SendMessageParseHandler
func DefaultSendMessageParseHandler(msg *camStructs.SendMessage) []byte {
	return camUtils.Json.Encode(map[string]interface{}{
		"i": msg.Id,
		"r": msg.Route,
		"d": string(msg.Data.([]byte)), // use Character stream
	})
}

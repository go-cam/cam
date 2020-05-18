package plugin

import (
	"github.com/go-cam/cam/base/camStructs"
	"github.com/go-cam/cam/base/camUtils"
)

// recvMessage parse handler
type RecvMessageParseHandler func(recv []byte) *camStructs.Message

// default RecvMessageParseHandler
func DefaultRecvToMessageHandler(recv []byte) *camStructs.Message {
	if recv == nil {
		return nil
	}
	msgStruct := new(camStructs.Message)
	camUtils.Json.DecodeToObj(recv, msgStruct)
	return msgStruct
}

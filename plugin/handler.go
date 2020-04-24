package plugin

import (
	"github.com/go-cam/cam/base/camStructs"
	"github.com/go-cam/cam/base/camUtils"
)

// recvMessage parse handler
type RecvMessageParseHandler func(recv []byte) (message *camStructs.Message, values map[string]interface{})

// default RecvMessageParseHandler
func DefaultRecvToMessageHandler(recv []byte) (*camStructs.Message, map[string]interface{}) {
	if recv == nil {
		return nil, nil
	}
	msgStruct := new(camStructs.Message)
	resStruct := new(camStructs.Response)
	camUtils.Json.DecodeToObj(recv, msgStruct)
	camUtils.Json.DecodeToObj([]byte(msgStruct.Data), resStruct)
	return msgStruct, resStruct.Values
}

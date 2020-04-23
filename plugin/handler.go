package plugin

import (
	"github.com/go-cam/cam/base/camStructs"
	"github.com/go-cam/cam/base/camUtils"
	"strings"
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

// default handler.
// Parse the received data to: controllerName、actionName、values
// only on controller action mode
// Deprecated: remove on v0.5.0
func DefaultMessageParseHandler(message []byte) (controllerName string, actionName string, values map[string]interface{}) {
	messageModel := new(camStructs.Message)
	responseModel := new(camStructs.Response)
	camUtils.Json.DecodeToObj(message, messageModel)
	camUtils.Json.DecodeToObj([]byte(messageModel.Data), responseModel)

	if messageModel.Route == "" {
		return "", "", responseModel.Values
	}
	if !strings.Contains(messageModel.Route, "/") {
		return messageModel.Route, "", responseModel.Values
	}
	tmpArr := strings.Split(messageModel.Route, "/")
	return camUtils.Url.UrlToHump(tmpArr[0]), camUtils.Url.UrlToHump(tmpArr[1]), responseModel.Values
}

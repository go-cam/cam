package camBase

import (
	"github.com/gorilla/websocket"
	"net/http"
)

// app status
type ApplicationStatus int

// CamModule type
type CamModuleType int

// websocket component message parse handler
type WebsocketMessageParseHandler func(message []byte) (controllerName string, actionName string, values map[string]interface{})

// log level
type LogLevel uint8

// http custom route handler
type HttpRouteHandler func(responseWriter http.ResponseWriter, request *http.Request)

// websocket custom route handler
type WebsocketRouteHandler func(conn *websocket.Conn) []byte

// recover handler result
type RecoverHandlerResult uint8

// component recover handler
type RecoverHandler func(rec interface{}) RecoverHandlerResult

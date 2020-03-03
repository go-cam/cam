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
type LogLevel int

// http custom route handler
type HttpRouteHandler func(responseWriter http.ResponseWriter, request *http.Request)

// websocket custom route handler
type WebsocketRouteHandler func(conn *websocket.Conn) []byte

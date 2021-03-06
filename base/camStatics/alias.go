package camStatics

import (
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"reflect"
)

// app status
type ApplicationStatus int

// CamModule type
type CamModuleType int

// message parse handler.
// it can read route and values info form the message
type MessageParseHandler func(message []byte) (controllerName string, actionName string, values map[string]interface{})

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

// socket conn handler
type SocketConnHandler func(conn net.Conn)

// socket custom route handler
type SocketRouteHandler func(conn net.Conn) []byte

// valid mode priority level
type ValidMode uint8

// valid handler
type ValidHandler func(value reflect.Value) error

// the next Handler.
// The processing method of Middleware
type NextHandler func() []byte

// route Handler
// Such as: Custom Handler, ControllerAction Handler
type RouteHandler func(ctxI ContextInterface) []byte

// load balancing login
type GRpcLoadBalancingLogic int

// Mysql column type
type MysqlColumnType string

// Mysql nullable type
type MysqlNullableType int

// Mysql key type
// Example: Primary Key, Index, Unique
type MysqlKeyType string

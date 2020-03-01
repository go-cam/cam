package camBase

// app status
type ApplicationStatus uint8

// CamModule type
type CamModuleType uint8

// websocket component message parse handler
type WebsocketComponentMessageParseHandler func(message []byte) (controllerName string, actionName string, values map[string]interface{})

// log level
type LogLevel int

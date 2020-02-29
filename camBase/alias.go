package camBase

type ApplicationStatus uint8 // app status
type CamModuleType uint8     // CamModule type
type WebsocketComponentMessageParseHandler func(message []byte) (controllerName string, actionName string, values map[string]interface{})

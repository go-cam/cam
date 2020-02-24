package camConfigs

import (
	"github.com/go-cam/cam/camBase"
)

// router plugin.
// It save controller and action config
type PluginRouter struct {
	ControllerList            []camBase.ControllerInterface                           // http or websocket controller list
	ConsoleControllerList     []camBase.ControllerInterface                           // console controller list
	OnWebsocketMessageHandler func(conn camBase.ContextInterface, recvMessage []byte) // on websocket receive message
}

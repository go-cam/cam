package camConfigs

import (
	"github.com/go-cam/cam/camBase"
)

// router plugin.
// It save controller and action config
type RouterPlugin struct {
	ControllerList []camBase.ControllerInterface // controller list
	// Deprecated:
	ConsoleControllerList     []camBase.ControllerInterface                           // console controller list
	OnWebsocketMessageHandler func(conn camBase.ContextInterface, recvMessage []byte) // on websocket receive message
}

func (plugin *RouterPlugin) Init() {
	plugin.ControllerList = []camBase.ControllerInterface{}
	plugin.OnWebsocketMessageHandler = nil
}

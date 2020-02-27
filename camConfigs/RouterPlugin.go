package camConfigs

import (
	"github.com/go-cam/cam/camBase"
)

// router plugin.
// It save controller and action config
type RouterPlugin struct {
	ControllerList []camBase.ControllerBakInterface // controller list
	// Deprecated:
	ConsoleControllerList     []camBase.ControllerBakInterface                        // console controller list
	OnWebsocketMessageHandler func(conn camBase.ContextInterface, recvMessage []byte) // on websocket receive message
}

func (plugin *RouterPlugin) Init() {
	plugin.ControllerList = []camBase.ControllerBakInterface{}
	plugin.OnWebsocketMessageHandler = nil
}

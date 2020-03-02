package camPluginRouter

import "github.com/go-cam/cam/base/camBase"

// router plugin.
// it can inject component if it need to
type RouterPluginConfig struct {
	camBase.PluginConfigInterface
	ControllerList []camBase.ControllerInterface // controller list
}

// init
func (config *RouterPluginConfig) Init() {
	config.ControllerList = []camBase.ControllerInterface{}
}

// register controller
// controller: Inout &Controller{} OR new(Controller) OR Controller{} is ok
func (config *RouterPluginConfig) Register(controller camBase.ControllerInterface) {
	config.ControllerList = append(config.ControllerList, controller)
}

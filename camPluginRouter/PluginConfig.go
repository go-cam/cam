package camPluginRouter

import "github.com/go-cam/cam/camBase"

// router plugin.
// it can inject component if it need to
type PluginConfig struct {
	camBase.PluginConfigInterface
	ControllerList []camBase.ControllerInterface // controller list
}

// init
func (config *PluginConfig) Init() {
	config.ControllerList = []camBase.ControllerInterface{}
}

// register controller
// controller: Inout &Controller{} OR new(Controller) OR Controller{} is ok
func (config *PluginConfig) Register(controller camBase.ControllerInterface) {
	config.ControllerList = append(config.ControllerList, controller)
}

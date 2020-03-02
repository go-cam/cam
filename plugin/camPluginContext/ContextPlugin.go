package camPluginContext

import (
	"github.com/go-cam/cam/base/camBase"
	"reflect"
)

type ContextPlugin struct {
	camBase.PluginInterface

	config *ContextPluginConfig
}

func (plugin *ContextPlugin) Init(config *ContextPluginConfig) {
	plugin.config = config
}

// new context by type
func (plugin *ContextPlugin) NewContext() camBase.ContextInterface {
	contextValue := reflect.New(plugin.config.contextType)
	return contextValue.Interface().(camBase.ContextInterface)
}

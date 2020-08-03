package camContext

import (
	"github.com/go-cam/cam/base/camStatics"
	"reflect"
)

type ContextPlugin struct {
	camStatics.PluginInterface

	config *ContextPluginConfig
}

func (plugin *ContextPlugin) Init(config *ContextPluginConfig) {
	plugin.config = config
}

// new context by type
func (plugin *ContextPlugin) NewContext() camStatics.ContextInterface {
	contextValue := reflect.New(plugin.config.contextType)
	return contextValue.Interface().(camStatics.ContextInterface)
}

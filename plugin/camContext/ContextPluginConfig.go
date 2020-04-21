package camContext

import (
	"github.com/go-cam/cam/base/camBase"
	"reflect"
)

type ContextPluginConfig struct {
	camBase.PluginConfigInterface

	contextType reflect.Type
}

func (config *ContextPluginConfig) Init() {
	config.SetContextStruct(&Context{})
}

// set Context type
func (config *ContextPluginConfig) SetContextStruct(v camBase.ContextInterface) {
	config.contextType = reflect.TypeOf(v)
	if config.contextType.Kind() == reflect.Ptr {
		config.contextType = config.contextType.Elem()
	}
}

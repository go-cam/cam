package camContext

import (
	"github.com/go-cam/cam/base/camStatics"
	"reflect"
)

type ContextPluginConfig struct {
	camStatics.PluginConfigInterface

	contextType reflect.Type
}

func (config *ContextPluginConfig) Init() {
	config.SetContextStruct(&Context{})
}

// set Context type
func (config *ContextPluginConfig) SetContextStruct(v camStatics.ContextInterface) {
	config.contextType = reflect.TypeOf(v)
	if config.contextType.Kind() == reflect.Ptr {
		config.contextType = config.contextType.Elem()
	}
}

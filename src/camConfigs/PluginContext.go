package camConfigs

import (
	"github.com/go-cam/cam/camBase"
	"github.com/go-cam/cam/camModels"
	"reflect"
)

// HTTP and websocket service custom context
type PluginContext struct {
	contextType reflect.Type
}

func (plugin *PluginContext) InitPluginContext() {
	plugin.SetContextStruct(new(camModels.BaseContext))
}

func (plugin *PluginContext) SetContextStruct(v interface{}) {
	plugin.contextType = reflect.TypeOf(v)
	if plugin.contextType.Kind() == reflect.Ptr {
		plugin.contextType = plugin.contextType.Elem()
	}
	contextValue := reflect.New(plugin.contextType)
	_, ok := contextValue.Interface().(camBase.ContextInterface)
	if !ok {
		panic("context struct must implements camBase.ContextInterface")
	}
}

func (plugin *PluginContext) NewContext() camBase.ContextInterface {
	contextValue := reflect.New(plugin.contextType)
	return contextValue.Interface().(camBase.ContextInterface)
}

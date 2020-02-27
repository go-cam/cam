package camConfigs

import (
	"github.com/go-cam/cam/camBase"
	"github.com/go-cam/cam/camModels"
	"reflect"
)

// HTTP and websocket service custom context
// Deprecated: remove on v0.3.0
type ContextPlugin struct {
	contextType reflect.Type
}

func (plugin *ContextPlugin) InitContextPlugin() {
	plugin.SetContextStruct(new(camModels.BaseContext))
}

func (plugin *ContextPlugin) SetContextStruct(v interface{}) {
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

func (plugin *ContextPlugin) NewContext() camBase.ContextInterface {
	contextValue := reflect.New(plugin.contextType)
	return contextValue.Interface().(camBase.ContextInterface)
}
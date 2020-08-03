package component

import (
	"github.com/go-cam/cam/base/camStatics"
	"reflect"
)

// base config
type ComponentConfig struct {
	camStatics.ComponentConfigInterface
	Component      camStatics.ComponentInterface // Instance of corresponding component
	recoverHandler camStatics.RecoverHandler
}

// get component instance
func (config *ComponentConfig) NewComponent() camStatics.ComponentInterface {
	if config.Component == nil {
		return nil
	}

	t := reflect.TypeOf(config.Component)
	componentType := t.Elem()
	componentValue := reflect.New(componentType)
	componentI := componentValue.Interface().(camStatics.ComponentInterface)

	return componentI
}

// get recover handler
func (config *ComponentConfig) GetRecoverHandler() camStatics.RecoverHandler {
	return config.recoverHandler
}

// set recover handler
// It can recover panic, but not all, only component release.
// For example, the panic thrown by the controller can be handled
func (config *ComponentConfig) RecoverHandler(handler camStatics.RecoverHandler) {
	config.recoverHandler = handler
}

// init all pluginConfig in configInterface
//func (config *ComponentConfig) InitPlugin(configI ComponentConfigInterface) {
//	rValue := camUtils.Reflect.ValueOfElem(configI)
//
//	num := rValue.NumField()
//	pluginIType := reflect.TypeOf((*PluginConfigInterface)(nil))
//	for i := 0; i < num; i++ {
//		fieldValue := rValue.Field(i)
//		if !fieldValue.CanInterface() {
//			continue
//		}
//		fieldI := fieldValue.Interface()
//		isImp := reflect.TypeOf(fieldI).Implements(pluginIType)
//		if !isImp {
//			continue
//		}
//		pluginConfigI, ok := fieldI.(PluginConfigInterface)
//		if !ok {
//			continue
//		}
//
//		pluginConfigI.Init()
//	}
//}

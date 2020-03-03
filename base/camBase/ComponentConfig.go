package camBase

import (
	"reflect"
)

// base config
type ComponentConfig struct {
	ComponentConfigInterface
	Component ComponentInterface // Instance of corresponding component
}

// get component instance
func (config *ComponentConfig) NewComponent() ComponentInterface {
	if config.Component == nil {
		return nil
	}

	t := reflect.TypeOf(config.Component)
	componentType := t.Elem()
	componentValue := reflect.New(componentType)
	componentI := componentValue.Interface().(ComponentInterface)

	return componentI
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

package camBase

import (
	"reflect"
)

// base component struct
type Component struct {
	ComponentInterface

	// component name
	name string
}

// init config
func (component *Component) Init(configInterface ComponentConfigInterface) {
	component.name = component.getComponentName(configInterface.NewComponent())
}

// start
func (component *Component) Start() {

}

// stop
func (component *Component) Stop() {

}

// get component struct name
func (component *Component) getComponentName(componentInterface ComponentInterface) string {
	t := reflect.TypeOf(componentInterface)
	return t.Elem().Name()
}

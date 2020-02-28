package camBase

import (
	"reflect"
)

// base component struct
type Component struct {
	ComponentInterface

	// component name
	name string
	// App instance
	App ApplicationInterface
}

// set App instance
func (component *Component) SetApp(app ApplicationInterface) {
	component.App = app
}

// init config
func (component *Component) Init(configInterface ConfigComponentInterface) {
	component.name = component.getComponentName(configInterface.GetComponent())
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

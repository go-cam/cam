package camComponents

import (
	"github.com/go-cam/cam/camBase"
	"reflect"
)

// base component struct
type Base struct {
	camBase.ComponentInterface

	// component name
	name string
	// app instance
	app camBase.ApplicationInterface
}

// set app instance
func (component *Base) SetApp(app camBase.ApplicationInterface) {
	component.app = app
}

// init config
func (component *Base) Init(configInterface camBase.ConfigComponentInterface) {
	component.name = component.getComponentName(configInterface.GetComponent())
}

// start
func (component *Base) Start() {

}

// stop
func (component *Base) Stop() {

}

// get component struct name
func (component *Base) getComponentName(componentInterface camBase.ComponentInterface) string {
	t := reflect.TypeOf(componentInterface)
	return t.Elem().Name()
}

package component

import (
	"github.com/go-cam/cam/base/camStatics"
	"reflect"
)

// base component struct
type Component struct {
	camStatics.ComponentInterface

	// component name
	name string
	// recover handler
	recoverHandler camStatics.RecoverHandler
}

// init config
func (comp *Component) Init(configInterface camStatics.ComponentConfigInterface) {
	comp.name = comp.getComponentName(configInterface.NewComponent())
	comp.recoverHandler = configInterface.GetRecoverHandler()
	if comp.recoverHandler == nil {
		comp.recoverHandler = comp.defaultRecoverHandler
	}
}

// start
func (comp *Component) Start() {

}

// stop
func (comp *Component) Stop() {

}

// get component struct name
func (comp *Component) getComponentName(componentInterface camStatics.ComponentInterface) string {
	t := reflect.TypeOf(componentInterface)
	return t.Elem().Name()
}

// component name
func (comp *Component) Name() string {
	return comp.name
}

// recover
func (comp *Component) Recover(rec interface{}) {
	if comp.recoverHandler(rec) == camStatics.RecoverHandlerResultPanic {
		panic(rec)
	}
}

// default recover handler
func (comp *Component) defaultRecoverHandler(rec interface{}) camStatics.RecoverHandlerResult {
	return camStatics.RecoverHandlerResultPanic
}

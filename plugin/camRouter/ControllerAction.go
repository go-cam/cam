package camRouter

import (
	"github.com/go-cam/cam/base/camStatics"
	"reflect"
)

// controller action
type ControllerAction struct {
	camStatics.ControllerActionInterface

	route  string
	rValue *reflect.Value
	isCall bool
}

// new action
func NewControllerAction(route string, rValue *reflect.Value) *ControllerAction {
	action := new(ControllerAction)
	action.route = route
	action.rValue = rValue
	return action
}

// action route
func (action *ControllerAction) Route() string {
	return action.route
}

// call action. only call once
func (action *ControllerAction) Call() {
	if action.isCall {
		return
	}
	action.isCall = true
	_ = action.rValue.Call([]reflect.Value{})
}

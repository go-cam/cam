package camRouter

import (
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/base/camUtils"
)

// base controller
type Controller struct {
	camStatics.ControllerInterface

	context camStatics.ContextInterface

	values        map[string]interface{}      // controller values
	responseBytes []byte                      // response bytes
	recover       camStatics.RecoverInterface // get recover
}

// OVERWRITE:
func (ctrl *Controller) Init() {
	ctrl.values = map[string]interface{}{}
	ctrl.responseBytes = []byte("")
}

// OVERWRITE
func (ctrl *Controller) BeforeAction(action camStatics.ControllerActionInterface) bool {
	return true
}

// OVERWRITE
func (ctrl *Controller) AfterAction(action camStatics.ControllerActionInterface, response []byte) []byte {
	return response
}

// OVERWRITE
func (ctrl *Controller) SetContext(context camStatics.ContextInterface) {
	ctrl.context = context
}

// OVERWRITE
func (ctrl *Controller) GetContext() camStatics.ContextInterface {
	return ctrl.context
}

// OVERWRITE
// Deprecated: remove on v0.5.0
// Instead: GetContext().SetSession();
func (ctrl *Controller) SetSession(session camStatics.SessionInterface) {
	ctrl.context.SetSession(session)
}

// OVERWRITE
// Deprecated: remove on v0.5.0
// Instead: GetContext().GetSession();
func (ctrl *Controller) GetSession() camStatics.SessionInterface {
	return ctrl.context.GetSession()
}

// OVERWRITE
// set values
func (ctrl *Controller) SetValues(values map[string]interface{}) {
	ctrl.values = values
}

// get all values
func (ctrl *Controller) GetValues() map[string]interface{} {
	return ctrl.values
}

// OVERWRITE
func (ctrl *Controller) AddValue(key string, value interface{}) {
	ctrl.values[key] = value
}

// get value by key
func (ctrl *Controller) GetValue(key string) interface{} {
	value, has := ctrl.values[key]
	if !has {
		value = nil
	}
	return value
}

// get the value with string type
func (ctrl *Controller) GetString(key string) string {
	return camUtils.C.InterfaceToString(ctrl.GetValue(key))
}

// set response content
// Deprecated: remove on v0.5.0
// Instead: GetContext().Write();
func (ctrl *Controller) SetResponse(bytes []byte) {
	ctrl.responseBytes = bytes
	ctrl.GetContext().Write(bytes)
}

// OVERWRITE
// return action write
// Deprecated: remove on v0.5.0
// Instead: GetContext().Read();
func (ctrl *Controller) GetResponse() []byte {
	return ctrl.responseBytes
}

// OVERWRITE
func (ctrl *Controller) GetDefaultActionName() string {
	return ""
}

// set recover
// Deprecated
func (ctrl *Controller) SetRecover(rec camStatics.RecoverInterface) {
	ctrl.GetContext().SetRecover(rec)
}

// get recover
func (ctrl *Controller) GetRecover() camStatics.RecoverInterface {
	return ctrl.GetContext().GetRecover()
}

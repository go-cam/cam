package camBase

import (
	"net/http"
)

// base controller
type Controller struct {
	ControllerInterface

	app     ApplicationInterface // app instance
	context ContextInterface

	values            map[string]interface{} // controller values
	responseBytes     []byte                 // response bytes
	defaultActionName string

	httpResponseWriter http.ResponseWriter
	httpRequest        *http.Request
}

// OVERWRITE:
func (controller *Controller) Init() {
	controller.values = map[string]interface{}{}
	controller.responseBytes = []byte("")
	controller.defaultActionName = ""
}

// OVERWRITE
func (controller *Controller) BeforeAction(action ControllerActionInterface) bool {
	return true
}

// OVERWRITE
func (controller *Controller) AfterAction(action ControllerActionInterface, response []byte) []byte {
	return response
}

// OVERWRITE
func (controller *Controller) SetContext(context ContextInterface) {
	controller.context = context
}

// OVERWRITE
func (controller *Controller) GetContext() ContextInterface {
	return controller.context
}

// OVERWRITE
// set values
func (controller *Controller) SetValues(values map[string]interface{}) {
	controller.values = values
}

// get all values
func (controller *Controller) GetValues() map[string]interface{} {
	return controller.values
}

// OVERWRITE
func (controller *Controller) AddValue(key string, value interface{}) {
	controller.values[key] = value
}

// get value by key
func (controller *Controller) GetValue(key string) interface{} {
	value, has := controller.values[key]
	if !has {
		value = nil
	}
	return value
}

// OVERWRITE
// set app instance
func (controller *Controller) SetApp(app ApplicationInterface) {
	controller.app = app
}

// Return app instance
func (controller *Controller) GetApp() ApplicationInterface {
	return controller.app
}

// set response content
func (controller *Controller) SetResponse(bytes []byte) {
	controller.responseBytes = bytes
}

// OVERWRITE
// return action write
func (controller *Controller) GetResponse() []byte {
	return controller.responseBytes
}

func (controller *Controller) SetDefaultActionName(actionName string) {
	controller.defaultActionName = actionName
}

// OVERWRITE
func (controller *Controller) GetDefaultActionName() string {
	return controller.defaultActionName
}

// Only support on http request
func (controller *Controller) GetHttpResponseWrite() http.ResponseWriter {
	return controller.httpResponseWriter
}

// Only support on http request
func (controller *Controller) GetHttpRequest() *http.Request {
	return controller.httpRequest
}

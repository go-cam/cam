package camPluginRouter

import (
	"github.com/go-cam/cam/base/camBase"
	"net/http"
)

// base controller
type Controller struct {
	camBase.ControllerInterface

	context camBase.ContextInterface
	session camBase.SessionInterface

	values        map[string]interface{} // controller values
	responseBytes []byte                 // response bytes

	httpResponseWriter http.ResponseWriter
	httpRequest        *http.Request
}

// OVERWRITE:
func (controller *Controller) Init() {
	controller.values = map[string]interface{}{}
	controller.responseBytes = []byte("")
}

// OVERWRITE
func (controller *Controller) BeforeAction(action camBase.ControllerActionInterface) bool {
	return true
}

// OVERWRITE
func (controller *Controller) AfterAction(action camBase.ControllerActionInterface, response []byte) []byte {
	return response
}

// OVERWRITE
func (controller *Controller) SetContext(context camBase.ContextInterface) {
	controller.context = context
}

// OVERWRITE
func (controller *Controller) GetContext() camBase.ContextInterface {
	return controller.context
}

// OVERWRITE
func (controller *Controller) SetSession(session camBase.SessionInterface) {
	controller.session = session
}

// OVERWRITE
func (controller *Controller) GetSession() camBase.SessionInterface {
	return controller.session
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

// set response content
func (controller *Controller) SetResponse(bytes []byte) {
	controller.responseBytes = bytes
}

// OVERWRITE
// return action write
func (controller *Controller) GetResponse() []byte {
	return controller.responseBytes
}

// OVERWRITE
func (controller *Controller) GetDefaultActionName() string {
	return ""
}

// Only support on http request
func (controller *Controller) GetHttpResponseWrite() http.ResponseWriter {
	return controller.httpResponseWriter
}

// Only support on http request
func (controller *Controller) GetHttpRequest() *http.Request {
	return controller.httpRequest
}
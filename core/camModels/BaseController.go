package camModels

import (
	"github.com/go-cam/cam/core/camBase"
	"net/http"
)

// base controller
type BaseController struct {
	camBase.ControllerInterface

	app           camBase.ApplicationInterface // app instance
	context       camBase.ContextInterface
	values        map[string]interface{}
	responseBytes []byte
}

// OVERWRITE:
func (controller *BaseController) Init() {
	controller.values = map[string]interface{}{}
	controller.responseBytes = []byte("")
}

// OVERWRITE:
func (controller *BaseController) BeforeAction(action string) bool {
	return true
}

// OVERWRITE:
func (controller *BaseController) AfterAction(action string, response []byte) []byte {
	return response
}

// OVERWRITE:
func (controller *BaseController) SetContext(context camBase.ContextInterface) {
	controller.context = context
}

// OVERWRITE:
func (controller *BaseController) GetContext() camBase.ContextInterface {
	return controller.context
}

// set http values by http.ResponseWriter and http.Request
// 	Q:	what are the values?
//	A:	values are collection of http's get and post data sent by the client
// OVERWRITE:
func (controller *BaseController) SetHttpValues(w http.ResponseWriter, r *http.Request) {
	// 接收 get 和 post 参数
	_ = r.ParseForm()
	for key, value := range r.Form {
		controller.values[key] = value
	}

	// TODO 处理数组、对象、数组和对象混合的数据
}

// set values
// OVERWRITE:
func (controller *BaseController) SetValues(values map[string]interface{}) {
	controller.values = values
}

// OVERWRITE:
func (controller *BaseController) AddValue(key string, value interface{}) {
	controller.values[key] = value
}

// get value by key
func (controller *BaseController) GetValue(key string) interface{} {
	value, has := controller.values[key]
	if !has {
		value = nil
	}
	return value
}

// set app instance
// OVERWRITE:
func (controller *BaseController) SetApp(app camBase.ApplicationInterface) {
	controller.app = app
}

// Return app instance
func (controller *BaseController) GetAppInterface() camBase.ApplicationInterface {
	return controller.app
}

// set response content
func (controller *BaseController) Write(bytes []byte) {
	controller.responseBytes = bytes
}

// return action write
// OVERWRITE:
func (controller *BaseController) Read() []byte {
	return controller.responseBytes
}

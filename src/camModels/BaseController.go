package camModels

import (
	"github.com/go-cam/cam/camBase"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// base controller
type BaseController struct {
	camBase.ControllerInterface

	app           camBase.ApplicationInterface // app instance
	context       camBase.ContextInterface
	values        map[string]interface{}
	responseBytes []byte
	// default action name. default value: ""
	// Example: "Login", "Action".
	DefaultAction string
}

// OVERWRITE:
func (controller *BaseController) Init() {
	controller.values = map[string]interface{}{}
	controller.responseBytes = []byte("")
	controller.DefaultAction = ""
}

// OVERWRITE
func (controller *BaseController) BeforeAction(action string) bool {
	return true
}

// OVERWRITE
func (controller *BaseController) AfterAction(action string, response []byte) []byte {
	return response
}

// OVERWRITE
func (controller *BaseController) SetContext(context camBase.ContextInterface) {
	controller.context = context
}

// OVERWRITE
func (controller *BaseController) GetContext() camBase.ContextInterface {
	return controller.context
}

// OVERWRITE
// set http values by http.ResponseWriter and http.Request
//	Q:	what are the values?
//	A:	values are collection of http's get and post data sent by the client
func (controller *BaseController) SetHttpValues(w http.ResponseWriter, r *http.Request) {
	controller.parseUrlValues(r)
	controller.parseFormValues(r)
}

// OVERWRITE
// set values
func (controller *BaseController) SetValues(values map[string]interface{}) {
	controller.values = values
}

// OVERWRITE
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

// OVERWRITE
// set app instance
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

// OVERWRITE
// return action write
func (controller *BaseController) Read() []byte {
	return controller.responseBytes
}

// OVERWRITE
func (controller *BaseController) GetDefaultAction() string {
	return controller.DefaultAction
}

// parse params from request url
func (controller *BaseController) parseUrlValues(r *http.Request) {
	_ = r.ParseForm()
	for key, value := range r.Form {
		controller.values[key] = value
	}
}

// parse params from form data
func (controller *BaseController) parseFormValues(r *http.Request) {
	// multipart/form-data; boundary=----WebKitFormBoundaryDumfytNg1NzoZq2r
	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "multipart/form-data") {
		boundaryRegexp, _ := regexp.Compile("boundary=([-|0-9a-zA-Z]+)")
		boundaries := boundaryRegexp.FindStringSubmatch(contentType)
		if len(boundaries) < 2 {
			panic("fail to parse form values")
		}
		boundary := "--" + boundaries[1]

		bytes, _ := ioutil.ReadAll(r.Body)
		bodyStr := string(bytes)
		paramsStrList := strings.Split(bodyStr, boundary)

		for _, row := range paramsStrList {
			if row == "" || !strings.Contains(row, "\"") {
				// exclude row
				continue
			}

			repl := "Content-Disposition: form-data; name=\"([0-9a-zA-Z|_]+)\""
			keyRegexp, _ := regexp.Compile(repl)
			keyList := keyRegexp.FindStringSubmatch(row)
			key := keyList[1]

			valueRow := keyRegexp.ReplaceAllString(row, "")
			value := strings.Trim(valueRow, "\n")
			value = strings.Trim(value, "\r")
			value = strings.Trim(value, "\r\n")
			value = strings.Trim(value, " ")

			controller.AddValue(key, value)
		}
	}

}

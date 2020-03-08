package camHttp

import (
	"fmt"
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/plugin/camPluginRouter"
	"net/http"
)

// http controller interface
type HttpControllerInterface interface {
	setResponseWriterAndRequest(w *http.ResponseWriter, r *http.Request)
}

// http controller
type HttpController struct {
	camPluginRouter.Controller
	HttpControllerInterface

	responseWriter *http.ResponseWriter
	request        *http.Request
}

// set response writer and request
func (ctrl *HttpController) setResponseWriterAndRequest(w *http.ResponseWriter, r *http.Request) {
	ctrl.responseWriter = w
	ctrl.request = r
}

func (ctrl *HttpController) GetRequestWriter() *http.ResponseWriter {
	return ctrl.responseWriter
}

func (ctrl *HttpController) GetRequest() *http.Request {
	return ctrl.request
}

// Deprecated: cannot work
func (ctrl *HttpController) GetFile(key string) {
	file, _, err := ctrl.request.FormFile(key)
	if err != nil {
		camBase.App.Error("HttpController.GetFile", err.Error())
		return
	}

	fmt.Println(file)
}

package camHttp

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/plugin/camPluginRouter"
	"net/http"
)

// http controller interface
type HttpControllerInterface interface {
	// Deprecated: Please use context to implement camBase.ContextHttpInterface
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

// Deprecated: remove on v0.5.0 . Please use context to implement camBase.ContextHttpInterface
func (ctrl *HttpController) GetRequestWriter() *http.ResponseWriter {
	return ctrl.responseWriter
}

// Deprecated: remove on v0.5.0 . Please use context to implement camBase.ContextHttpInterface
func (ctrl *HttpController) GetRequest() *http.Request {
	return ctrl.request
}

// Get upload file
// call UploadFile.Save(...) if you want to save the upload file
func (ctrl *HttpController) GetFile(key string) *UploadFile {
	file, header, err := ctrl.GetRequest().FormFile(key)
	if err != nil {
		camBase.App.Error("HttpController.GetFile", err.Error())
		return nil
	}

	return NewUploadFile(file, header)
}

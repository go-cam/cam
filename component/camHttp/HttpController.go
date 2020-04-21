package camHttp

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/plugin/camRouter"
	"net/http"
)

// http controller interface
// Deprecated: remove on v0.5.0
type HttpControllerInterface interface {
	// Deprecated: remove on v0.5.0
	// Instead: Please use context to implement camBase.ContextHttpInterface
	setResponseWriterAndRequest(w *http.ResponseWriter, r *http.Request)
}

// http controller
type HttpController struct {
	camRouter.Controller

	// Deprecated: remove on v0.5.0
	responseWriter *http.ResponseWriter
	// Deprecated: remove on v0.5.0
	request *http.Request
}

// set response writer and request
// Deprecated: remove on v0.5.0
func (ctrl *HttpController) setResponseWriterAndRequest(w *http.ResponseWriter, r *http.Request) {
	ctrl.responseWriter = w
	ctrl.request = r
}

// Deprecated: remove on v0.5.0
// Instead: Please use context to implement camBase.ContextHttpInterface
func (ctrl *HttpController) GetRequestWriter() *http.ResponseWriter {
	return ctrl.responseWriter
}

// Deprecated: remove on v0.5.0
// Instead: Please use context to implement camBase.ContextHttpInterface
func (ctrl *HttpController) GetRequest() *http.Request {
	return ctrl.request
}

// Get HttpContextInterface
func (ctrl *HttpController) GetHttpContext() camBase.HttpContextInterface {
	ctxI := ctrl.GetContext()
	httpCtxI, ok := ctxI.(camBase.HttpContextInterface)
	if !ok {
		return nil
	}
	return httpCtxI
}

// Get upload file
// call UploadFile.Save(...) if you want to save the upload file
func (ctrl *HttpController) GetFile(key string) *UploadFile {
	ctxI := ctrl.GetHttpContext()
	if ctxI == nil {
		return nil
	}

	file, header, err := ctxI.GetHttpRequest().FormFile(key)
	if err != nil {
		camBase.App.Error("HttpController.GetFile", err.Error())
		return nil
	}

	return NewUploadFile(file, header)
}

package camHttp

import (
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/plugin/camRouter"
)

// http controller
type HttpController struct {
	camRouter.Controller
}

// Get HttpContextInterface
func (ctrl *HttpController) GetHttpContext() camStatics.HttpContextInterface {
	ctxI := ctrl.GetContext()
	httpCtxI, ok := ctxI.(camStatics.HttpContextInterface)
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
		camStatics.App.Error("HttpController.GetFile", err.Error())
		return nil
	}

	return NewUploadFile(file, header)
}

package camConsole

import (
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/base/camUtils"
)

type GrpcOption struct {
	// the module's dirs
	//
	// Example: []string{"backend-grpc", "center-grpc"}
	GrpcDirs []string
}

type GrpcController struct {
	ConsoleController
}

func (ctrl *GrpcController) DownloadProtocGenGo() {
	err := camUtils.Console.Start("go get github.com/golang/protobuf/protoc-gen-go@v1.3.5")
	if err != nil {
		camStatics.App.Error("GrpcController.DownloadProtocGenGo", "err: "+err.Error())
	}
}

//// Compile proto To go's files
//func (ctrl *GrpcController) Compile() {
//	option := ctrl.GetConsoleComponent().config.grpcOption
//
//}

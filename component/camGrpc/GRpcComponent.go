package camGrpc

import (
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/component"
	"google.golang.org/grpc"
)

type GRpcComponent struct {
	component.Component

	conf *GRpcComponentConfig
}

func (comp *GRpcComponent) Init(confI camStatics.ComponentConfigInterface) {
	comp.Component.Init(confI)

	var ok bool
	comp.conf, ok = confI.(*GRpcComponentConfig)
	if !ok {
		camStatics.App.Fatal("GRpcComponent", "invalid config")
		return
	}
}

func (comp *GRpcComponent) Start() {
	comp.Component.Start()
}

func (comp *GRpcComponent) Stop() {
	defer comp.Component.Stop()
}

func (comp *GRpcComponent) Conn() *grpc.ClientConn {
	conn, err := grpc.Dial(comp.serverAddr())
	if err != nil {
		panic(err)
	}
	return conn
}

// TODO
func (comp *GRpcComponent) serverAddr() string {
	return comp.conf.opt.Client.ServerList[0]
}

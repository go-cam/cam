package camGRpcServer

import (
	"crypto/tls"
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/base/camUtils"
	"github.com/go-cam/cam/component"
	"github.com/go-cam/cam/plugin/camTls"
	"google.golang.org/grpc"
	"net"
	"reflect"
)

type GRpcServerComponent struct {
	component.Component

	conf *GRpcServerComponentConfig
	camTls.TlsPlugin
}

// init conf
func (comp *GRpcServerComponent) Init(confI camStatics.ComponentConfigInterface) {
	comp.Component.Init(confI)

	conf, ok := confI.(*GRpcServerComponentConfig)
	if !ok {
		camStatics.App.Fatal("GRpcServerComponentConfig", "invalid conf")
		return
	}
	comp.conf = conf
	comp.TlsPlugin.Init(&comp.conf.TlsPluginConfig)
	comp.TlsPlugin.SetListenHandler(comp.listen, comp.listenTls)
}

// start
func (comp *GRpcServerComponent) Start() {
	comp.Component.Start()
	go comp.StartListenServer()
}

// stop
func (comp *GRpcServerComponent) Stop() {
	defer comp.Component.Stop()
}

func (comp *GRpcServerComponent) listen() {
	addr := ":" + camUtils.C.Uint16ToString(comp.conf.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		camStatics.App.Fatal("GRpcServerComponent", "Listen failed. Err:" + err.Error())
		return
	}
	server := comp.getService()
	camStatics.App.Trace("GRpcServerComponent", "listen: tcp://" + addr)
	err = server.Serve(lis)
	if err != nil {
		camStatics.App.Fatal("GRpcServerComponent", "Listen failed. Err:" + err.Error())
	}
}

func (comp *GRpcServerComponent) listenTls() {
	cert, err := tls.LoadX509KeyPair(comp.conf.TlsCertFile, comp.conf.TlsKeyFile)
	if err != nil {
		camStatics.App.Fatal("GRpcServerComponent", "Listen failed. Err:" + err.Error())
		return
	}

	tlsConf := &tls.Config{Certificates: []tls.Certificate{cert}}
	addr := ":" + camUtils.C.Uint16ToString(comp.conf.TlsPort)
	lis, err := tls.Listen("tcp", addr, tlsConf)
	if err != nil {
		camStatics.App.Fatal("GRpcServerComponent", "Listen failed. Err:" + err.Error())
		return
	}
	server := comp.getService()
	camStatics.App.Trace("GRpcServerComponent", "listen tls: tcp://" + addr)
	err = server.Serve(lis)
	if err != nil {
		camStatics.App.Fatal("GRpcServerComponent", "Listen failed. Err:" + err.Error())
	}
}

func (comp *GRpcServerComponent) getService() *grpc.Server {
	server := grpc.NewServer(comp.conf.DialOptions...)
	for _, srvOpt := range comp.conf.services {
		regHandlerV := reflect.ValueOf(srvOpt.regHandler)
		v1 := reflect.ValueOf(server)
		v2 := reflect.ValueOf(srvOpt.srv)
		_ = regHandlerV.Call([]reflect.Value{v1, v2})
	}
	return server
}

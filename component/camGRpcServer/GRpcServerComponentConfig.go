package camGRpcServer

import (
	"github.com/go-cam/cam/component"
	"github.com/go-cam/cam/plugin/camTls"
	"google.golang.org/grpc"
)

type ServiceOption struct {
	srv        interface{}
	regHandler interface{}
}

type Option struct {
	Port        uint16
	DialOptions []grpc.ServerOption
	services    []ServiceOption
}

type GRpcServerComponentConfig struct {
	component.ComponentConfig
	camTls.TlsPluginConfig

	Option
}

func NewGRpcServer() *GRpcServerComponentConfig {
	conf := new(GRpcServerComponentConfig)
	conf.Component = &GRpcServerComponent{}
	conf.Option = Option{}
	conf.services = []ServiceOption{}
	return conf
}

func (conf *GRpcServerComponentConfig) Register(service interface{}, regHandler interface{}) {
	conf.services = append(conf.services, ServiceOption{srv: service, regHandler: regHandler})
}

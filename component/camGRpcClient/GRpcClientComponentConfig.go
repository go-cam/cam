package camGRpcClient

import (
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/component"
	"google.golang.org/grpc"
)

// server struct
type Server struct {
	// addr
	// Example: "localhost:50051"
	Addr        string
	// dial options
	// Example: grpc.WithInsecure(), grpc.WithBlock()
	DialOptions []grpc.DialOption
}

// client options
type Option struct {
	LoadBalancingLogic camStatics.GRpcLoadBalancingLogic
	// Server config
	Servers []*Server
}

// gRpc client component's conf
type GRpcClientComponentConfig struct {
	component.ComponentConfig
	Option
}

// new gRpc client conf
func NewGRpcClient() *GRpcClientComponentConfig {
	conf := new(GRpcClientComponentConfig)
	conf.Component = &GRpcClientComponent{}
	conf.Option = Option{}
	return conf
}

// set Option
func (conf *GRpcClientComponentConfig) SetOption(option *Option) {
	conf.Option = *option
	if conf.LoadBalancingLogic == 0 {
		conf.LoadBalancingLogic = camStatics.GRpcLoadBalancingLogicSequence
	}
}

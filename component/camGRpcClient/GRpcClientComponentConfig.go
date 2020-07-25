package camGRpcClient

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camConstants"
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
type GRpcClientOption struct {
	LoadBalancingLogic camBase.GRpcLoadBalancingLogic
	// Server config
	Servers []*Server
}

// gRpc client component's conf
type GRpcClientComponentConfig struct {
	component.ComponentConfig
	option *GRpcClientOption
}

// new gRpc client conf
func NewGRpcClient() *GRpcClientComponentConfig {
	conf := new(GRpcClientComponentConfig)
	conf.Component = &GRpcClientComponent{}
	return conf
}

// set option
func (conf *GRpcClientComponentConfig) SetOption(option *GRpcClientOption) {
	conf.option = option
	if conf.option.LoadBalancingLogic == 0 {
		conf.option.LoadBalancingLogic = camConstants.GRpcLoadBalancingLogicSequence
	}
}

package camGrpc

import (
	"github.com/go-cam/cam/component"
	"google.golang.org/grpc"
)

// What do you want the component to be the client or server
type Type string

const (
	TypeClient Type = "CLIENT"
	TypeServer Type = "SERVER"
)

type ClientOption struct {
	// Server addresses
	ServerList []string
}

// server option
type ServerOption struct {
	// Server start address
	Addr string
	// Server grpc's options
	Opts []grpc.ServerOption
	// Friends list. Distributed logic can be implemented
	Friends []string
}

// config options
type Option struct {
	// grpc type.
	// What do you want the component to be the client or server
	Type   Type
	Client ClientOption
	Server ServerOption
}

// Deprecated
type GRpcComponentConfig struct {
	component.ComponentConfig

	opt     *Option
	srvList []interface{}
}

// Deprecated: instead by camGRpcClient and camGRpcServer
// new GrpcComponent config
func NewGRpcConfig(opt *Option) *GRpcComponentConfig {
	conf := new(GRpcComponentConfig)
	if opt == nil {
		panic("opt cannot be nil")
	}
	conf.opt = opt
	conf.srvList = []interface{}{}
	return conf
}

// Register server
func (conf *GRpcComponentConfig) RegisterServer(srv interface{}) {
	conf.srvList = append(conf.srvList, srv)
}

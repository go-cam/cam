package camGRpcClient

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camConstants"
	"github.com/go-cam/cam/component"
	"google.golang.org/grpc"
)

type sequenceOption struct {
	keys  []string
	len   int
	index int
}

func newSequenceOption(connDict map[string]*grpc.ClientConn) *sequenceOption {
	opt := new(sequenceOption)
	for key, _ := range connDict {
		opt.keys = append(opt.keys, key)
	}
	opt.len = len(connDict)
	opt.index = 0
	return opt
}

// get now key and move index to the next
func (opt *sequenceOption) nextKey() string {
	key := opt.keys[opt.index]
	if opt.len > 1 {
		opt.index++
		if opt.index == opt.len {
			opt.index = 0
		}
	}
	return key
}

type GRpcClientComponent struct {
	component.Component

	conf     *GRpcClientComponentConfig
	connDict map[string]*grpc.ClientConn
	// only sequence logic will be create
	seqOpt   *sequenceOption
}

// init conf
func (comp *GRpcClientComponent) Init(confI camBase.ComponentConfigInterface) {
	comp.Component.Init(confI)

	conf, ok := confI.(*GRpcClientComponentConfig)
	if !ok {
		camBase.App.Fatal("GRpcClientComponentConfig", "invalid conf")
		return
	}
	comp.conf = conf
	comp.connDict = map[string]*grpc.ClientConn{}
}

// start
func (comp *GRpcClientComponent) Start() {
	comp.Component.Start()

	// Create connections
	if comp.conf.option.Servers != nil {
		for _, server := range comp.conf.option.Servers {
			conn, err := grpc.Dial(server.Addr, server.DialOptions...)
			if err != nil {
				camBase.App.Error("GRpcClientComponent.Start()", err.Error())
				continue
			}
			comp.connDict[server.Addr] = conn
		}
	}
	if len(comp.connDict) == 0 {
		panic("There's no connection")
	}
	comp.initOptionByLoadBalancingLogin()
}

// stop
func (comp *GRpcClientComponent) Stop() {
	defer comp.Component.Stop()
}

// init option by load balancing login
func (comp *GRpcClientComponent) initOptionByLoadBalancingLogin() {
	switch comp.conf.option.LoadBalancingLogic {
	case camConstants.GRpcLoadBalancingLogicSequence:
		comp.seqOpt = newSequenceOption(comp.connDict)
	default:
		panic("There's no gRpc Load Balancing config")
	}
}

// get connect, use this conn to create GRpc Service
func (comp *GRpcClientComponent) GetConn() *grpc.ClientConn {
	var conn *grpc.ClientConn

	switch comp.conf.option.LoadBalancingLogic {
	case camConstants.GRpcLoadBalancingLogicSequence:
		conn = comp.getConnBySequence()
	}

	return conn
}

func (comp *GRpcClientComponent) getConnBySequence() *grpc.ClientConn {
	key := comp.seqOpt.nextKey()
	return comp.connDict[key]
}

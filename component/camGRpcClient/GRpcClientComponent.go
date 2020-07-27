package camGRpcClient

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camConstants"
	"github.com/go-cam/cam/component"
	"google.golang.org/grpc"
	"sync"
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
	seqOpt *sequenceOption
	// lock func reconnect()
	checkAndReconnectMutex sync.Mutex
	// If it true, when call GetConn() doesn't need to wait for a connection to be detected
	canGoOn                bool
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
	comp.checkAndReconnectMutex = sync.Mutex{}
	comp.canGoOn = false
}

// start
func (comp *GRpcClientComponent) Start() {
	comp.Component.Start()

	if len(comp.conf.Servers) == 0 {
		camBase.App.Fatal("GRpcClientComponent", "There are not connections.")
		return
	}

	// Create connections
	if comp.conf.Servers != nil {
		for _, server := range comp.conf.Servers {
			comp.createConn(server)
		}
	}
	comp.initOptionByLoadBalancingLogin()
}

// stop
func (comp *GRpcClientComponent) Stop() {
	defer comp.Component.Stop()
}

// init Option by load balancing login
func (comp *GRpcClientComponent) initOptionByLoadBalancingLogin() {
	switch comp.conf.LoadBalancingLogic {
	case camConstants.GRpcLoadBalancingLogicSequence:
		comp.seqOpt = newSequenceOption(comp.connDict)
	default:
		panic("There's no gRpc Load Balancing config")
	}
}

// get connect, use this conn to create GRpc Service
func (comp *GRpcClientComponent) GetConn() *grpc.ClientConn {
	var conn *grpc.ClientConn

	var goOn = make(chan bool)
	go comp.reconnect(goOn)
	if <- goOn {
	}

	switch comp.conf.LoadBalancingLogic {
	case camConstants.GRpcLoadBalancingLogicSequence:
		conn = comp.getConnBySequence()
	}

	// TODO check conn status, if cannot used, get new one

	return conn
}

// Create to gRpc server connection.
// Return:
//	true: 	conn success.
//	false:	conn failed.
func (comp *GRpcClientComponent) createConn(server *Server) bool {
	conn, err := grpc.Dial(server.Addr, server.DialOptions...)
	if err != nil {
		camBase.App.Error("GRpcClientComponent", "Connective failed: "+err.Error())
		comp.connDict[server.Addr] = nil
		return false
	}

	camBase.App.Trace("GRpcClientComponent", "Connect to addr: " + server.Addr)
	comp.connDict[server.Addr] = conn
	return true
}

func (comp *GRpcClientComponent) reconnect(goOn chan bool) {
	comp.checkAndReconnectMutex.Lock()
	defer func() {
		comp.checkAndReconnectMutex.Unlock()
		goOn <- true
	}()

	if comp.canGoOn {
		goOn <- true
		return
	}

	for _, server := range comp.conf.Servers {
		conn, has := comp.connDict[server.Addr]
		if has && conn != nil {
			goOn <- true
			continue
		}
		done := comp.createConn(server)
		if done {
			goOn <- true
		}
	}
}

func (comp *GRpcClientComponent) getConnBySequence() *grpc.ClientConn {
	key := comp.seqOpt.nextKey()
	conn := comp.connDict[key]
	if conn == nil {
		return comp.getConnBySequence()
	}
	return conn
}

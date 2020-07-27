package camGRpcClient

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camConstants"
	"github.com/go-cam/cam/component"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"sync"
)




type GRpcClientComponent struct {
	component.Component

	conf     *GRpcClientComponentConfig
	// connect dict
	connDict *connDict
	// only sequence logic will be create
	seqOpt *sequenceOption
	// lock func reconnect()
	reconnectMutex sync.Mutex
	// If it true, when call GetConn() doesn't need to wait for a connection to be detected
	canGoOn                bool
	// serverOption key => index to the conf.Servers
	serverAddrIndexDict map[string]int
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
	comp.connDict = &connDict{sync.Map{}}
	comp.reconnectMutex = sync.Mutex{}
	comp.canGoOn = false
	comp.serverAddrIndexDict = map[string]int{}
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
		for i, server := range comp.conf.Servers {
			comp.createConn(server)
			comp.serverAddrIndexDict[server.Addr] = i
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

	switch comp.conf.LoadBalancingLogic {
	case camConstants.GRpcLoadBalancingLogicSequence:
		conn = comp.getConnBySequence()
	}

	if !comp.checkConn(conn) {
		var goOn = make(chan bool)
		go comp.reconnect(goOn)
		<- goOn
		close(goOn)
		conn = comp.GetConn()
	}

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
		comp.connDict.Set(server.Addr, nil)
		return false
	}

	if !comp.checkConn(conn) {
		camBase.App.Error("GRpcClientComponent", "Connective failed")
		return false
	}
	comp.connDict.Set(server.Addr, conn)
	camBase.App.Trace("GRpcClientComponent", "Connect to addr: " + server.Addr)
	return true
}

// reconnect
func (comp *GRpcClientComponent) reconnect(goOn chan bool) {
	comp.reconnectMutex.Lock()
	defer func() {
		comp.reconnectMutex.Unlock()
		goOn <- true
	}()

	for _, server := range comp.conf.Servers {
		conn, has := comp.connDict.Get(server.Addr)
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

// If the connection was broken, it will be reconnected
func (comp *GRpcClientComponent) fixConn() {
	comp.connDict.Range(func(addr string, conn *grpc.ClientConn) bool {
		if !comp.checkConn(conn) {
			server := comp.getServerOptByAddr(addr)
			comp.createConn(server)
		}
		return true
	})
}

// Check if the connection is available
func (comp *GRpcClientComponent) checkConn(conn *grpc.ClientConn) bool {
	if conn == nil {
		return false
	}
	if conn.GetState() != connectivity.Idle {
		return false
	}
	return true
}

// sequence logic
func (comp *GRpcClientComponent) getConnBySequence() *grpc.ClientConn {
	key := comp.seqOpt.nextKey()
	conn, has := comp.connDict.Get(key)
	if !has || conn == nil {
		return comp.getConnBySequence()
	}
	return conn
}

func (comp *GRpcClientComponent) getServerOptByAddr(addr string) *Server {
	i, has := comp.serverAddrIndexDict[addr]
	if !has {
		camBase.App.Fatal("GRpcClientComponent.getServerOptByAddr()", "addr has no config. addr: " + addr)
		return nil
	}
	return comp.conf.Servers[i]
}

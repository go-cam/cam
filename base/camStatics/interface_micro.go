package camStatics

import "google.golang.org/grpc"

// IMicroServerComponent
// register center
type IMicroServerComponent interface {
	Register(client IMicroApp)
	GetAddress(name string) (string, error)
	RefreshHeartbeat(microApp IMicroApp)
}

type IMicroServerComponentConfig interface {
	Port() uint16
}

// IMicroClientComponent
// discovery client
type IMicroClientComponent interface {
	GetGrpcConn(name string) (*grpc.ClientConn, error)
}

// IMicroApp
// client info
type IMicroApp interface {
	AppName() string
	Address() string
}

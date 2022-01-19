package camStatics

type MicroProtocol uint8

const (
	MicroProtocolGrpc MicroProtocol = iota
)

type MicroType uint8

const (
	MicroTypeClient MicroType = iota
	MicroTypeServer
)

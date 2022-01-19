package camStructs

import "github.com/go-cam/cam/base/camStatics"

type MicroApp struct {
	camStatics.IMicroApp
	appName      string
	address      string
	microType    camStatics.MicroType
	protocolList []camStatics.MicroProtocol // support protocol
}

func (m *MicroApp) AppName() string {
	return m.appName
}

func (m *MicroApp) Address() string {
	return m.address
}

func NewMicroApp(appName string, address string) *MicroApp {
	client := new(MicroApp)
	client.microType = camStatics.MicroTypeClient
	client.appName = appName
	client.address = address
	return client
}

func NewMicroAppServer(appName string, address string, protocolList ...camStatics.MicroProtocol) *MicroApp {
	if len(protocolList) == 0 {
		panic("Please select at least one protocol")
	}
	client := new(MicroApp)
	client.microType = camStatics.MicroTypeServer
	client.appName = appName
	client.address = address
	client.protocolList = protocolList
	return client
}

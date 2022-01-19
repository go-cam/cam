package camMicroServer

import (
	"errors"
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/component"
	"github.com/go-cam/cam/proto/github_io_cam_micro"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"net"
	"strconv"
	"time"
)

type MicroServerComponent struct {
	camStatics.IMicroServerComponent
	component.Component

	conf             camStatics.IMicroServerComponentConfig
	microAppDictDict *microAppDictDict
	upgrader         websocket.Upgrader
}

// Init init conf
func (comp *MicroServerComponent) Init(confI camStatics.IComponentConfig) {
	comp.Component.Init(confI)

	conf, ok := confI.(camStatics.IMicroServerComponentConfig)
	if !ok {
		camStatics.App.Fatal("MicroServerComponent", "invalid conf")
		return
	}
	comp.conf = conf
	comp.microAppDictDict = new(microAppDictDict)
}

func (comp *MicroServerComponent) Start() {
	comp.Component.Start()

	// start grpc microAppInfo
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(int(comp.conf.Port())))
	if err != nil {
		panic(err)
	}
	err = comp.getService().Serve(lis)
	if err != nil {
		panic(err)
	}
}

func (comp *MicroServerComponent) Stop() {
	defer comp.Component.Stop()
}

// Register register micro client
func (comp *MicroServerComponent) Register(client camStatics.IMicroApp) {
	info := new(microAppInfo)
	info.name = client.AppName()
	info.address = client.Address()
	comp.microAppDictDict.put(info)
}

// GetAddress get micro microAppInfo address
func (comp *MicroServerComponent) GetAddress(name string) (string, error) {
	var microApp *microAppInfo = nil
	comp.microAppDictDict.getMicroAppInfoDict(name).Range(func(address, tmpMicroApp interface{}) bool {
		microApp = tmpMicroApp.(*microAppInfo)
		return false
	})
	if microApp == nil {
		return "", errors.New("no microApp[" + name + "] found")
	}
	return microApp.address, nil
}

func (comp *MicroServerComponent) RefreshHeartbeat(microApp camStatics.IMicroApp) {
	info := comp.microAppDictDict.getMicroAppInfo(microApp.AppName(), microApp.Address())
	if info == nil {
		comp.microAppDictDict.put(info)
		comp.RefreshHeartbeat(microApp)
	}
	info.lastHeartbeatMS = time.Now().UnixMilli()
	camStatics.App.Trace("MicroServerComponent:RefreshHeartbeat", microApp.AppName()+" "+microApp.Address())
}

func (comp *MicroServerComponent) getService() *grpc.Server {
	server := grpc.NewServer()
	github_io_cam_micro.RegisterCamMicroServer(server, &CamMicroServer{})
	return server
}

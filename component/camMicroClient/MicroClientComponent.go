package camMicroClient

import (
	"context"
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/component"
	"github.com/go-cam/cam/proto/github_io_cam_micro"
	"google.golang.org/grpc"
	"time"
)

var heartbeatDuration = time.Second * 10

type MicroClientComponent struct {
	component.Component
	conf        *MicroClientComponentConfig
	lastAliveMS int64
}

func (comp *MicroClientComponent) Init(iConfig camStatics.IComponentConfig) {
	comp.Component.Init(iConfig)
	comp.lastAliveMS = 0
}

func (comp *MicroClientComponent) Start() {
	comp.Component.Start()
	go comp.connectServer()
}

func (comp *MicroClientComponent) Stop() {
	comp.Component.Stop()
}

func (comp *MicroClientComponent) GetGrpcConn(name string) (*grpc.ClientConn, error) {
	client := github_io_cam_micro.NewCamMicroClient(comp.getMicroServerConn())
	out, err := client.GetServer(context.Background(), &github_io_cam_micro.CamMicroGetServerIn{AppName: name})
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(out.GetAddress())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (comp *MicroClientComponent) refreshAlive() {
	comp.lastAliveMS = time.Now().UnixMilli()
}

func (comp *MicroClientComponent) isNeedPing() bool {
	return comp.lastAliveMS+int64(heartbeatDuration) < time.Now().UnixMilli()
}

func (comp *MicroClientComponent) getMicroServerConn() *grpc.ClientConn {
	conn, err := grpc.Dial(comp.conf.ServerAddress)
	if err == nil {
		panic(err)
	}
	return conn
}

func (comp *MicroClientComponent) connectServer() {
	defer func() {
		err := recover()
		if err != nil {
			time.Sleep(heartbeatDuration)
			go comp.connectServer() // 重新连接
		}
	}()

	client := github_io_cam_micro.NewCamMicroClient(comp.getMicroServerConn())
	stream, err := client.Heartbeat(context.Background())
	if err != nil {
		panic(err)
	}

	// 接收数据
	go func() {
		out, err := stream.Recv()
		for err != nil {
			if out.Type == github_io_cam_micro.CamMicroHeartbeatType_Ping {
				err = stream.Send(&github_io_cam_micro.CamMicroHeartbeatSteam{
					Type: github_io_cam_micro.CamMicroHeartbeatType_Pong,
				})
				if err != nil {
					panic(err)
				}
				comp.refreshAlive()
			} else if out.Type == github_io_cam_micro.CamMicroHeartbeatType_Pong {
				comp.refreshAlive()
			}
			out, err = stream.Recv()
		}
		panic(err)
	}()

	// 主动发送心跳
	go func() {
		var err error
		for {
			if comp.isNeedPing() {
				err = stream.Send(&github_io_cam_micro.CamMicroHeartbeatSteam{
					Type: github_io_cam_micro.CamMicroHeartbeatType_Pong,
				})
				if err != nil {
					panic(err)
				}
			}
			time.Sleep(heartbeatDuration)
		}
	}()
}

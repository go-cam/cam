package camMicroServer

import (
	"context"
	"errors"
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/base/camStructs"
	"github.com/go-cam/cam/proto/github_io_cam_micro"
	"time"
)

var microServer camStatics.IMicroServerComponent = nil

type CamMicroServer struct {
	github_io_cam_micro.CamMicroServer
	microApp camStatics.IMicroApp
}

func (s *CamMicroServer) Register(ctx context.Context, in *github_io_cam_micro.CamMicroRegisterIn) (*github_io_cam_micro.CamMicroRegisterOut, error) {
	s.microApp = camStructs.NewMicroApp(in.GetAppName(), in.GetAddress())
	s.getMicroServerComponent().Register(s.microApp)

	out := &github_io_cam_micro.CamMicroRegisterOut{
		Done: true,
	}
	return out, nil
}

func (s *CamMicroServer) Heartbeat(stream github_io_cam_micro.CamMicro_HeartbeatServer) error {
	if s.microApp == nil {
		return errors.New("please register first")
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
				s.getMicroServerComponent().RefreshHeartbeat(s.microApp)
			} else if out.Type == github_io_cam_micro.CamMicroHeartbeatType_Pong {
				s.getMicroServerComponent().RefreshHeartbeat(s.microApp)
			}
			out, err = stream.Recv()
		}
		panic(err)
	}()

	var err error
	for {
		if true {
			err = stream.Send(&github_io_cam_micro.CamMicroHeartbeatSteam{
				Type: github_io_cam_micro.CamMicroHeartbeatType_Ping,
			})
			if err != nil {
				panic(err)
			}
		}
		time.Sleep(time.Second * 10)
	}
	return nil
}

func (s *CamMicroServer) GetServer(ctx context.Context, in *github_io_cam_micro.CamMicroGetServerIn) (*github_io_cam_micro.CamMicroGetServerOut, error) {
	address, err := s.getMicroServerComponent().GetAddress(in.GetAppName())
	if err != nil {
		return nil, errors.New("no [appName: " + in.GetAppName() + "] found")
	}
	return &github_io_cam_micro.CamMicroGetServerOut{
		Address: address,
	}, nil
}

func (s *CamMicroServer) getMicroServerComponent() camStatics.IMicroServerComponent {
	if microServer == nil {
		iComp := camStatics.App.GetComponent(&MicroServerComponent{})
		if iComp == nil {
			panic("no micro microAppInfo component")
		}
		iMicroServerComp, ok := iComp.(camStatics.IMicroServerComponent)
		if !ok {
			panic("no micro microAppInfo component")
		}
		microServer = iMicroServerComp
	}
	return microServer
}

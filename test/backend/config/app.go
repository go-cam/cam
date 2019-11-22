package config

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/core/base"
)

// 获取默认配置
func GetApp() *cam.Config {
	config := cam.NewConfig()
	config.AppConfig = cam.NewAppConfig()
	config.ComponentDict = map[string]base.ConfigComponentInterface{
		"ws":      cam.NewConfigWebsocketServer(24600),
		"http":    cam.NewConfigHttpServer(24601).SetSessionName("test"),
		"db":      cam.NewConfigDatabase("mysql", "127.0.0.1", "3306", "cin", "root", "root"),
		"console": cam.NewConfigConsole(),
	}
	return config
}

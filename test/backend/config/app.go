package config

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/core/camBase"
)

// 获取默认配置
func GetApp() *cam.Config {
	config := cam.NewConfig()
	config.AppConfig = cam.NewAppConfig()
	config.ComponentDict = map[string]camBase.ConfigComponentInterface{
		"ws":      cam.NewWebsocketServerConfig(24600),
		"http":    cam.NewHttpServerConfig(24601).SetSessionName("test"),
		"db":      cam.NewDatabaseConfig("mysql", "127.0.0.1", "3306", "cin", "root", "root"),
		"console": cam.NewConsoleConfig(),
	}
	return config
}

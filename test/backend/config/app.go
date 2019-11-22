package config

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/core/base"
)

// 获取默认配置
func GetApp() *cin.Config {
	config := cin.NewConfig()
	config.AppConfig = cin.NewAppConfig()
	config.ComponentDict = map[string]base.ConfigComponentInterface{
		"ws":      cin.NewConfigWebsocketServer(24600),
		"http":    cin.NewConfigHttpServer(24601).SetSessionName("test"),
		"db":      cin.NewConfigDatabase("mysql", "127.0.0.1", "3306", "cin", "root", "root"),
		"console": cin.NewConfigConsole(),
	}
	return config
}

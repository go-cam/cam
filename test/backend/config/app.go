package config

import (
	"cin"
	"cin/base"
)

// 获取默认配置
func GetApp() *cin.Config {
	config := cin.NewConfig()
	config.ComponentDict = map[string]base.ConfigComponentInterface{
		"ws":   cin.NewConfigWebsocketServer(24600),
		"http": cin.NewConfigHttpServer(24601).SetSessionName("test"),
		"db": cin.NewConfigDatabase("mysql", "127.0.0.1", "3306", "cin", "root", "root"),
		"console": cin.NewConfigConsole(),
	}
	return config
}
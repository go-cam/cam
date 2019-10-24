package config

import (
	"cin"
	"cin/base"
	"cin/components"
	"cin/configs"
)

// 获取默认配置
func GetApp() *cin.Config {
	config := cin.NewConfig()
	config.ComponentDict = map[string]base.ConfigComponentInterface{
		"ws":   configs.NewWebsocketServer(&components.WebsocketServer{}, 10001),
		"http": configs.NewHttpServer(&components.HttpServer{}, 10000).SetSessionName("test"),
	}
	return config
}
package main

import (
	cin "cin/src"
	"cin/src/configs"
)

func main() {
	cin.App.AddConfig(config())
	cin.App.Run()
}

// 配置
func config() *cin.Config {
	config := new(cin.Config)
	config.SetParams(map[string]interface{}{
		"test": 123123,
	})
	config.AddComponents("ws", &configs.WebsocketServer{
		Port: 10001,
		Mode: cin.WebsocketServerModeAutoHandler,
	})
	return config
}
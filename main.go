package main

import (
	cin "cin/src"
	"cin/src/base"
	"cin/src/components"
	"cin/src/configs"
)

func main() {
	cin.App.AddConfig(config())
	cin.App.Run()
}

// 配置
func config() *cin.Config {
	config := new(cin.Config)
	config.Params = map[string]interface{}{
		"test": 123123,
	}
	config.ComponentDict = map[string]base.ConfigComponentInterface{
		"ws": configs.NewWebsocketServer(&components.WebsocketServer{}, 10001),
	}
	return config
}
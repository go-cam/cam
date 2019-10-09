package main

import (
	"cin"
	"cin/base"
	"cin/components"
	"cin/configs"
	"cin/test/controllers"
)

func main() {
	cin.App.AddConfig(config())
	router := cin.App.GetRouter()
	router.Register(&controllers.TestController{})
	cin.App.Run()
}

// 配置
func config() *cin.Config {
	config := cin.NewConfig()
	config.Params = map[string]interface{}{
		"test": 123123,
	}
	config.ComponentDict = map[string]base.ConfigComponentInterface{
		"ws":   configs.NewWebsocketServer(&components.WebsocketServer{}, 10001),
		"http": configs.NewHttpServer(&components.HttpServer{}, 10000).SetSessionName("test"),
	}
	return config
}

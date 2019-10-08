package main

import (
	"fmt"
	"cin"
	"cin/base"
	"cin/components"
	"cin/configs"
	"cin/models"
	"cin/test/controllers"
)

func main() {
	cin.App.AddConfig(config())
	router := cin.App.GetRouter()
	router.Register(&controllers.TestController{})
	router.OnWebsocketMessage(func(conn *models.WebsocketSession, recvMessage []byte) {
		fmt.Println("接收：" + string(recvMessage))
		conn.Send([]byte("收到了"))
	})
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

package main

import (
	cin "../../src"
	"base"
	"components"
	"configs"
	"fmt"
	"models"
	"test/controllers"
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
		"ws": configs.NewWebsocketServer(&components.WebsocketServer{}, 10001),
	}
	return config
}
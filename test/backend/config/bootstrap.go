package config

import (
	"cin"
	"cin/test/backend/controllers"
	"cin/test/common/config"
)

// 加载配置
func LoadConfig() {
	// 加载 common 配置
	cin.App.AddConfig(config.GetApp())
	cin.App.AddConfig(config.GetAppLocal())
	// 加载 本模块配置
	cin.App.AddConfig(GetApp())
	cin.App.AddConfig(GetAppLocal())

	routeConfig()
}

// 路由设置
func routeConfig() {
	router := cin.App.GetRouter()
	router.Register(&controllers.TestController{})
}
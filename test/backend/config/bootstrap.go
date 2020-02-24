package config

import (
	"github.com/go-cam/cam"
	_ "github.com/go-sql-driver/mysql"
	"test/backend/controllers"
	_ "test/backend/database/migrations"
	"test/common/config"
)

// 加载配置
func LoadConfig() {
	// load common's config
	cam.App.AddConfig(config.GetApp())
	cam.App.AddConfig(config.GetAppLocal())
	// load module's config
	cam.App.AddConfig(GetApp())

	routeConfig()
}

// 路由设置
func routeConfig() {
	router := cam.App.GetRouter()
	router.Register(new(controllers.TestController))
}

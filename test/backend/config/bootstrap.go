package config

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/test/backend/controllers"
	_ "github.com/go-cam/cam/test/backend/database/migrations"
	"github.com/go-cam/cam/test/common/config"
	_ "github.com/go-sql-driver/mysql"
)

// 加载配置
func LoadConfig() {
	// load common's config
	cam.App.AddConfig(config.GetApp())
	cam.App.AddConfig(config.GetAppLocal())
	// load module's config
	cam.App.AddConfig(GetApp())
	cam.App.AddConfig(GetAppLocal())

	routeConfig()
}

// 路由设置
func routeConfig() {
	router := cam.App.GetRouter()
	router.Register(new(controllers.TestController))
}

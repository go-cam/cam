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
	core.App.AddConfig(config.GetApp())
	core.App.AddConfig(config.GetAppLocal())
	// load module's config
	core.App.AddConfig(GetApp())

	routeConfig()
}

// 路由设置
func routeConfig() {
	router := core.App.GetRouter()
	router.Register(new(controllers.TestController))
}

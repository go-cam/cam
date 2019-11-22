package config

import (
	"github.com/cinling/cam"
	"github.com/cinling/cam/test/backend/controllers"
	_ "github.com/cinling/cam/test/backend/database/migrations"
	"github.com/cinling/cam/test/common/config"
	_ "github.com/go-sql-driver/mysql"
)

// 加载配置
func LoadConfig() {
	// load common's config
	cin.App.AddConfig(config.GetApp())
	cin.App.AddConfig(config.GetAppLocal())
	// load module's config
	cin.App.AddConfig(GetApp())
	cin.App.AddConfig(GetAppLocal())

	routeConfig()
}

// 路由设置
func routeConfig() {
	router := cin.App.GetRouter()
	router.Register(&controllers.TestController{})
}

package cam

import (
	"github.com/go-cam/cam/core/components"
	"github.com/go-cam/cam/core/configs"
	"github.com/go-cam/cam/core/models"
	"github.com/go-cam/cam/core/utils"
)

// new application config
func NewAppConfig() *models.AppConfig {
	appConfig := new(models.AppConfig)
	appConfig.DefaultDBName = "db"
	return appConfig
}

// 新建 websocket server 配置
func NewConfigWebsocketServer(port uint16) *configs.WebsocketServer {
	config := new(configs.WebsocketServer)
	config.Port = port
	config.Component = &components.WebsocketServer{}
	config.MessageParseHandler = nil
	return config
}

// 新建 http server 配置
func NewConfigHttpServer(port uint16) *configs.HttpServer {
	config := new(configs.HttpServer)
	config.Component = &components.HttpServer{}
	config.Port = port
	config.SessionKey = "cam-key"
	config.SessionName = "cam"
	return config
}

// 新建 数据库配置
func NewConfigDatabase(driverName string, host string, port string, name string, username string, password string) *configs.Database {
	config := new(configs.Database)
	config.Component = &components.Database{}
	config.DriverName = driverName
	config.Host = host
	config.Port = port
	config.Name = name
	config.Username = username
	config.Password = password
	config.SetDBFileDir(utils.File.GetRunPath() + "/database")
	// TODO fix path
	config.SetXormTemplateDir("D:\\workspace\\cin\\core\\templates\\xorm")
	config.AutoMigrate = false
	return config
}

// 新建 控制台配置
func NewConfigConsole() *configs.Console {
	config := new(configs.Console)
	config.Component = &components.Console{}
	return config
}

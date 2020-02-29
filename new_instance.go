package cam

import (
	"github.com/go-cam/cam/camConsole"
	"github.com/go-cam/cam/camDatabase"
	"github.com/go-cam/cam/camHttp"
	"github.com/go-cam/cam/camLog"
	"github.com/go-cam/cam/camModels"
	"github.com/go-cam/cam/camWebsocket"
)

// new Application config
func NewAppConfig() *camModels.AppConfig {
	appConfig := new(camModels.AppConfig)
	appConfig.DefaultDBName = "db"
	appConfig.DefaultTemplatesDir = "common/templates"
	return appConfig
}

// new WebsocketComponent config
func NewWebsocketConfig(port uint16) *camWebsocket.WebsocketComponentConfig {
	return camWebsocket.NewWebsocketComponentConfig(port)
}

// new ConsoleComponent config
func NewHttpServerConfig(port uint16) *camHttp.HttpComponentConfig {
	return camHttp.NewHttpComponentConfig(port)
}

// new DatabaseComponent config
func NewDatabaseConfig(driverName string, host string, port string, name string, username string, password string) *camDatabase.DatabaseComponentConfig {
	return camDatabase.NewDatabaseComponentConfig(driverName, host, port, name, username, password)
}

// new ConsoleComponent config
func NewConsoleConfig() *camConsole.ConsoleComponentConfig {
	return camConsole.NewConsoleComponentConfig()
}

// new log config
func NewLogConfig() *camLog.LogComponentConfig {
	config := camLog.NewLogConfig()
	config.Component = new(camLog.LogComponent)
	return config
}

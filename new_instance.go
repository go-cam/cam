package cam

import (
	"github.com/go-cam/cam/camComponents"
	"github.com/go-cam/cam/camConfigs"
	"github.com/go-cam/cam/camConsole"
	"github.com/go-cam/cam/camDatabase"
	"github.com/go-cam/cam/camHttp"
	"github.com/go-cam/cam/camModels"
)

// new application config
func NewAppConfig() *camModels.AppConfig {
	appConfig := new(camModels.AppConfig)
	appConfig.DefaultDBName = "db"
	appConfig.DefaultTemplatesDir = "common/templates"
	return appConfig
}

// new WebsocketServer config
func NewWebsocketServerConfig(port uint16) *camConfigs.WebsocketServer {
	config := new(camConfigs.WebsocketServer)
	config.Port = port
	config.Component = &camComponents.WebsocketServer{}
	config.MessageParseHandler = nil
	config.InitContextPlugin()
	config.InitSslPlugin()
	return config
}

// Deprecated: instead by NewWebsocketServerConfig()
// Remove after 0.3.0
func NewConfigWebsocketServer(port uint16) *camConfigs.WebsocketServer {
	return NewWebsocketServerConfig(port)
}

// new ConsoleComponent config
func NewHttpServerConfig(port uint16) *camHttp.HttpComponentConfig {
	return camHttp.NewHttpComponentConfig(port)
}

// Deprecated: instead by NewHttpServerConfig()
// Remove after 0.3.0
func NewConfigHttpServer(port uint16) *camHttp.HttpComponentConfig {
	return NewHttpServerConfig(port)
}

// new DatabaseComponent config
func NewDatabaseConfig(driverName string, host string, port string, name string, username string, password string) *camDatabase.DatabaseComponentConfig {
	return camDatabase.NewDatabaseComponentConfig(driverName, host, port, name, username, password)
}

// Deprecated: instead by NewDatabaseConfig()
// Remove after 0.3.0
func NewConfigDatabase(driverName string, host string, port string, name string, username string, password string) *camDatabase.DatabaseComponentConfig {
	return NewDatabaseConfig(driverName, host, port, name, username, password)
}

// new ConsoleComponent config
func NewConsoleConfig() *camConsole.ConsoleComponentConfig {
	return camConsole.NewConsoleComponentConfig()
}

// Deprecated: instead by NewConsoleComponentConfig()
// Remove after 0.3.0
func NewConfigConsole() *camConsole.ConsoleComponentConfig {
	return NewConsoleConfig()
}

// new log config
func NewLogConfig() *camConfigs.Log {
	config := camConfigs.NewLogConfig()
	config.Component = new(camComponents.Log)
	return config
}

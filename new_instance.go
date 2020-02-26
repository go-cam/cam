package cam

import (
	"github.com/go-cam/cam/camComponents"
	"github.com/go-cam/cam/camConfigs"
	"github.com/go-cam/cam/camModels"
	"github.com/go-cam/cam/camUtils"
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
// Remove after 1.0.0
func NewConfigWebsocketServer(port uint16) *camConfigs.WebsocketServer {
	return NewWebsocketServerConfig(port)
}

// new HttpServer config
func NewHttpServerConfig(port uint16) *camConfigs.HttpServer {
	config := new(camConfigs.HttpServer)
	config.Component = &camComponents.HttpServer{}
	config.Port = port
	config.SessionKey = "cam-key"
	config.SessionName = "cam"
	config.InitContextPlugin()
	config.InitSslPlugin()
	return config
}

// Deprecated: instead by NewHttpServerConfig()
// Remove after 1.0.0
func NewConfigHttpServer(port uint16) *camConfigs.HttpServer {
	return NewHttpServerConfig(port)
}

// new Database config
func NewDatabaseConfig(driverName string, host string, port string, name string, username string, password string) *camConfigs.Database {
	config := new(camConfigs.Database)
	config.Component = &camComponents.Database{}
	config.DriverName = driverName
	config.Host = host
	config.Port = port
	config.Name = name
	config.Username = username
	config.Password = password
	config.SetDBFileDir(camUtils.File.GetRunPath() + "/database")
	rootPath := App.GetEvn("ROOT_PATH")
	templateDir := App.config.AppConfig.DefaultTemplatesDir
	if rootPath != "" && templateDir != "" {
		//config.SetXormTemplateDir("D:\\workspace\\cin\\core\\templates\\xorm")
		xormTemplateDir := rootPath + "/" + templateDir + "/xorm"
		if camUtils.File.Exists(xormTemplateDir) {
			config.SetXormTemplateDir(xormTemplateDir)
		}
	}
	config.AutoMigrate = false
	return config
}

// Deprecated: instead by NewDatabaseConfig()
// Remove after 1.0.0
func NewConfigDatabase(driverName string, host string, port string, name string, username string, password string) *camConfigs.Database {
	return NewDatabaseConfig(driverName, host, port, name, username, password)
}

// new Console config
func NewConsoleConfig() *camConfigs.Console {
	config := new(camConfigs.Console)
	config.Component = &camComponents.Console{}
	return config
}

// Deprecated: instead by NewConsoleConfig()
// Remove after 1.0.0
func NewConfigConsole() *camConfigs.Console {
	return NewConsoleConfig()
}

// new log config
func NewLogConfig() *camConfigs.Log {
	config := camConfigs.NewLogConfig()
	config.Component = new(camComponents.Log)
	return config
}

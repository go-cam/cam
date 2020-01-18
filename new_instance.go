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
	appConfig.DefaultTemplatesDir = "common/templates"
	return appConfig
}

// new WebsocketServer config
func NewWebsocketServerConfig(port uint16) *configs.WebsocketServer {
	config := new(configs.WebsocketServer)
	config.Port = port
	config.Component = &components.WebsocketServer{}
	config.MessageParseHandler = nil
	return config
}

// Deprecated: instead by NewWebsocketServerConfig()
func NewConfigWebsocketServer(port uint16) *configs.WebsocketServer {
	return NewWebsocketServerConfig(port)
}

// new HttpServer config
func NewHttpServerConfig(port uint16) *configs.HttpServer {
	config := new(configs.HttpServer)
	config.Component = &components.HttpServer{}
	config.Port = port
	config.SessionKey = "cam-key"
	config.SessionName = "cam"
	return config
}

// Deprecated: instead by NewHttpServerConfig()
func NewConfigHttpServer(port uint16) *configs.HttpServer {
	return NewHttpServerConfig(port)
}

// new Database config
func NewDatabaseConfig(driverName string, host string, port string, name string, username string, password string) *configs.Database {
	config := new(configs.Database)
	config.Component = &components.Database{}
	config.DriverName = driverName
	config.Host = host
	config.Port = port
	config.Name = name
	config.Username = username
	config.Password = password
	config.SetDBFileDir(utils.File.GetRunPath() + "/database")
	rootPath := App.GetEvn("ROOT_PATH")
	templateDir := App.config.AppConfig.DefaultTemplatesDir
	if rootPath != "" && templateDir != "" {
		//config.SetXormTemplateDir("D:\\workspace\\cin\\core\\templates\\xorm")
		xormTemplateDir := rootPath + "/" + templateDir + "/xorm"
		if utils.File.Exists(xormTemplateDir) {
			config.SetXormTemplateDir(xormTemplateDir)
		}
	}
	config.AutoMigrate = false
	return config
}

// Deprecated: instead by NewDatabaseConfig()
func NewConfigDatabase(driverName string, host string, port string, name string, username string, password string) *configs.Database {
	return NewDatabaseConfig(driverName, host, port, name, username, password)
}

// new Console config
func NewConsoleConfig() *configs.Console {
	config := new(configs.Console)
	config.Component = &components.Console{}
	return config
}

// Deprecated: instead by NewConsoleConfig()
func NewConfigConsole() *configs.Console {
	return NewConsoleConfig()
}

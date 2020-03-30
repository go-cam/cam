package cam

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camConfig"
	"github.com/go-cam/cam/component/camCache"
	"github.com/go-cam/cam/component/camConsole"
	"github.com/go-cam/cam/component/camDatabase"
	"github.com/go-cam/cam/component/camHttp"
	"github.com/go-cam/cam/component/camLog"
	"github.com/go-cam/cam/component/camMail"
	"github.com/go-cam/cam/component/camWebsocket"
	"github.com/go-cam/cam/template"
)

// new Application config
func NewAppConfig() *camConfig.AppConfig {
	appConfig := new(camConfig.AppConfig)
	appConfig.DefaultDBName = "db"
	return appConfig
}

// new WebsocketComponent config
func NewWebsocketConfig(port uint16) *camWebsocket.WebsocketComponentConfig {
	return camWebsocket.NewWebsocketComponentConfig(port)
}

// new ConsoleComponent config
func NewHttpConfig(port uint16) *camHttp.HttpComponentConfig {
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
	return camLog.NewLogConfig()
}

// new config
func NewConfig() *camConfig.Config {
	return camConfig.NewConfig()
}

// new cache config
func NewCacheConfig() *camCache.CacheComponentConfig {
	return camCache.NewCacheConfig()
}

func NewFileCache() *camCache.FileCache {
	return camCache.NewFileCache()
}

func NewRedisCache() *camCache.RedisCache {
	return camCache.NewRedisEngine()
}

func NewMailConfig(email string, password string, host string) *camMail.MailComponentConfig {
	return camMail.NewMailConfig(email, password, host)
}

func NewTemplateCommand() *template.Command {
	return template.NewCommand()
}

func NewRecover(message string) *camBase.Recover {
	return camBase.NewRecoverable(message)
}

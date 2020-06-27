package cam

// export common constants, structures, methods

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camConfig"
	"github.com/go-cam/cam/base/camConstants"
	"github.com/go-cam/cam/base/camStructs"
	"github.com/go-cam/cam/base/camUtils"
	"github.com/go-cam/cam/component/camCache"
	"github.com/go-cam/cam/component/camConsole"
	"github.com/go-cam/cam/component/camDatabase"
	"github.com/go-cam/cam/component/camHttp"
	"github.com/go-cam/cam/component/camLog"
	"github.com/go-cam/cam/component/camMail"
	"github.com/go-cam/cam/component/camSocket"
	"github.com/go-cam/cam/component/camValidation"
	"github.com/go-cam/cam/component/camWebsocket"
	"github.com/go-cam/cam/plugin/camContext"
	"github.com/go-cam/cam/plugin/camRouter"
	"github.com/go-cam/cam/template"
)

// #################### [START] constant export ####################

// Log
const (
	LogLevelTrace   = camConstants.LevelTrace   // log level: trace
	LogLevelDebug   = camConstants.LevelDebug   // log level: debug
	LogLevelInfo    = camConstants.LevelInfo    // log level: info
	LogLevelWarn    = camConstants.LevelWarn    // log level: warning
	LogLevelError   = camConstants.LevelError   // log level: error
	LogLevelFatal   = camConstants.LevelFatal   // log level: fatal
	LogLevelNone    = camConstants.LevelNone    // none
	LogLevelSuggest = camConstants.LevelSuggest // suggest this level to write file
	LogLevelAll     = camConstants.LevelAll     // all
)

// Validation
const (
	ValidModeInterface = camConstants.ModeInterface // Interface mode
	ValidModeTag       = camConstants.ModeTag       // Tag mode
	ValidModeBot       = camConstants.ModeBoth      // Interface and Tag mode
)

// #################### [END] constant export ####################

// #################### [START] struct export ####################

type Controller struct {
	camRouter.Controller
}

type ConstantController struct {
	camConsole.ConsoleController
}

type HttpController struct {
	camHttp.HttpController
}

type ControllerAction struct {
	camRouter.ControllerAction
}

type Context struct {
	camContext.Context
}

type ValidInterface interface {
	camBase.ValidInterface
}

type MiddlewareInterface interface {
	camBase.MiddlewareInterface
}

// #################### [END] struct export ####################

// #################### [START] new instance func export ####################

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

// new file cache engine
func NewFileCache() *camCache.FileCache {
	return camCache.NewFileCache()
}

// new redis engine
func NewRedisCache() *camCache.RedisCache {
	return camCache.NewRedisEngine()
}

func NewMailConfig(email string, password string, host string) *camMail.MailComponentConfig {
	return camMail.NewMailConfig(email, password, host)
}

func NewCamManager() *template.CamManager {
	return template.NewCamManager()
}

func NewRecover(message string) *camStructs.Recover {
	return camStructs.NewRecoverable(message)
}

// new SocketComponentConfig
func NewSocketConfig(port uint16) *camSocket.SocketComponentConfig {
	return camSocket.NewSocketComponentConfig(port)
}

// new ValidationComponentConfig
func NewValidationConfig() *camValidation.ValidationComponentConfig {
	return camValidation.NewValidationConfig()
}

// new rule
func NewRule(fields []string, handlers ...camBase.ValidHandler) *camStructs.Rule {
	return camStructs.NewRule(fields, handlers...)
}

// #################### [END] new instance func export ####################

// #################### [START] instance export ####################
var Rule = camValidation.Rule

// #################### [END] instance export ####################

// #################### [START] other export ####################
// Framework version
func Version() string {
	return camUtils.C.Version()
}

// #################### [END] other export ####################

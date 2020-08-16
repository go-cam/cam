package cam

// export common constants, structures, methods

import (
	"github.com/go-cam/cam/base/camConfig"
	"github.com/go-cam/cam/base/camStatics"
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
	// Deprecated: Remove on v0.6.0 instead by camStatics.*
	LogLevelTrace   = camStatics.LogLevelTrace // log level: trace
	// Deprecated: Remove on v0.6.0 instead by camStatics.*
	LogLevelDebug   = camStatics.LogLevelDebug // log level: debug
	// Deprecated: Remove on v0.6.0 instead by camStatics.*
	LogLevelInfo    = camStatics.LogLevelInfo // log level: info
	// Deprecated: Remove on v0.6.0 instead by camStatics.*
	LogLevelWarn    = camStatics.LogLevelWarn // log level: warning
	// Deprecated: Remove on v0.6.0 instead by camStatics.*
	LogLevelError   = camStatics.LogLevelError // log level: error
	// Deprecated: Remove on v0.6.0 instead by camStatics.*
	LogLevelFatal   = camStatics.LogLevelFatal // log level: fatal
	// Deprecated: Remove on v0.6.0 instead by camStatics.*
	LogLevelNone    = camStatics.LogLevelNone // none
	// Deprecated: Remove on v0.6.0 instead by camStatics.*
	LogLevelSuggest = camStatics.LogLevelSuggest // suggest this level to write file
	// Deprecated: Remove on v0.6.0 instead by camStatics.*
	LogLevelAll     = camStatics.LogLevelAll // all
)

// Validation
const (
	// Deprecated: Remove on v0.6.0 instead by camStatics.*
	ValidModeInterface = camStatics.ModeInterface // Interface mode
	// Deprecated: Remove on v0.6.0 instead by camStatics.*
	ValidModeTag       = camStatics.ModeTag       // Tag mode
	// Deprecated: Remove on v0.6.0 instead by camStatics.*
	ValidModeBot       = camStatics.ModeBoth      // Interface and Tag mode
)

// #################### [END] constant export ####################

// #################### [START] struct export ####################

// Deprecated: Remove on v0.6.0
type Controller struct {
	camRouter.Controller
}

// Deprecated: Remove on v0.6.0
type ConstantController struct {
	camConsole.ConsoleController
}

// Deprecated: Remove on v0.6.0
type HttpController struct {
	camHttp.HttpController
}

// Deprecated: Remove on v0.6.0
type ControllerAction struct {
	camRouter.ControllerAction
}

// Deprecated: Remove on v0.6.0
type Context struct {
	camContext.Context
}

// Deprecated: Remove on v0.6.0
type ValidInterface interface {
	camStatics.ValidInterface
}

// Deprecated: Remove on v0.6.0
type MiddlewareInterface interface {
	camStatics.MiddlewareInterface
}

// #################### [END] struct export ####################

// #################### [START] new instance func export ####################

// new Application config
func NewAppConfig() *camConfig.AppConfig {
	appConfig := new(camConfig.AppConfig)
	appConfig.DefaultDBName = "db"
	return appConfig
}

// Deprecated: Remove on v0.6.0
// new WebsocketComponent config
func NewWebsocketConfig(port uint16) *camWebsocket.WebsocketComponentConfig {
	return camWebsocket.NewWebsocketComponentConfig(port)
}

// Deprecated: Remove on v0.6.0
// new ConsoleComponent config
func NewHttpConfig(port uint16) *camHttp.HttpComponentConfig {
	return camHttp.NewHttpComponentConfig(port)
}

// Deprecated: Remove on v0.6.0
// new DatabaseComponent config
func NewDatabaseConfig(driverName string, host string, port string, name string, username string, password string) *camDatabase.DatabaseComponentConfig {
	return camDatabase.NewDatabaseComponentConfig(driverName, host, port, name, username, password)
}

// Deprecated: Remove on v0.6.0
// new ConsoleComponent config
func NewConsoleConfig() *camConsole.ConsoleComponentConfig {
	return camConsole.NewConsoleComponentConfig()
}

// Deprecated: Remove on v0.6.0
// new log config
func NewLogConfig() *camLog.LogComponentConfig {
	return camLog.NewLogConfig()
}

// Deprecated: Remove on v0.6.0
// new config
func NewConfig() *camConfig.Config {
	return camConfig.NewConfig()
}

// Deprecated: Remove on v0.6.0
// new cacheComp config
func NewCacheConfig() *camCache.CacheComponentConfig {
	return camCache.NewCacheConfig()
}

// Deprecated: Remove on v0.6.0
// new file cacheComp engine
func NewFileCache() *camCache.FileCache {
	return camCache.NewFileCache()
}

// Deprecated: Remove on v0.6.0
// new redis engine
func NewRedisCache() *camCache.RedisCache {
	return camCache.NewRedisEngine()
}

// Deprecated: Remove on v0.6.0
func NewMailConfig(email string, password string, host string) *camMail.MailComponentConfig {
	return camMail.NewMailConfig(email, password, host)
}

func NewCamManager() *template.CamManager {
	return template.NewCamManager()
}

// Deprecated: Remove on v0.6.0
func NewRecover(message string) *camStructs.Recover {
	return camStructs.NewRecoverable(message)
}

// Deprecated: Remove on v0.6.0
// new SocketComponentConfig
func NewSocketConfig(port uint16) *camSocket.SocketComponentConfig {
	return camSocket.NewSocketComponentConfig(port)
}

// Deprecated: Remove on v0.6.0
// new ValidationComponentConfig
func NewValidationConfig() *camValidation.ValidationComponentConfig {
	return camValidation.NewValidationConfig()
}

// Deprecated: Remove on v0.6.0
// new rule
func NewRule(fields []string, handlers ...camStatics.ValidHandler) *camStructs.Rule {
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

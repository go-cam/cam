package camBase

import (
	"time"
	"xorm.io/xorm"
)

// application interface
// NODE: Provides interface function to the module inner framework
type ApplicationInterface interface {
	// get Component instance by reflect
	GetComponent(v ComponentInterface) ComponentInterface
	// get Component instance by component name
	GetComponentByName(name string) ComponentInterface
	// get default db component's interface
	GetDB() DatabaseComponentInterface
	// run application
	Run()
	// stop application
	Stop()
	// add migration struct
	AddMigration(m MigrationInterface)
	// log trace
	Trace(title string, content string)
	// log debug
	Debug(title string, content string)
	// log info
	Info(title string, content string)
	// log warn
	Warn(title string, content string)
	// log error
	Error(title string, content string)
	// log fatal
	Fatal(title string, content string)
	// Add config
	AddConfig(config AppConfigInterface)
	// get value form .evn file
	GetEvn(key string) string
	// get params form camAppConfig.Config.Params
	GetParam(key string) interface{}
	// get migrate dict
	GetMigrateDict() map[string]MigrationInterface
	// get cache component
	GetCache() CacheComponentInterface
	// get mail component
	GetMail() MailComponentInterface
}

// component config interface
type ComponentConfigInterface interface {
	// new component
	NewComponent() ComponentInterface
	// get recover handler
	GetRecoverHandler() RecoverHandler
}

// Component interface
type ComponentInterface interface {
	// init
	Init(configInterface ComponentConfigInterface)
	// start
	Start()
	// stop
	Stop()
	// set app instance
	SetApp(app ApplicationInterface)
}

// controller interface
type ControllerInterface interface {
	// init
	Init()
	// before action
	BeforeAction(action ControllerActionInterface) bool
	// after action
	AfterAction(action ControllerActionInterface, response []byte) []byte
	// set context
	SetContext(context ContextInterface)
	// get context
	GetContext() ContextInterface
	// set session
	SetSession(session SessionInterface)
	// get session
	GetSession() SessionInterface
	// set values.
	// it will replace the original values
	SetValues(values map[string]interface{})
	// get default action
	GetDefaultActionName() string
	// set response
	SetResponse([]byte)
	// get response
	GetResponse() []byte
}

// controller action interface
type ControllerActionInterface interface {
	// controller route
	Route() string
	// call action
	Call()
}

// context interface
type ContextInterface interface {
}

// session interface
type SessionInterface interface {
	// get sessionId
	GetSessionId() string
	// set key-value in session
	Set(key interface{}, value interface{})
	// get value by key
	Get(key interface{}) interface{}
	// destroy session
	Destroy()
}

// migration interface
type MigrationInterface interface {
	// update migration
	Up()
	// recall migration
	Down()
	// get up sql list
	GetSqlList() []string
}

type PluginConfigInterface interface {
	Init()
}

type PluginInterface interface {
	Init(configInterface PluginConfigInterface)
}

// database component interface
type DatabaseComponentInterface interface {
	// get data source name.
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	GetDSN() string
	// get engine
	GetEngine() *xorm.Engine
	// get xorm session
	NewSession() *xorm.Session
}

// console component interface
type ConsoleComponentInterface interface {
	// get database dir
	GetDatabaseDir() string
	// get xorm template dir
	GetXormTemplateDir() string
}

// app config interface
type AppConfigInterface interface {
}

// cache component interface
type CacheComponentInterface interface {
	// set cache storage default duration
	Set(key string, value interface{}) error
	// set cache storage custom duration
	SetDuration(key string, value interface{}, duration time.Duration) error
	// whether the key exists
	Exists(key string) bool
	// get value by key
	Get(key string) interface{}
	// delete cache
	Del(keys ...string) error
	// delete all cache
	Flush() error
}

// mail component interface
type MailComponentInterface interface {
	Send(subject string, body string, to ...string) error
}

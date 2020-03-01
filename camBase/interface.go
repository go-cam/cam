package camBase

import (
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
	// log info
	Debug(title string, content string)
	// log info
	Info(title string, content string)
	// log info
	Warn(title string, content string)
	// log info
	Error(title string, content string)
	// Add config
	AddConfig(config AppConfigInterface)
	// get value form .evn file
	GetEvn(key string) string
	// get params form camAppConfig.Config.Params
	GetParam(key string) interface{}
	// get migrate dict
	GetMigrateDict() map[string]MigrationInterface
}

// component config interface
type ComponentConfigInterface interface {
	GetComponent() ComponentInterface
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
	// set values.
	// it will replace the original values
	SetValues(values map[string]interface{})
	// append values to values. new value replace old value
	AppendValues(values map[string]interface{})
	// set app instance
	SetApp(app ApplicationInterface)
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
	// set session
	SetSession(session SessionInterface)
	// get session
	GetSession() SessionInterface
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

type AppConfigInterface interface {
}

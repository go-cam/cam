package camBase

import "net/http"

// application interface
// NODE: Provides interface function to the module inner framework
type ApplicationInterface interface {
	// get Component instance by reflect
	GetComponent(v ComponentInterface) ComponentInterface
	// get Component instance by component name
	GetComponentByName(name string) ComponentInterface
	// get default db component's interface
	GetDBInterface() ComponentInterface
}

// component config interface
type ConfigComponentInterface interface {
	GetComponent() ComponentInterface
}

// Component interface
type ComponentInterface interface {
	// init
	Init(configInterface ConfigComponentInterface)
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
	BeforeAction(action string) bool
	// after action
	AfterAction(action string, response []byte) []byte

	// set context
	SetContext(context ContextInterface)
	// get context
	GetContext() ContextInterface

	// set http values by http.ResponseWriter and http.Request
	// 	Q:	what are the values?
	//	A:	values are collection of http's get and post data sent by the client
	SetHttpValues(w http.ResponseWriter, r *http.Request)
	// set values.
	// it will replace the original values
	SetValues(values map[string]interface{})
	// add value
	// it will add key-value to values
	AddValue(key string, value interface{})
	// set app instance
	SetApp(app ApplicationInterface)
	// get action return
	Read() []byte
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

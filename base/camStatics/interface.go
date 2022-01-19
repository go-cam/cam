package camStatics

import (
	"google.golang.org/grpc"
	"net/http"
	"time"
	"xorm.io/xorm"
)

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
	// Deprecated: remove on v0.5.0
	// Instead: GetContext().SetSession();
	SetSession(session SessionInterface)
	// get session
	// Deprecated: remove on v0.5.0
	// Instead: GetContext().GetSession();
	GetSession() SessionInterface
	// set values.
	// it will replace the original values
	SetValues(values map[string]interface{})
	// get default action
	GetDefaultActionName() string
	// set response
	// Deprecated: remove on v0.5.0
	// Instead: GetContext().Write();
	SetResponse([]byte)
	// get response
	// Deprecated: remove on v0.5.0
	// Instead: GetContext().Read();
	GetResponse() []byte
	// set recover
	// Deprecated: remove on v0.5.0
	// Instead: GetContext().SetRecover();
	SetRecover(rec RecoverInterface)
	// get recover
	// Deprecated: remove on v0.5.0
	// Instead: GetContext().GetRecover();
	GetRecover() RecoverInterface
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
	// Write returned data
	Write(res []byte)
	// Read returned data
	Read() []byte
	// Set RecoverInterface. As a error panic by last handle route of handler
	SetRecover(rec RecoverInterface)
	// Get RecoverInterface
	GetRecover() RecoverInterface
	// Set route
	SetRoute(route string)
	// Get route
	GetRoute() string
}

// http context
// it will inject http.ResponseWriter and request *http.Request to context
// Deprecated: remove on v0.5.0
// Instead: HttpContextInterface
type ContextHttpInterface interface {
	ContextInterface
	SetHttpResponseWriter(responseWriter http.ResponseWriter)
	GetHttpResponseWriter() http.ResponseWriter
	SetHttpRequest(request *http.Request)
	GetHttpRequest() *http.Request
}

// http context
type HttpContextInterface interface {
	ContextInterface
	SetHttpResponseWriter(responseWriter http.ResponseWriter)
	GetHttpResponseWriter() http.ResponseWriter
	SetHttpRequest(request *http.Request)
	GetHttpRequest() *http.Request
	CloseHandler(handler func())
	GetCookie(name string) *http.Cookie
	SetCookie(cookie *http.Cookie)
	DelCookie(name string)
	SetCookieValue(name string, value string)
	GetCookieValue(name string) string
	Close()
}

// session interface
type SessionInterface interface {
	// get sessionId
	GetSessionId() string
	// set key-value in session
	Set(key string, value interface{})
	// get value by key
	Get(key string) interface{}
	// delete value by key
	Del(key string)
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

// AppConfigInterface
// Deprecated: remove after v0.6.0-beta
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

// recoverable interface
type RecoverInterface interface {
	Error() string
	GetError() error
}

// valid interface
type ValidInterface interface {
	// get rules
	Rules() []RuleInterface
}

// rule
type RuleInterface interface {
	// get fields of validation
	Fields() []string
	// get handlers of validation
	Handlers() []ValidHandler
}

// validation component interface
type ValidationComponentInterface interface {
	// valid struct
	Valid(v interface{}) map[string][]error
}

// middleware interface
type MiddlewareInterface interface {
	// call when route has middleware
	Handler(ctx ContextInterface, next NextHandler) []byte
}

// grpc component interface
type GRpcComponentInterface interface {
	// get client conn
	Conn() *grpc.ClientConn
}

// Mysql column's builder
type MysqlColumnBuilderInterface interface {
	Unsigned() MysqlColumnBuilderInterface
	NotNull() MysqlColumnBuilderInterface
	Null() MysqlColumnBuilderInterface
	Default(value interface{}) MysqlColumnBuilderInterface
	AutoIncrement() MysqlColumnBuilderInterface
	Comment(comment string) MysqlColumnBuilderInterface
	After(name string) MysqlColumnBuilderInterface
	PrimaryKey() MysqlColumnBuilderInterface
	Index() MysqlColumnBuilderInterface
	Unique() MysqlColumnBuilderInterface
	ToSql() string
	GetKeyPartSql() string
}

type GrpcClientComponentInterface interface {
	GetConn() *grpc.ClientConn
}

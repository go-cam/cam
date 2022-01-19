package camStatics

import "google.golang.org/grpc"

// IApplication
// NODE: Provides interface function to the module inner framework
type IApplication interface {
	// GetComponent get Component instance by reflect
	GetComponent(component IComponent) IComponent
	// GetComponentByName get Component instance by component name
	GetComponentByName(name string) IComponent
	// GetDB get default db component's interface
	GetDB() DatabaseComponentInterface
	// Run run application
	Run()
	// Stop stop application
	Stop()
	// AddMigration add migration struct
	AddMigration(m MigrationInterface)
	// AddConfig
	// Deprecated
	AddConfig(config AppConfigInterface)
	SetApplicationConfig(config IApplicationConfig)

	// Trace log trace
	Trace(title string, content string)
	// Debug log debug
	Debug(title string, content string)
	// Info log info
	Info(title string, content string)
	// Warn log warn
	Warn(title string, content string)
	// Error log error
	Error(title string, content string)
	// Fatal log fatal
	Fatal(title string, content string)

	// GetEvn get value form .evn file
	GetEvn(key string) string
	// GetParam get params form camAppConfig.Config.Params
	GetParam(key string) interface{}
	// GetMigrateDict get migrate dict
	GetMigrateDict() map[string]MigrationInterface
	// GetCache get cache component
	GetCache() CacheComponentInterface
	// GetMail get mail component
	GetMail() MailComponentInterface
	// Valid valid struct.
	Valid(v interface{}) (firstErr error, errDict map[string][]error)
	// GetGrpcClientConn Get grpc client conn
	// name: Component name
	GetGrpcClientConn(name string) *grpc.ClientConn

	// BeforeInit Before app init
	BeforeInit(handler func())
	// AfterInit After app init
	AfterInit(handler func())
	// BeforeStart Before app start
	BeforeStart(handler func())
	// AfterStart After app start
	AfterStart(handler func())
	// BeforeStop Before app stop
	BeforeStop(handler func())
	// AfterStop After app start
	AfterStop(handler func())

	// GetMicroGrpcConn get the micro server *grpc.ClientConn
	// 获取微服务中的 grpc 客户端连接
	GetMicroGrpcConn(appName string) *grpc.ClientConn
}

// IApplicationConfig
// 应用配置
type IApplicationConfig interface {
	// AppName
	// 应用名字
	AppName() string

	// DefaultDBName
	// The component name of the default components.Database
	// 默认数据库名字
	DefaultDBName() string

	// AddComponent
	// name: component name
	// config: component configuration
	AddComponent(name string, config IComponentConfig)

	GetComponentDict() map[string]IComponentConfig
}

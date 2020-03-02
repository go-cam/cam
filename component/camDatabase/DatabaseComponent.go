package camDatabase

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camUtils"
	"github.com/go-cam/cam/component/camConsole"
	"xorm.io/xorm"
)

// database component
type DatabaseComponent struct {
	camBase.Component
	camBase.DatabaseComponentInterface

	config *DatabaseComponentConfig
	engine *xorm.Engine
}

// init
func (component *DatabaseComponent) Init(configInterface camBase.ComponentConfigInterface) {
	component.Component.Init(configInterface)
	var done bool
	component.config, done = configInterface.(*DatabaseComponentConfig)
	if !done {
		panic("configInterface type error. need [*configs.DatabaseComponent]")
	}
	component.engine = nil
}

// start
func (component *DatabaseComponent) Start() {
	component.Component.Start()
}

// stop
func (component *DatabaseComponent) Stop() {
	component.Component.Stop()
}

// create migrations's version record table
func (component *DatabaseComponent) createMigrateVersionTable() error {
	session := component.NewSession()
	migration := new(camConsole.Migration)
	return session.Sync2(migration)
}

// get xorm engine
func (component *DatabaseComponent) GetEngine() *xorm.Engine {
	if component.engine == nil {
		var err error
		component.engine, err = xorm.NewEngine(component.config.DriverName, component.GetDSN())
		camUtils.Error.Panic(err)
	}
	return component.engine
}

// get data source name.
// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
func (component *DatabaseComponent) GetDSN() string {
	host := component.config.Host
	port := component.config.Port
	name := component.config.Name
	username := component.config.Username
	password := component.config.Password

	return username + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?multiStatements=true&charset=utf8"
}

// new session
func (component *DatabaseComponent) NewSession() *xorm.Session {
	return component.GetEngine().NewSession()
}

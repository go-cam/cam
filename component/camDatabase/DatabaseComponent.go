package camDatabase

import (
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/base/camUtils"
	"github.com/go-cam/cam/component"
	"github.com/go-cam/cam/component/camConsole"
	"xorm.io/xorm"
)

// database component
type DatabaseComponent struct {
	component.Component
	camStatics.DatabaseComponentInterface

	config *DatabaseComponentConfig
	engine *xorm.Engine
}

// init
func (comp *DatabaseComponent) Init(configI camStatics.ComponentConfigInterface) {
	comp.Component.Init(configI)
	var ok bool
	comp.config, ok = configI.(*DatabaseComponentConfig)
	if !ok {
		camStatics.App.Fatal("DatabaseComponent", "invalid config")
		return
	}
	comp.engine = nil
}

// start
func (comp *DatabaseComponent) Start() {
	comp.Component.Start()
}

// stop
func (comp *DatabaseComponent) Stop() {
	defer comp.Component.Stop()
}

// create migrations's version record table
func (comp *DatabaseComponent) createMigrateVersionTable() error {
	session := comp.NewSession()
	migration := new(camConsole.Migration)
	return session.Sync2(migration)
}

// get xorm engine
func (comp *DatabaseComponent) GetEngine() *xorm.Engine {
	if comp.engine == nil {
		var err error
		comp.engine, err = xorm.NewEngine(comp.config.DriverName, comp.GetDSN())
		camUtils.Error.Panic(err)
	}
	return comp.engine
}

// get data source name.
// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
func (comp *DatabaseComponent) GetDSN() string {
	host := comp.config.Host
	port := comp.config.Port
	name := comp.config.Name
	username := comp.config.Username
	password := comp.config.Password

	return username + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?multiStatements=true&charset=utf8"
}

// new session
func (comp *DatabaseComponent) NewSession() *xorm.Session {
	return comp.GetEngine().NewSession()
}

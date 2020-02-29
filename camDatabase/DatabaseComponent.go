package camDatabase

import (
	"fmt"
	"github.com/go-cam/cam/camBase"
	"github.com/go-cam/cam/camModels/camModelsTables"
	"github.com/go-cam/cam/camUtils"
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
	if component.config.AutoMigrate {
		component.MigrateUp()
	}
}

// stop
func (component *DatabaseComponent) Stop() {
	component.Component.Stop()
}

// get migrate up version list.
func (component *DatabaseComponent) GetMigrateUpVersionList() []string {
	lastVersion := component.MigrateLastVersion()
	var versionList []string
	for version, _ := range component.config.MigrationDict {
		if version <= lastVersion {
			continue
		}
		versionList = append(versionList, version)
	}
	return versionList
}

// up all database version
func (component *DatabaseComponent) MigrateUp() {
	fmt.Println("Migrate up start.")

	lastVersion := component.MigrateLastVersion()
	var err error
	for version, m := range component.config.MigrationDict {
		if version <= lastVersion {
			continue
		}

		fmt.Print("\tup version: " + version + " ...")

		m.Up()
		sqlList := m.GetSqlList()

		session := component.NewSession()
		err = session.Begin()
		camUtils.Error.Panic(err)

		for _, sqlStr := range sqlList {
			_, err = session.Exec(sqlStr)
			if err != nil {
				panic(err)
			}
		}
		migration := new(camModelsTables.Migration)
		migration.Version = version
		_, err := session.Insert(migration)
		if err != nil {
			_ = session.Rollback()
			camUtils.Error.Panic(err)
		} else {
			err = session.Commit()
			camUtils.Error.Panic(err)
		}

		fmt.Println(" done.")
	}
	fmt.Println("Migrate up finished.")
}

// down last database version
func (component *DatabaseComponent) MigrateDown() {
	lastVersion := component.MigrateLastVersion()
	m, has := component.config.MigrationDict[lastVersion]
	if !has {
		if lastVersion == "" {
			fmt.Println("no version can be down.")
		} else {
			fmt.Println("not found " + lastVersion + "' struct")
		}
		return
	}
	fmt.Println("version: " + lastVersion)
	fmt.Print("Do you want to down this version ?[Y/N]:")
	if !camUtils.Console.IsPressY() {
		return
	}

	m.Down()
	sqlList := m.GetSqlList()

	var err error
	session := component.NewSession()
	err = session.Begin()
	camUtils.Error.Panic(err)
	defer func() {
		if rec := recover(); rec != nil {
			_ = session.Rollback()
		} else {
			_ = session.Commit()
		}
	}()

	for _, sqlStr := range sqlList {
		_, err = session.Exec(sqlStr)
		if err != nil {
			panic(err)
		}
	}

	_, err = session.ID(lastVersion).Delete(camModelsTables.Migration{})
	camUtils.Error.Panic(err)

	fmt.Println(" done.")
}

// get last version
func (component *DatabaseComponent) MigrateLastVersion() string {
	session := component.NewSession()
	exists, err := session.IsTableExist(new(camModelsTables.Migration))
	camUtils.Error.Panic(err)
	if !exists {
		err = component.createMigrateVersionTable()
		camUtils.Error.Panic(err)
		return ""
	}

	migration := new(camModelsTables.Migration)
	has, err := session.Desc("version").Get(migration)
	camUtils.Error.Panic(err)

	version := ""
	if has {
		version = migration.Version
	}
	return version
}

// create migrations's version record table
func (component *DatabaseComponent) createMigrateVersionTable() error {
	session := component.NewSession()
	migration := new(camModelsTables.Migration)
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

// get xorm models's dir
func (component *DatabaseComponent) GetXormModelDir() string {
	return component.config.DBFileDir + "/models"
}

func (component *DatabaseComponent) GetMigrateDir() string {
	return component.config.DBFileDir + "/migrations"
}

// get xorm template dir
func (component *DatabaseComponent) GetXormTemplateDir() string {
	return component.config.XormTemplateDir
}

func (component *DatabaseComponent) GetDatabaseDir() string {
	return component.config.DBFileDir
}

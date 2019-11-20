package components

import (
	"fmt"
	"github.com/cinling/cin/core/base"
	"github.com/cinling/cin/core/configs"
	"github.com/cinling/cin/core/models/tables"
	"github.com/cinling/cin/core/utils"
	"github.com/go-xorm/xorm"
)

// 数据库组件
type Database struct {
	Base

	config *configs.Database
	engine *xorm.Engine
}

// 初始化
func (component *Database) Init(configInterface base.ConfigComponentInterface) {
	component.Base.Init(configInterface)
	var done bool
	component.config, done = configInterface.(*configs.Database)
	if !done {
		panic("configInterface type error. need [*configs.Database]")
	}
	component.engine = nil
}

// 启动
func (component *Database) Start() {
	component.Base.Start()
	component.MigrateUp()
}

// 结束
func (component *Database) Stop() {
	component.Base.Stop()
}

// up all database version
func (component *Database) MigrateUp() {
	lastVersion := component.MigrateLastVersion()
	var err error
	for version, m := range component.config.MigrationDict {
		if version <= lastVersion {
			continue
		}

		fmt.Print("\tMigrate up: " + version)

		m.Up()
		sqlList := m.GetSqlList()

		session := component.GetSession()
		err = session.Begin()
		utils.Error.Panic(err)

		for _, sqlStr := range sqlList {
			_, err = session.Exec(sqlStr)
			if err != nil {
				panic(err)
			}
		}
		migration := new(tables.Migration)
		migration.Version = version
		_, err := session.Insert(migration)
		if err != nil {
			_ = session.Rollback()
			utils.Error.Panic(err)
		} else {
			err = session.Commit()
			utils.Error.Panic(err)
		}

		fmt.Println(" done.")
	}
}

// TODO
// down last database version
func (component *Database) MigrateDown() {

}

// get last version
func (component *Database) MigrateLastVersion() string {
	session := component.GetSession()
	exists, err := session.IsTableExist(new(tables.Migration))
	utils.Error.Panic(err)
	if !exists {
		err = component.createMigrateVersionTable()
		utils.Error.Panic(err)
		return ""
	}

	migration := new(tables.Migration)
	has, err := session.Desc("version").Get(migration)
	utils.Error.Panic(err)

	version := ""
	if has {
		version = migration.Version
	}
	return version
}

// create migrations's version record table
func (component *Database) createMigrateVersionTable() error {
	session := component.GetSession()
	migration := new(tables.Migration)
	return session.Sync2(migration)
}

// get xorm engine
func (component *Database) GetEngine() *xorm.Engine {
	if component.engine == nil {
		var err error
		component.engine, err = xorm.NewEngine(component.config.DriverName, component.GetDSN())
		utils.Error.Panic(err)
	}
	return component.engine
}

// get data source name.
// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
func (component *Database) GetDSN() string {
	host := component.config.Host
	port := component.config.Port
	name := component.config.Name
	username := component.config.Username
	password := component.config.Password

	return username + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?multiStatements=true&charset=utf8"
}

// TODO not testing
func (component *Database) GetSession() *xorm.Session {
	return component.GetEngine().NewSession()
}

// get xorm models's dir
func (component *Database) GetXormModelDir() string {
	return component.config.DBFileDir + "/models"
}

func (component *Database) GetMigrateDir() string {
	return component.config.DBFileDir + "/migrations"
}

// get xorm template dir
func (component *Database) GetXormTemplateDir() string {
	return component.config.XormTemplateDir
}

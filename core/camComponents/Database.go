package camComponents

import (
	"bufio"
	"fmt"
	"github.com/go-cam/cam/core/camBase"
	"github.com/go-cam/cam/core/camConfigs"
	"github.com/go-cam/cam/core/camModels/camModelsTables"
	"github.com/go-cam/cam/core/camUtils"
	"os"
	"strings"
	"xorm.io/xorm"
)

// 数据库组件
type Database struct {
	Base

	config *camConfigs.Database
	engine *xorm.Engine
}

// 初始化
func (component *Database) Init(configInterface camBase.ConfigComponentInterface) {
	component.Base.Init(configInterface)
	var done bool
	component.config, done = configInterface.(*camConfigs.Database)
	if !done {
		panic("configInterface type error. need [*configs.Database]")
	}
	component.engine = nil
}

// 启动
func (component *Database) Start() {
	component.Base.Start()
	if component.config.AutoMigrate {
		component.MigrateUp()
	}
}

// 结束
func (component *Database) Stop() {
	component.Base.Stop()
}

// up all database version
func (component *Database) MigrateUp() {
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

// TODO
// down last database version
func (component *Database) MigrateDown() {
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
	fmt.Print("Do you want to down this version: " + lastVersion + " ?[Y/N]:")
	input := bufio.NewScanner(os.Stdin)
	if !input.Scan() {
		return
	}
	str := strings.ToLower(input.Text())
	if str != "y" {
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

	fmt.Println("\tdone.")
}

// get last version
func (component *Database) MigrateLastVersion() string {
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
func (component *Database) createMigrateVersionTable() error {
	session := component.NewSession()
	migration := new(camModelsTables.Migration)
	return session.Sync2(migration)
}

// get xorm engine
func (component *Database) GetEngine() *xorm.Engine {
	if component.engine == nil {
		var err error
		component.engine, err = xorm.NewEngine(component.config.DriverName, component.GetDSN())
		camUtils.Error.Panic(err)
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
func (component *Database) NewSession() *xorm.Session {
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

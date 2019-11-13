package components

import (
	"database/sql"
	"errors"
	"github.com/cinling/cin/core/base"
	"github.com/cinling/cin/core/configs"
	"github.com/cinling/cin/core/constants"
	"github.com/cinling/cin/core/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
}

// 启动
func (component *Database) Start() {
	component.Base.Start()
	component.initEngine()
	component.MigrateUp()
}

// 结束
func (component *Database) Stop() {
	component.Base.Stop()
}

// 升级数据库
func (component *Database) MigrateUp() {
	migrateDir := component.config.GetMigrateDir()
	if migrateDir == "" {
		// 如果没有设置 migrate 路径，则不进行更新
		return
	}
	driverName := component.config.DriverName
	host := component.config.Host
	port := component.config.Port
	name := component.config.Name
	username := component.config.Username
	password := component.config.Password

	driver, err := component.getDriver(driverName, host, port, name, username, password)
	if err != nil {
		panic(err.Error())
	}
	m, err := migrate.NewWithDatabaseInstance("file://"+migrateDir, driverName, driver)
	if err != nil {
		panic(err.Error())
	}
	_, dirty, err := m.Version()
	if dirty {
		// 如果已经被弄脏，则尝试强制修复
		// fix dirty version. but it's not work. must be fixed manually
		err := m.Steps(-1)
		utils.Error.Panic(err)
	}

	err = m.Up()

	if err == migrate.ErrNoChange {
		// 不需要升级的情况
		err = nil
	}
	utils.Error.Panic(err)
}

// 通过 driverName 获取 driver
func (component *Database) getDriver(driverName string, host string, port string, name string, username string, password string) (database.Driver, error) {
	if driverName == constants.DatabaseDriverMysql {
		db, err := sql.Open(driverName, component.GetDSN())
		if err != nil {
			return nil, err
		}
		driver, err := mysql.WithInstance(db, &mysql.Config{})
		if err != nil {
			return nil, err
		}
		return driver, nil
	} else {
		return nil, errors.New("not support driver: " + driverName)
	}
}

// get xorm engine
func (component *Database) GetEngine() *xorm.Engine {
	return component.engine
}

// init xorm engine
func (component *Database) initEngine() {
	var err error
	component.engine, err = xorm.NewEngine(component.config.DriverName, component.GetDSN())
	utils.Error.Panic(err)
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

func (component *Database) GetSession() *xorm.Session {
	component.engine.Sync()
	return component.engine.NewSession()
}

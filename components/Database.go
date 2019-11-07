package components

import (
	"database/sql"
	"errors"
	"github.com/cinling/cin/base"
	"github.com/cinling/cin/configs"
	"github.com/cinling/cin/constants"
	"github.com/cinling/cin/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"reflect"
)

// 数据库组件
type Database struct {
	Base

	config *configs.Database
}

// 初始化
func (component *Database) Init(configInterface base.ConfigComponentInterface) {
	component.Base.Init(configInterface)

	configValue := reflect.ValueOf(configInterface)
	var config *configs.Database
	if configValue.Kind() == reflect.Ptr {
		config = configValue.Interface().(*configs.Database)
	} else if configValue.Kind() == reflect.Struct {
		configStruct := configValue.Interface().(configs.Database)
		config = &configStruct
	} else {
		panic("illegal config")
	}
	component.config = config
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

// 升级数据库
func (component *Database) MigrateUp() {
	migrateDir := component.config.MigrateDir
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
		db, err := sql.Open(driverName, username+":"+password+"@tcp("+host+":"+port+")/"+name+"?multiStatements=true")
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

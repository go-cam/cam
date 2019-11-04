package components

import (
	"cin/base"
	"cin/configs"
	"cin/constants"
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database"
	"github.com/golang-migrate/migrate/database/mysql"
)

// 数据库组件
type Database struct {
	Base

	config *configs.Database
}

// 初始化
func (component *Database) Init(configInterface base.ConfigComponentInterface) {
	component.Base.Init(configInterface)
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

	db, err := sql.Open(driverName, username + ":" + password + "@tcp(" + host + ":" + port + ")/" + name +  "?multiStatements=true")
	if err != nil {
		panic(err.Error())
	}
	driver, err := component.getDriver(db, driverName)
	if err != nil {
		panic(err.Error())
	}
	m, err := migrate.NewWithDatabaseInstance("file://" + migrateDir, driverName, driver)
	if err != nil {
		panic(err.Error())
	}
	err = m.Steps(10000)
	if err != nil {
		panic(err.Error())
	}
}

// 通过 driverName 获取 driver
func (component *Database) getDriver(db *sql.DB, driverName string) (database.Driver, error) {
	if driverName == constants.DatabaseDriverMysql {
		driver, err := mysql.WithInstance(db, &mysql.Config{})
		if err != nil {
			return nil, err
		}
		return driver, nil
	} else {
		return nil, errors.New("not support driver: " + driverName)
	}
}
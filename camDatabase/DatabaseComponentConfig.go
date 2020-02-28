package camDatabase

import (
	"github.com/go-cam/cam/camBase"
	"github.com/go-cam/cam/camConfigs"
	"github.com/go-cam/cam/camUtils"
	"strings"
)

// database config
type DatabaseComponentConfig struct {
	camBase.Config
	camConfigs.MigratePlugin

	DriverName      string // driver name. Example: "mysql", "sqlite"
	Host            string // database hostname
	Port            string // database port
	Name            string // database name
	Username        string // username
	Password        string // password
	DBFileDir       string // DatabaseComponentConfig file storage path. Default is: /[path to run dir]/database
	XormTemplateDir string // xorm template path.
	AutoMigrate     bool   // run migrate up on component startup
}

// new instance
func NewDatabaseComponentConfig(driverName string, host string, port string, name string, username string, password string) *DatabaseComponentConfig {
	config := new(DatabaseComponentConfig)
	config.Component = &DatabaseComponent{}
	config.DriverName = driverName
	config.Host = host
	config.Port = port
	config.Name = name
	config.Username = username
	config.Password = password
	runPath := camUtils.File.GetRunPath()
	rootPath := camUtils.File.Dir(runPath)
	config.SetDBFileDir(runPath + "/database")
	config.SetXormTemplateDir(rootPath + "/common/templates")
	config.AutoMigrate = false
	return config
}

// set migration's file dir
func (config *DatabaseComponentConfig) SetDBFileDir(dir string) *DatabaseComponentConfig {
	config.DBFileDir = strings.Replace(dir, "\\", "/", -1)
	return config
}

// set xorm dir
func (config *DatabaseComponentConfig) SetXormTemplateDir(dir string) *DatabaseComponentConfig {
	config.XormTemplateDir = strings.Replace(dir, "\\", "/", -1)
	return config
}

// migrate up on component startup
func (config *DatabaseComponentConfig) SetAutoMigrate() *DatabaseComponentConfig {
	config.AutoMigrate = true
	return config
}

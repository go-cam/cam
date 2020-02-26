package camConfigs

import "strings"

// database config
type Database struct {
	BaseConfig
	MigratePlugin

	DriverName      string // driver name. Example: "mysql", "sqlite"
	Host            string // database hostname
	Port            string // database port
	Name            string // database name
	Username        string // username
	Password        string // password
	DBFileDir       string // Database file storage path. Default is: /[path to run dir]/database
	XormTemplateDir string // xorm template path.
	AutoMigrate     bool   // run migrate up on component startup
}

// set migration's file dir
func (config *Database) SetDBFileDir(dir string) *Database {
	config.DBFileDir = strings.Replace(dir, "\\", "/", -1)
	return config
}

// set xorm dir
func (config *Database) SetXormTemplateDir(dir string) *Database {
	config.XormTemplateDir = strings.Replace(dir, "\\", "/", -1)
	return config
}

// migrate up on component startup
func (config *Database) SetAutoMigrate() *Database {
	config.AutoMigrate = true
	return config
}

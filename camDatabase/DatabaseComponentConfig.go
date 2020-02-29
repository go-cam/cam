package camDatabase

import (
	"github.com/go-cam/cam/camBase"
)

// database config
type DatabaseComponentConfig struct {
	camBase.ComponentConfig

	DriverName string // driver name. Example: "mysql", "sqlite"
	Host       string // database hostname
	Port       string // database port
	Name       string // database name
	Username   string // username
	Password   string // password
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
	return config
}

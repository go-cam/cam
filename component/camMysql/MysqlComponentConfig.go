package camMysql

import "github.com/go-cam/cam/component"

type MysqlComponentConfig struct {
	component.ComponentConfig
}

func NewMysqlComponentConfig() *MysqlComponentConfig {
	conf := new(MysqlComponentConfig)
	conf.Component = &MysqlComponent{}
	return conf
}

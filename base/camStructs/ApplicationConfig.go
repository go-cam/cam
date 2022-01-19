package camStructs

import "github.com/go-cam/cam/base/camStatics"

func NewApplicationConfig(appName string) *ApplicationConfig {
	conf := new(ApplicationConfig)
	conf.appName = appName
	conf.defaultDBName = "db"
	conf.params = map[string]interface{}{}
	conf.componentDict = map[string]camStatics.IComponentConfig{}
	return conf
}

type ApplicationConfig struct {
	camStatics.IApplicationConfig
	appName       string
	defaultDBName string
	params        map[string]interface{}
	componentDict map[string]camStatics.IComponentConfig
}

func (c *ApplicationConfig) AppName() string {
	return c.appName
}

func (c *ApplicationConfig) DefaultDBName() string {
	return c.defaultDBName
}

func (c *ApplicationConfig) AddComponent(name string, config camStatics.IComponentConfig) {
	c.componentDict[name] = config
}

func (c *ApplicationConfig) GetComponentDict() map[string]camStatics.IComponentConfig {
	return c.componentDict
}

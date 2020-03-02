package camConsole

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camUtils"
	"github.com/go-cam/cam/plugin/camPluginRouter"
	"strings"
)

// console config
type ConsoleComponentConfig struct {
	camBase.ComponentConfig
	camPluginRouter.RouterPluginConfig

	DatabaseDir     string // DatabaseComponentConfig file storage path. Default is: /[path to run dir]/database
	XormTemplateDir string // xorm template path.
}

// new console config
func NewConsoleComponentConfig() *ConsoleComponentConfig {
	config := new(ConsoleComponentConfig)
	config.Component = &ConsoleComponent{}
	runPath := camUtils.File.GetRunPath()
	rootPath := camUtils.File.Dir(runPath)
	config.SetDatabaseDir(runPath + "/database")
	config.SetXormTemplateDir(rootPath + "/common/templates")
	config.RouterPluginConfig.Init()
	config.registerFrameworkController()
	return config
}

// register controller in the framework
func (config *ConsoleComponentConfig) registerFrameworkController() {
	config.RouterPluginConfig.Register(&MigrateController{})
	config.RouterPluginConfig.Register(&XormController{})
}

// set migration's file dir
func (config *ConsoleComponentConfig) SetDatabaseDir(dir string) *ConsoleComponentConfig {
	config.DatabaseDir = strings.Replace(dir, "\\", "/", -1)
	return config
}

// set xorm dir
func (config *ConsoleComponentConfig) SetXormTemplateDir(dir string) *ConsoleComponentConfig {
	config.XormTemplateDir = strings.Replace(dir, "\\", "/", -1)
	return config
}

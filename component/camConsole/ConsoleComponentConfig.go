package camConsole

import (
	"github.com/go-cam/cam/base/camUtils"
	"github.com/go-cam/cam/component"
	"github.com/go-cam/cam/plugin/camContext"
	"github.com/go-cam/cam/plugin/camRouter"
	"strings"
)

// console config
type ConsoleComponentConfig struct {
	component.ComponentConfig
	camRouter.RouterPluginConfig
	camContext.ContextPluginConfig

	DatabaseDir     string // DatabaseComponentConfig file storage path. Default is: /[path to run dir]/database
	XormTemplateDir string // Deprecated: xorm template path.
	grpcOption      *GrpcOption
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
	config.SetContextStruct(&camContext.Context{})
	config.grpcOption = &GrpcOption{}
	return config
}

// register controller in the framework
func (conf *ConsoleComponentConfig) registerFrameworkController() {
	conf.RouterPluginConfig.Register(&MigrateController{})
	conf.RouterPluginConfig.Register(&XormController{})
}

// set migration's file dir
func (conf *ConsoleComponentConfig) SetDatabaseDir(dir string) *ConsoleComponentConfig {
	conf.DatabaseDir = strings.Replace(dir, "\\", "/", -1)
	return conf
}

// set xorm dir
func (conf *ConsoleComponentConfig) SetXormTemplateDir(dir string) *ConsoleComponentConfig {
	conf.XormTemplateDir = strings.Replace(dir, "\\", "/", -1)
	return conf
}

// set grpc option
func (conf *ConsoleComponentConfig) SetGrpcOption(option *GrpcOption) {
	conf.grpcOption = option
}

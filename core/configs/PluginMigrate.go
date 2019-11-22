package configs

import (
	"github.com/go-cam/cam/core/base"
	"github.com/go-cam/cam/core/utils"
)

// migrate plugin
type PluginMigrate struct {
	MigrationDict map[string]base.MigrationInterface
}

// add migration struct
func (plugin *PluginMigrate) Add(m base.MigrationInterface) {
	id := utils.Reflect.GetStructName(m)
	plugin.MigrationDict[id] = m
}

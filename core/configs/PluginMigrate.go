package configs

import (
	"github.com/cinling/cin/core/base"
	"github.com/cinling/cin/core/utils"
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

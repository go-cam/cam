package camConfigs

import (
	"github.com/go-cam/cam/core/camBase"
	"github.com/go-cam/cam/core/camUtils"
)

// migrate plugin
type PluginMigrate struct {
	MigrationDict map[string]camBase.MigrationInterface
}

// add migration struct
func (plugin *PluginMigrate) Add(m camBase.MigrationInterface) {
	id := camUtils.Reflect.GetStructName(m)
	plugin.MigrationDict[id] = m
}

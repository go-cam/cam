package camConfigs

import (
	"github.com/go-cam/cam/camBase"
	"github.com/go-cam/cam/camUtils"
)

// migrate plugin
// Deprecated: remove on v0.3.0
type MigratePlugin struct {
	MigrationDict map[string]camBase.MigrationInterface
}

// add migration struct
func (plugin *MigratePlugin) Add(m camBase.MigrationInterface) {
	id := camUtils.Reflect.GetStructName(m)
	plugin.MigrationDict[id] = m
}

package camMiddleware

import (
	"github.com/go-cam/cam/base/camBase"
	"strings"
)

// middleware plugin
type MiddlewarePlugin struct {
	camBase.PluginInterface
	conf *MiddlewarePluginConfig
}

// init
func (plugin *MiddlewarePlugin) Init(config *MiddlewarePluginConfig) {
	plugin.conf = config
}

// get list of camBase.MiddlewareInterface
func (plugin *MiddlewarePlugin) GetMiddlewareList(route string) []camBase.MiddlewareInterface {
	var midIList []camBase.MiddlewareInterface
	for prefix, midI := range plugin.conf.middlewareDict {
		if strings.HasPrefix(route, prefix) {
			midIList = append(midIList, midI)
		}
	}
	return midIList
}

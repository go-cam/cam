package camMiddleware

import (
	"github.com/go-cam/cam/base/camStatics"
)

// middleware plugin conf
type MiddlewarePluginConfig struct {
	camStatics.PluginConfigInterface
	middlewareDict map[string][]camStatics.MiddlewareInterface
}

func (config *MiddlewarePluginConfig) Init() {
	// default middleware
	config.middlewareDict = map[string][]camStatics.MiddlewareInterface{}
}

// add middleware
func (config *MiddlewarePluginConfig) AddMiddleware(prefix string, midI camStatics.MiddlewareInterface) *MiddlewarePluginConfig {
	_, has := config.middlewareDict[prefix]
	if !has {
		config.middlewareDict[prefix] = []camStatics.MiddlewareInterface{}
	}
	config.middlewareDict[prefix] = append(config.middlewareDict[prefix], midI)
	return config
}

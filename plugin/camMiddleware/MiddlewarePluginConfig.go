package camMiddleware

import "github.com/go-cam/cam/base/camBase"

// middleware plugin conf
type MiddlewarePluginConfig struct {
	camBase.PluginConfigInterface
	middlewareDict map[string][]camBase.MiddlewareInterface
}

func (config *MiddlewarePluginConfig) Init() {
	// default middleware
	config.middlewareDict = map[string][]camBase.MiddlewareInterface{}
}

// add middleware
func (config *MiddlewarePluginConfig) AddMiddleware(prefix string, midI camBase.MiddlewareInterface) *MiddlewarePluginConfig {
	_, has := config.middlewareDict[prefix]
	if !has {
		config.middlewareDict[prefix] = []camBase.MiddlewareInterface{}
	}
	config.middlewareDict[prefix] = append(config.middlewareDict[prefix], midI)
	return config
}

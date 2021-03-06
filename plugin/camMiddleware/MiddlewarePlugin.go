package camMiddleware

import (
	"github.com/go-cam/cam/base/camStatics"
	"reflect"
	"strings"
)

// middleware plugin
type MiddlewarePlugin struct {
	camStatics.PluginInterface
	conf *MiddlewarePluginConfig
}

// init
func (plugin *MiddlewarePlugin) Init(config *MiddlewarePluginConfig) {
	plugin.conf = config
}

// get list of camBase.MiddlewareInterface
func (plugin *MiddlewarePlugin) GetMiddlewareList(route string) []camStatics.MiddlewareInterface {
	var midIList []camStatics.MiddlewareInterface
	for prefix, tmpMidIList := range plugin.conf.middlewareDict {
		if strings.HasPrefix(route, prefix) {
			for _, midI := range tmpMidIList {
				midIList = append(midIList, midI)
			}
		}
	}
	return midIList
}

// call next with middleware
func (plugin *MiddlewarePlugin) CallWithMiddleware(ctx camStatics.ContextInterface, route string, next camStatics.NextHandler) []byte {
	midIList := plugin.GetMiddlewareList(route)
	return plugin.recursionCall(ctx, midIList, next)
}

// recursion call
func (plugin *MiddlewarePlugin) recursionCall(ctx camStatics.ContextInterface, midIList []camStatics.MiddlewareInterface, next camStatics.NextHandler) []byte {
	length := len(midIList)
	if length == 0 {
		return next()
	}

	midI := midIList[0]
	midIList = midIList[1:]
	midV := reflect.New(reflect.TypeOf(midI).Elem())
	mid, ok := midV.Interface().(camStatics.MiddlewareInterface)
	if !ok {
		panic("middleware not implements camBase.MiddlewareInterface")
	}

	return plugin.recursionCall(ctx, midIList, func() []byte {
		return mid.Handler(ctx, next)
	})
}

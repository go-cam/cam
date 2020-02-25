package camComponents

import (
	"github.com/go-cam/cam/camConfigs"
	"reflect"
)

// router plugin.
type RouterPlugin struct {
	config               *camConfigs.RouterPlugin   // router config
	controllerDict       map[string]reflect.Type    // controller reflect.Type dict
	controllerActionDict map[string]map[string]bool // map[controllerName]map[actionName]
}

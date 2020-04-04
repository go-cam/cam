// Help framework easy start, default setting, simplified application method
package cam

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/plugin/camPluginRouter"
)

// record some values
type defaultAo struct {
	camPluginRouter.RouterPluginConfig

	compDict map[string]camBase.ComponentConfigInterface
}

var ao = newDefaultAo()

func newDefaultAo() *defaultAo {
	ao := new(defaultAo)
	ao.compDict = map[string]camBase.ComponentConfigInterface{}
	return ao
}

// default http server config
func defaultHttpConfig() camBase.ComponentConfigInterface {
	conf := NewHttpConfig(20200)
	conf.SessionName = "cam"
	conf.RouterPluginConfig = ao.RouterPluginConfig
	return conf
}

// default websocket server config
func defaultWebsocketConfig() camBase.ComponentConfigInterface {
	conf := NewWebsocketConfig(20201)
	conf.RouterPluginConfig = ao.RouterPluginConfig
	return conf
}

// default tcp socket server config
func defaultSocketConfig() camBase.ComponentConfigInterface {
	conf := NewSocketConfig(20202)
	conf.Trace = true
	conf.RouterPluginConfig = ao.RouterPluginConfig
	return conf
}

// must run before cam.RunDefault
func AddComponent(name string, conf camBase.ComponentConfigInterface) {
	ao.compDict[name] = conf
}

// must run before cam.RunDefault register controller
func RegisterController(ctrl camBase.ControllerInterface) {
	ao.Register(ctrl)
}

// run application
func RunDefault() {
	conf := NewConfig()
	conf.ComponentDict = map[string]camBase.ComponentConfigInterface{
		"http": defaultHttpConfig(),
		"ws":   defaultWebsocketConfig(),
		"tcp":  defaultSocketConfig(),
	}
	for compName, comp := range ao.compDict {
		conf.ComponentDict[compName] = comp
	}
	App.AddConfig(conf)

	App.Run()
}

// get env file values
func Env(key string) string {
	return App.GetEvn(key)
}

// get config param
func Param(key string) interface{} {
	return App.GetParam(key)
}

// trace log
func Trace(title, content string) {
	App.Trace(title, content)
}

// debug log
func Debug(title, content string) {
	App.Debug(title, content)
}

// info log
func Info(title, content string) {
	App.Info(title, content)
}

// warn log
func Warn(title, content string) {
	App.Warn(title, content)
}

// error log
func Error(title, content string) {
	App.Error(title, content)
}

// fatal log
func Fatal(title, content string) {
	App.Fatal(title, content)
}

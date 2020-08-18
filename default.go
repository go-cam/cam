package cam

// Help framework easy start, default setting, simplified application method

import (
	"github.com/go-cam/cam/base/camConfig"
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/component/camHttp"
	"github.com/go-cam/cam/component/camSocket"
	"github.com/go-cam/cam/component/camWebsocket"
	"github.com/go-cam/cam/plugin/camRouter"
)

// record some values
type defaultAo struct {
	camRouter.RouterPluginConfig

	compDict map[string]camStatics.ComponentConfigInterface
}

var ao = newDefaultAo()

func newDefaultAo() *defaultAo {
	ao := new(defaultAo)
	ao.compDict = map[string]camStatics.ComponentConfigInterface{}
	return ao
}

// default http server config
func defaultHttpConfig() camStatics.ComponentConfigInterface {
	conf := camHttp.NewHttpComponentConfig(20200)
	conf.RouterPluginConfig = ao.RouterPluginConfig
	return conf
}

// default websocket server config
func defaultWebsocketConfig() camStatics.ComponentConfigInterface {
	conf := camWebsocket.NewWebsocketComponentConfig(20201)
	conf.RouterPluginConfig = ao.RouterPluginConfig
	return conf
}

// default tcp socket server config
func defaultSocketConfig() camStatics.ComponentConfigInterface {
	conf := camSocket.NewSocketComponentConfig(20202)
	conf.Trace = true
	conf.RouterPluginConfig = ao.RouterPluginConfig
	return conf
}

// must run before cam.RunDefault
func AddComponent(name string, conf camStatics.ComponentConfigInterface) {
	ao.compDict[name] = conf
}

// must run before cam.RunDefault register controller
func RegisterController(ctrl camStatics.ControllerInterface) {
	ao.Register(ctrl)
}

// run application
func RunDefault() {
	conf := camConfig.NewConfig()
	conf.ComponentDict = map[string]camStatics.ComponentConfigInterface{
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

// validComp struct
func Valid(v interface{}) (firstErr error, errDict map[string][]error) {
	return App.Valid(v)
}

package cin

import (
	"cin/base"
	"cin/configs"
	"reflect"
	"time"
)

// 服务器全局类
type application struct {
	// 应用状态（初始化、开始、运行中、停止、销毁）[onInit, onStart, onRun, onStop, onDestroy]
	status base.ApplicationStatus
	// 应用全部配置
	config *Config
	// 应用组件名
	componentDict map[string]base.ComponentInterface
	// url 路游戏（websocket、http通用）
	router *router
}

// 应用全局实例（只需要一个实例即可操作整个应用）
var App = newApplication()

// 框架方法
func newApplication() *application {
	app := new(application)
	app.status = ApplicationStatusInit
	app.config = NewConfig()
	app.componentDict = map[string]base.ComponentInterface{}
	app.router = newRouter()
	return app
}

// 添加配置（有先后顺序。后面的配置将覆盖前面的配置）
func (app *application) AddConfig(config *Config) {
	for key, value := range config.Params {
		app.config.Params[key] = value
	}
	for name, componentConfig := range config.ComponentDict {
		app.config.ComponentDict[name] = componentConfig
	}
}

// 获取路游戏
func (app *application) GetRouter() *router {
	return app.router
}

// 启动应用
func (app *application) Run() {
	app.onInit()
	app.onStart()
	app.wait()
	app.onStop()
}

// 应用初始化
func (app *application) onInit() {
	// 实例化组件
	for name, config := range app.config.ComponentDict {
		// 使用结构体重新创建一个新对象（防止外部修改导致混乱）
		componentInterface := config.GetComponent()
		t := reflect.TypeOf(componentInterface)
		componentType := t.Elem()
		componentValue := reflect.New(componentType)
		componentInterface = componentValue.Interface().(base.ComponentInterface)

		// 写入插件的数据
		app.writePluginParams(config)

		componentInterface.Init(config)
		app.componentDict[name] = componentInterface
	}
}

// 应用开始（初始化组件）
func (app *application) onStart() {
	for _, component := range app.componentDict {
		go component.Start()
	}
}

// 应用向所有组件发送停止信号
func (app *application) onStop() {
	for _, component := range app.componentDict {
		component.Stop()
	}
}

// 等待（不会让程序结束）
func (app *application) wait() {
	for {
		time.Sleep(1 * time.Second)
	}
}

// 自动写入插件配置
func (app *application) writePluginParams(config base.ConfigComponentInterface) {
	t := reflect.TypeOf(config).Elem()
	v := reflect.ValueOf(config).Elem()
	// 写入路由插件数据
	if _, has := t.FieldByName("PluginRouter"); has {
		pluginRouter := v.FieldByName("PluginRouter").Interface().(configs.PluginRouter)
		pluginRouter.HandlerList = app.router.handlerList
		pluginRouter.OnWebsocketMessageHandler = app.router.onWebsocketMessageHandler
		v.FieldByName("PluginRouter").Set(reflect.ValueOf(pluginRouter))
	}
}

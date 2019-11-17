package cin

import (
	"fmt"
	"github.com/cinling/cin/core/base"
	"github.com/cinling/cin/core/components"
	"github.com/cinling/cin/core/configs"
	"github.com/cinling/cin/core/utils"
	"os"
	"reflect"
	"time"
)

// 服务器全局类
type application struct {
	base.ApplicationInterface

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
	app.config.AppConfig = NewAppConfig()
	app.componentDict = map[string]base.ComponentInterface{}
	app.router = newRouter()
	return app
}

// Add config
// Merge as much as possible, otherwise overwrite.
//
// config: new config
func (app *application) AddConfig(config *Config) {
	for key, value := range config.Params {
		app.config.Params[key] = value
	}
	for name, componentConfig := range config.ComponentDict {
		app.config.ComponentDict[name] = componentConfig
	}

	if config.AppConfig != nil {
		app.config.AppConfig = config.AppConfig
	}
}

// 获取路游戏
func (app *application) GetRouter() *router {
	return app.router
}

// 启动应用
func (app *application) Run() {
	fmt.Println("App: Initializing ...")
	app.onInit()
	if len(os.Args) >= 2 {
		// 如果运行参数大于1个，说明是一个一次单独的命令，不启动服务
		app.callConsole()
		return
	}
	fmt.Println("App: Starting up ...")
	app.onStart()
	fmt.Println("App: Startup done.")
	app.wait()
	fmt.Println("App: Stop now...")
	app.onStop()
	fmt.Println("App: Stop done. Application was exit.")
}

// 应用初始化
func (app *application) onInit() {
	components.SetApplication(app)

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
		componentInterface.SetApp(app)
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
		pluginRouter.ControllerList = app.router.controllerList
		pluginRouter.ConsoleControllerList = app.router.consoleControllerList
		pluginRouter.OnWebsocketMessageHandler = app.router.onWebsocketMessageHandler
		v.FieldByName("PluginRouter").Set(reflect.ValueOf(pluginRouter))
	}
}

// 调用命令行组件
func (app *application) callConsole() {
	isCallConsole := false

	for _, componentIns := range app.componentDict {
		name := utils.Reflect.GetClassName(componentIns)
		if name == "Console" {
			isCallConsole = true
			consoleComponent := componentIns.(*components.Console)
			consoleComponent.RunAction()
		}
	}

	if !isCallConsole {
		fmt.Println("the console component is not enabled.")
	}
}

// Overwrite: 实现获取组件实例的方法
func (app *application) GetComponent(v base.ComponentInterface) base.ComponentInterface {
	var componentIns base.ComponentInterface = nil

	targetName := utils.Reflect.GetClassName(v)
	for _, ins := range app.componentDict {
		if utils.Reflect.GetClassName(ins) == targetName {
			componentIns = ins
			break
		}
	}

	return componentIns
}

// Overwrite: get component instance by name
func (app *application) GetComponentByName(name string) base.ComponentInterface {
	componentIns, has := app.componentDict[name]
	if !has {
		return nil
	}
	return componentIns
}

// get default db component's interface
func (app *application) GetDBInterface() base.ComponentInterface {
	componentIns := app.GetComponentByName(app.config.AppConfig.DefaultDBName)
	if componentIns == nil {
		return nil
	}
	return componentIns
}

// get default db component
func (app *application) GetDB() *components.Database {
	return app.GetDBInterface().(*components.Database)
}

// TODO
// add migration
func (app *application) AddMigration(m base.MigrationInterface) {

}

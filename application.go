package cam

import (
	"fmt"
	"github.com/go-cam/cam/core/base"
	"github.com/go-cam/cam/core/components"
	"github.com/go-cam/cam/core/configs"
	"github.com/go-cam/cam/core/utils"
	"os"
	"reflect"
	"strconv"
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

	// migrations's struct dict
	migrationDict map[string]base.MigrationInterface

	// log component
	logComponent *components.Log
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
	app.migrationDict = map[string]base.MigrationInterface{}
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

	// read config component
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

	// init core component
	app.initCoreComponent()
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
	if _, has := t.FieldByName("PluginMigrate"); has {
		pluginMigrate := v.FieldByName("PluginMigrate").Interface().(configs.PluginMigrate)
		pluginMigrate.MigrationDict = app.migrationDict
		v.FieldByName("PluginMigrate").Set(reflect.ValueOf(pluginMigrate))
	}
}

// init core component
func (app *application) initCoreComponent() {

	app.initCoreComponentLog()
}

// init LogComponent. if LogComponent not in the dict, create one
func (app *application) initCoreComponentLog() {
	logComponent, _ := app.getComponentAndName(new(components.Log))
	if logComponent != nil {
		app.logComponent = logComponent.(*components.Log)
	} else {
		var name = "log"
		var has = true
		for i := 0; !has; i++ {
			if i != 0 {
				name = "log" + strconv.Itoa(i)
			}
			_, has = app.componentDict[name]
		}

		logConfig := new(configs.Log)
		logComponent = new(components.Log)
		logConfig.Component = logComponent
		logComponent.Init(logConfig)
		app.logComponent = logComponent.(*components.Log)
		app.componentDict[name] = logComponent
	}
}

// 调用命令行组件
func (app *application) callConsole() {
	isCallConsole := false

	for _, componentIns := range app.componentDict {
		name := utils.Reflect.GetStructName(componentIns)
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

// get component and the name in the dict
func (app *application) getComponentAndName(v base.ComponentInterface) (base.ComponentInterface, string) {
	var componentIns base.ComponentInterface = nil
	var componentName = ""

	targetName := utils.Reflect.GetStructName(v)
	for name, ins := range app.componentDict {
		if utils.Reflect.GetStructName(ins) == targetName {
			componentIns = ins
			componentName = name
			break
		}
	}

	return componentIns, componentName
}

// Overwrite: 实现获取组件实例的方法
func (app *application) GetComponent(v base.ComponentInterface) base.ComponentInterface {
	ins, _ := app.getComponentAndName(v)
	return ins
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
	ins := app.GetDBInterface()
	var db *components.Database = nil
	if ins != nil {
		db = ins.(*components.Database)
	}
	return db
}

// add migration struct
func (app *application) AddMigration(m base.MigrationInterface) {
	id := utils.Reflect.GetStructName(m)
	app.migrationDict[id] = m
}

// log info
func (app *application) Info(title string, content string) error {
	return app.logComponent.Info(title, content)
}

// log warning
func (app *application) Warn(title string, content string) error {
	return app.logComponent.Warn(title, content)
}

// log error
func (app *application) Error(title string, content string) error {
	return app.logComponent.Error(title, content)
}

// get one .evn file values
func (app *application) GetEvn(key string) string {
	return utils.Env.Get(key)
}

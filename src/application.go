package cin

import (
	"cin/src/base"
	"cin/src/models"
	"reflect"
)

// 服务器全局类
type Application struct {
	status        base.ApplicationStatus // 应用状态（初始化、开始、运行中、停止、销毁）[onInit, onStart, onRun, onStop, onDestroy]
	config        *Config                // 应用全部配置
	componentDict map[string]base.ComponentInterface
}

// 应用全局实例（只需要一个实例即可操作整个应用）
var App = NewApplication()

// 框架方法
func NewApplication() *Application {
	app := new(Application)
	app.status = ApplicationStatusInit
	app.config = new(Config)
	app.config.Config = *models.NewConfig()
	app.componentDict = map[string]base.ComponentInterface{}
	return app
}

// 添加配置（有先后顺序。后面的配置将覆盖前面的配置）
func (app *Application) AddConfig(config *Config) {
	for key, value := range config.Params {
		app.config.Params[key] = value
	}
	for name, componentConfig := range config.ComponentDict {
		app.config.ComponentDict[name] = componentConfig
	}
}

// 启动应用
func (app *Application) Run() {
	app.onInit()
	app.onStart()
	app.onRun()
	//app.onStop()
	//app.onDestroy()
}

// TODO
func (app *Application) onInit() {
	// 实例化组件
	for name, config := range app.config.ComponentDict {
		// 使用结构体重新创建一个新对象（防止外部修改导致混乱）
		componentInterface := config.GetComponent()
		t := reflect.TypeOf(componentInterface)
		componentType := t.Elem()
		componentValue := reflect.New(componentType)
		componentInterface = componentValue.Interface().(base.ComponentInterface)

		componentInterface.Init(config)
		app.componentDict[name] = componentInterface
	}
}

// TODO
func (app *Application) onStart() {

}

// TODO
func (app *Application) onRun() {

}

// TODO
func (app *Application) onStop() {

}

// TODO
func (app *Application) onDestroy() {

}

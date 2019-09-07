package cin

// 服务器全局类
type Application struct {
	status int8   // 应用状态（初始化、开始、运行中、停止、销毁）[onInit, onStart, onRun, onStop, onDestroy]
	config Config //
}

// 应用全局实例（只需要一个实例即可操作整个应用）
var App = NewApplication()

// 框架方法
func NewApplication() *Application {
	app := new(Application)
	return app
}

// 添加配置（有先后顺序。后面的配置将覆盖前面的配置）
// TODO
func (app *Application) AddConfig(config *Config) {

}

// 启动应用
func (app *Application) Run() {
	app.onInit()
	app.onStart()
	app.onRun()
	app.onStop()
	app.onDestroy()
}

// TODO
func (app *Application) onInit() {

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

package cam

import (
	"fmt"
	"github.com/go-cam/cam/core/camBase"
	"github.com/go-cam/cam/core/camComponents"
	"github.com/go-cam/cam/core/camConfigs"
	"github.com/go-cam/cam/core/camConstants"
	"github.com/go-cam/cam/core/camUtils"
	"os"
	"reflect"
	"strconv"
	"time"
)

// framework application global instance struct define
type application struct {
	camBase.ApplicationInterface

	status        camBase.ApplicationStatus // application status[onInit, onStart, onRun, onStop, onDestroy]
	config        *Config                   // application config
	router        *router
	logComponent  *camComponents.Log                    // log component
	componentDict map[string]camBase.ComponentInterface // components dict
	migrationDict map[string]camBase.MigrationInterface // migrations's struct dict
}

// single instance
var App = newApplication()

// new application instance
func newApplication() *application {
	app := new(application)
	app.status = camConstants.ApplicationStatusInit
	app.config = NewConfig()
	app.config.AppConfig = NewAppConfig()
	app.router = newRouter()
	app.componentDict = map[string]camBase.ComponentInterface{}
	app.migrationDict = map[string]camBase.MigrationInterface{}
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

// get router
func (app *application) GetRouter() *router {
	return app.router
}

// run application
func (app *application) Run() {
	fmt.Println("App: Initializing ...")
	app.onInit()
	if len(os.Args) >= 2 {
		// It's a console command
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

// init application and components
func (app *application) onInit() {
	camComponents.SetApplication(app)

	// read config component
	for name, config := range app.config.ComponentDict {
		componentInterface := config.GetComponent()
		t := reflect.TypeOf(componentInterface)
		componentType := t.Elem()
		componentValue := reflect.New(componentType)
		componentInterface = componentValue.Interface().(camBase.ComponentInterface)

		// input plugin params
		app.writePluginParams(config)

		componentInterface.Init(config)
		componentInterface.SetApp(app)
		app.componentDict[name] = componentInterface
	}

	// init core component
	app.initCoreComponent()
}

// startup all components
func (app *application) onStart() {
	for _, component := range app.componentDict {
		go component.Start()
	}
}

// stop all components
func (app *application) onStop() {
	for _, component := range app.componentDict {
		component.Stop()
	}
}

// Wait until the app call Stop()
func (app *application) wait() {
	for {
		time.Sleep(1 * time.Second)
	}
}

// input plugin params
func (app *application) writePluginParams(config camBase.ConfigComponentInterface) {
	t := reflect.TypeOf(config).Elem()
	v := reflect.ValueOf(config).Elem()
	// router plugin
	if _, has := t.FieldByName("PluginRouter"); has {
		pluginRouter := v.FieldByName("PluginRouter").Interface().(camConfigs.PluginRouter)
		pluginRouter.ControllerList = app.router.controllerList
		pluginRouter.ConsoleControllerList = app.router.consoleControllerList
		pluginRouter.OnWebsocketMessageHandler = app.router.onWebsocketMessageHandler
		v.FieldByName("PluginRouter").Set(reflect.ValueOf(pluginRouter))
	}
	// migrate plugin
	if _, has := t.FieldByName("PluginMigrate"); has {
		pluginMigrate := v.FieldByName("PluginMigrate").Interface().(camConfigs.PluginMigrate)
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
	logComponent, _ := app.getComponentAndName(new(camComponents.Log))
	if logComponent != nil {
		app.logComponent = logComponent.(*camComponents.Log)
	} else {
		var name = "log"
		var has = true
		for i := 0; !has; i++ {
			if i != 0 {
				name = "log" + strconv.Itoa(i)
			}
			_, has = app.componentDict[name]
		}

		logConfig := new(camConfigs.Log)
		logComponent = new(camComponents.Log)
		logConfig.Component = logComponent
		logComponent.Init(logConfig)
		app.logComponent = logComponent.(*camComponents.Log)
		app.componentDict[name] = logComponent
	}
}

// Call console
func (app *application) callConsole() {
	isCallConsole := false

	for _, componentIns := range app.componentDict {
		name := camUtils.Reflect.GetStructName(componentIns)
		if name == "Console" {
			isCallConsole = true
			consoleComponent := componentIns.(*camComponents.Console)
			consoleComponent.RunAction()
		}
	}

	if !isCallConsole {
		fmt.Println("the console component is not enabled.")
	}
}

// get component and the name in the dict
func (app *application) getComponentAndName(v camBase.ComponentInterface) (camBase.ComponentInterface, string) {
	var componentIns camBase.ComponentInterface = nil
	var componentName = ""

	targetName := camUtils.Reflect.GetStructName(v)
	for name, ins := range app.componentDict {
		if camUtils.Reflect.GetStructName(ins) == targetName {
			componentIns = ins
			componentName = name
			break
		}
	}

	return componentIns, componentName
}

// Overwrite:
// Try to get instance using struct type
func (app *application) GetComponent(v camBase.ComponentInterface) camBase.ComponentInterface {
	ins, _ := app.getComponentAndName(v)
	return ins
}

// Overwrite:
// Try to get component instance by name.
// The name is define in config
func (app *application) GetComponentByName(name string) camBase.ComponentInterface {
	componentIns, has := app.componentDict[name]
	if !has {
		return nil
	}
	return componentIns
}

// get default db component's interface
func (app *application) GetDBInterface() camBase.ComponentInterface {
	componentIns := app.GetComponentByName(app.config.AppConfig.DefaultDBName)
	if componentIns == nil {
		return nil
	}
	return componentIns
}

// get default db component
func (app *application) GetDB() *camComponents.Database {
	ins := app.GetDBInterface()
	var db *camComponents.Database = nil
	if ins != nil {
		db = ins.(*camComponents.Database)
	}
	return db
}

// add migration struct
func (app *application) AddMigration(m camBase.MigrationInterface) {
	id := camUtils.Reflect.GetStructName(m)
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
	return camUtils.Env.Get(key)
}

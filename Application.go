package cam

import (
	"github.com/go-cam/cam/camAppConfig"
	"github.com/go-cam/cam/camBase"
	"github.com/go-cam/cam/camConsole"
	"github.com/go-cam/cam/camConstants"
	"github.com/go-cam/cam/camLog"
	"github.com/go-cam/cam/camUtils"
	"reflect"
	"strconv"
	"time"
)

// framework Application global instance struct define
type Application struct {
	camBase.ApplicationInterface

	status        camBase.ApplicationStatus             // Application status[onInit, onStart, onRun, onStop, onDestroy]
	config        *camAppConfig.Config                  // Application config
	logComponent  *camLog.LogComponent                  // log component
	waitChan      chan bool                             // wait until call Application.Stop()'s sign
	componentDict map[string]camBase.ComponentInterface // components dict
	migrationDict map[string]camBase.MigrationInterface // migration map
}

var App camBase.ApplicationInterface

func init() {
	camBase.App = NewApplication()
	App = camBase.App
}

// new Application instance
func NewApplication() *Application {
	app := new(Application)
	app.status = camConstants.ApplicationStatusInit
	app.config = NewConfig()
	app.config.AppConfig = NewAppConfig()
	app.waitChan = make(chan bool)
	app.componentDict = map[string]camBase.ComponentInterface{}
	app.migrationDict = map[string]camBase.MigrationInterface{}
	return app
}

// Add config
// Merge as much as possible, otherwise overwrite.
//
// config: new config
func (app *Application) AddConfig(configI camBase.AppConfigInterface) {
	config, ok := configI.(*camAppConfig.Config)
	if !ok {
		panic("Wrong type. need: *camModels.Config")
	}

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

// run Application
func (app *Application) Run() {
	if camUtils.Console.IsRunByCommand() {
		app.onInit()
		app.callConsole()
	} else {
		app.onInit()
		app.onStart()
		select {
		case <-app.waitChan:
			app.onStop()
		}
	}
}

// init Application and components
func (app *Application) onInit() {
	// read config component
	for name, config := range app.config.ComponentDict {
		componentInterface := config.GetComponent()
		t := reflect.TypeOf(componentInterface)
		componentType := t.Elem()
		componentValue := reflect.New(componentType)
		componentInterface = componentValue.Interface().(camBase.ComponentInterface)
		componentInterface.Init(config)
		componentInterface.SetApp(app)
		app.componentDict[name] = componentInterface
	}
	// init core component
	app.initCoreComponent()
}

// startup all components
func (app *Application) onStart() {
	for name, component := range app.componentDict {
		go component.Start()
		app.Info("runtime", "start component:"+name)
	}
	app.Info("runtime", "Application start finished.")
}

// stop all components
func (app *Application) onStop() {
	for name, component := range app.componentDict {
		component.Stop()
		app.Info("runtime", "stop component:"+name)
	}
	app.Info("runtime", "Application stop finished.")
}

// Wait until the app call Stop()
func (app *Application) wait() {
	for {
		time.Sleep(1 * time.Second)
	}
}

// init core component
func (app *Application) initCoreComponent() {
	app.initCoreComponentLog()
}

// init LogComponent. if LogComponent not in the dict, create one
func (app *Application) initCoreComponentLog() {
	logComponent, _ := app.getComponentAndName(new(camLog.LogComponent))
	if logComponent != nil {
		app.logComponent = logComponent.(*camLog.LogComponent)
	} else {
		var name = "log"
		var has = true
		for i := 0; !has; i++ {
			if i != 0 {
				name = "log" + strconv.Itoa(i)
			}
			_, has = app.componentDict[name]
		}

		logConfig := camLog.NewLogConfig()
		logComponent = new(camLog.LogComponent)
		logConfig.Component = logComponent
		logComponent.Init(logConfig)
		app.logComponent = logComponent.(*camLog.LogComponent)
		app.componentDict[name] = logComponent
	}
}

// Call console
func (app *Application) callConsole() {
	isCallConsole := false

	for _, componentIns := range app.componentDict {
		name := camUtils.Reflect.GetStructName(componentIns)
		if name == "ConsoleComponent" {
			isCallConsole = true
			consoleComponent := componentIns.(*camConsole.ConsoleComponent)
			consoleComponent.RunAction()
		}
	}

	if !isCallConsole {
		app.Info("runtime-console", "the console component is not enabled.")
	}
}

// get component and the name in the dict
func (app *Application) getComponentAndName(v camBase.ComponentInterface) (camBase.ComponentInterface, string) {
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
func (app *Application) GetComponent(v camBase.ComponentInterface) camBase.ComponentInterface {
	ins, _ := app.getComponentAndName(v)
	return ins
}

// Overwrite:
// Try to get component instance by name.
// The name is define in config
func (app *Application) GetComponentByName(name string) camBase.ComponentInterface {
	componentIns, has := app.componentDict[name]
	if !has {
		return nil
	}
	return componentIns
}

// get default db component
func (app *Application) GetDB() camBase.DatabaseComponentInterface {
	dbCompI := app.GetComponentByName(app.config.AppConfig.DefaultDBName)
	if dbCompI == nil {
		return nil
	}

	dbComp, ok := dbCompI.(camBase.DatabaseComponentInterface)
	if !ok {
		return nil
	}

	return dbComp
}

// add migration struct
func (app *Application) AddMigration(m camBase.MigrationInterface) {
	id := camUtils.Reflect.GetStructName(m)
	app.migrationDict[id] = m
}

// log debug
func (app *Application) Debug(title string, content string) {
	err := app.logComponent.Debug(title, content)
	if err != nil {
		panic(err)
	}
}

// log info
func (app *Application) Info(title string, content string) {
	err := app.logComponent.Info(title, content)
	if err != nil {
		panic(err)
	}
}

// log warning
func (app *Application) Warn(title string, content string) {
	err := app.logComponent.Warn(title, content)
	if err != nil {
		panic(err)
	}
}

// log error
func (app *Application) Error(title string, content string) {
	err := app.logComponent.Error(title, content)
	if err != nil {
		panic(err)
	}
}

// get one .evn file values
func (app *Application) GetEvn(key string) string {
	return camUtils.Env.Get(key)
}

// stop Application
func (app *Application) Stop() {
	app.waitChan <- true
}

func (app *Application) GetMigrateDict() map[string]camBase.MigrationInterface {
	return app.migrationDict
}

// get value form app.config.Params.
func (app *Application) GetParam(key string) interface{} {
	i, has := app.config.Params[key]
	if !has {
		return nil
	}
	return i
}

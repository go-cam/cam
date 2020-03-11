package cam

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camConfig"
	"github.com/go-cam/cam/base/camConstants"
	"github.com/go-cam/cam/base/camUtils"
	"github.com/go-cam/cam/component/camCache"
	"github.com/go-cam/cam/component/camConsole"
	"github.com/go-cam/cam/component/camDatabase"
	"github.com/go-cam/cam/component/camLog"
	"github.com/go-cam/cam/component/camMail"
	"strconv"
	"time"
)

// framework Application global instance struct define
type Application struct {
	camBase.ApplicationInterface

	status        camBase.ApplicationStatus             // Application status[onInit, onStart, onRun, onStop, onDestroy]
	config        *camConfig.Config                     // Application config
	logComponent  *camLog.LogComponent                  // log component
	cache         camBase.CacheComponentInterface       // cache component
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
	app.status = camConstants.AppStatusBeforeInit
	app.config = NewConfig()
	app.config.AppConfig = NewAppConfig()
	app.cache = nil
	app.componentDict = map[string]camBase.ComponentInterface{}
	app.migrationDict = map[string]camBase.MigrationInterface{}
	return app
}

// Add config. Must be called before calling cam.App.Run ().
// Merge as much as possible, otherwise overwrite.
//
// config: new config
func (app *Application) AddConfig(configI camBase.AppConfigInterface) {
	config, ok := configI.(*camConfig.Config)
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
		app.wait()
	}
}

// init Application and components
func (app *Application) onInit() {
	// read config component
	for name, config := range app.config.ComponentDict {
		componentI := config.NewComponent()
		componentI.Init(config)
		app.componentDict[name] = componentI
	}
	// init core component
	app.initCoreComponent()

	app.status = camConstants.AppStatusBeforeStart
}

// startup all components
func (app *Application) onStart() {
	for name, component := range app.componentDict {
		go component.Start()
		app.Trace("Application.onStart", "start component:"+name)
	}
	app.Trace("Application.onStart", "Application start finished.")

	app.status = camConstants.AppStatusAfterStart
}

// stop all components
func (app *Application) onStop() {
	for name, component := range app.componentDict {
		component.Stop()
		delete(app.componentDict, name)
		app.Trace("Application.onStop", "stop component:"+name)
	}
	app.Trace("Application.onStop", "Application stop finished.")

	app.status = camConstants.AppStatusAfterStop
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
		app.Warn("Application.callConsole", "the console component is not enabled.")
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
		dbCompI = app.GetComponent(&camDatabase.DatabaseComponent{})
		if dbCompI == nil {
			return nil
		}
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

// base log
func (app *Application) basicLog(logLevel camBase.LogLevel, title string, content string) {
	err := app.logComponent.Record(logLevel, title, content)
	if err != nil {
		panic(err)
	}
}

// log trace
func (app *Application) Trace(title string, content string) {
	app.basicLog(LogLevelTrace, title, content)
}

// log debug
func (app *Application) Debug(title string, content string) {
	app.basicLog(LogLevelDebug, title, content)
}

// log info
func (app *Application) Info(title string, content string) {
	app.basicLog(LogLevelInfo, title, content)
}

// log warning
func (app *Application) Warn(title string, content string) {
	app.basicLog(LogLevelWarn, title, content)
}

// log error
func (app *Application) Error(title string, content string) {
	app.basicLog(LogLevelError, title, content)
}

// log fatal
func (app *Application) Fatal(title string, content string) {
	app.basicLog(LogLevelFatal, title, content)
}

// get one .evn file values
func (app *Application) GetEvn(key string) string {
	return camUtils.Env.Get(key)
}

// stop Application
func (app *Application) Stop() {
	app.status = camConstants.AppStatusBeforeStop
	app.onStop()
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

// get cache component
func (app *Application) GetCache() camBase.CacheComponentInterface {
	if app.cache != nil {
		return app.cache
	}

	var ok bool
	compI := app.GetComponent(&camCache.CacheComponent{})
	if compI == nil {
		compI = app.createDefaultCacheComponent()
		app.cache, ok = compI.(camBase.CacheComponentInterface)
		if !ok {
			app.Error("Application.GetCache", "create default cache fail")
		}
	} else {
		app.cache, ok = compI.(camBase.CacheComponentInterface)
		if !ok {
			app.Error("Application.GetCache", "convert fail")
			return nil
		}
	}

	return app.cache
}

// create default cache component
func (app *Application) createDefaultCacheComponent() camBase.ComponentInterface {
	cacheConfig := NewCacheConfig()
	componentI := cacheConfig.NewComponent()
	componentName := "cache"

	i := 0
	for {
		tmpComponent := app.GetComponentByName(componentName)
		if tmpComponent == nil {
			break
		}
		componentName = "cache_" + strconv.Itoa(i)
		i++
	}
	if app.status >= camConstants.AppStatusBeforeStart {
		componentI.Init(cacheConfig)
	}
	if app.status >= camConstants.AppStatusAfterStart {
		componentI.Start()
	}

	app.componentDict[componentName] = componentI
	return componentI
}

// get mail component
func (app *Application) GetMail() camBase.MailComponentInterface {
	compI := app.GetComponent(&camMail.MailComponent{})
	if compI == nil {
		return nil
	}
	mailCompI, ok := compI.(camBase.MailComponentInterface)
	if !ok {
		return nil
	}
	return mailCompI
}

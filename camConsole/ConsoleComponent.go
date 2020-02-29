package camConsole

import (
	"fmt"
	"github.com/go-cam/cam/camBase"
	"github.com/go-cam/cam/camPluginRouter"
	"github.com/go-cam/cam/camUtils"
	"os"
	"reflect"
	"xorm.io/xorm"
)

// command component
type ConsoleComponent struct {
	camBase.Component
	camPluginRouter.RouterPlugin

	config *ConsoleComponentConfig
}

// init
func (component *ConsoleComponent) Init(configInterface camBase.ComponentConfigInterface) {
	component.Component.Init(configInterface)

	configValue := reflect.ValueOf(configInterface)
	var config *ConsoleComponentConfig
	if configValue.Kind() == reflect.Ptr {
		config = configValue.Interface().(*ConsoleComponentConfig)
	} else if configValue.Kind() == reflect.Struct {
		configStruct := configValue.Interface().(ConsoleComponentConfig)
		config = &configStruct
	} else {
		panic("illegal config")
	}
	component.config = config

	// register controller
	//component.controllerDict, component.controllerActionDict = camComponents.Common.GetControllerDict(config.ControllerList)

	// init router plugin
	component.RouterPlugin.Init(&config.RouterPluginConfig)
}

// run command
// Example:
//	# go build cam.go
//	# ./cam controllerName/actionName param1 param2
func (component *ConsoleComponent) RunAction() {
	if len(os.Args) < 2 {
		fmt.Println("please input route")
		return
	}

	route := os.Args[1]
	controller, action := component.GetControllerAction(route)
	if controller == nil || action == nil {
		panic("route not found: " + route)
	}
	controller.Init()
	controller.SetApp(component.App)
	if !controller.BeforeAction(action) {
		panic("invalid call")
		return
	}
	action.Call()
	response := controller.AfterAction(action, controller.GetResponse())
	fmt.Println(string(response))
}

// get migrate up version list.
func (component *ConsoleComponent) GetMigrateUpVersionList() []string {
	lastVersion := component.MigrateLastVersion()
	var versionList []string
	for version, _ := range camBase.App.GetMigrateDict() {
		if version <= lastVersion {
			continue
		}
		versionList = append(versionList, version)
	}
	return versionList
}

func (component *ConsoleComponent) MigrateLastVersion() string {
	session := component.getDBSession()
	exists, err := session.IsTableExist(new(Migration))
	camUtils.Error.Panic(err)
	if !exists {
		err = component.createMigrateVersionTable()
		camUtils.Error.Panic(err)
		return ""
	}

	migration := new(Migration)
	has, err := session.Desc("version").Get(migration)
	camUtils.Error.Panic(err)

	version := ""
	if has {
		version = migration.Version
	}
	return version
}

// create migrations's version record table
func (component *ConsoleComponent) createMigrateVersionTable() error {
	session := component.getDBSession()
	migration := new(Migration)
	return session.Sync2(migration)
}

// up all database version
func (component *ConsoleComponent) MigrateUp() {
	fmt.Println("Migrate up start.")

	lastVersion := component.MigrateLastVersion()
	var err error
	for version, m := range camBase.App.GetMigrateDict() {
		if version <= lastVersion {
			continue
		}

		fmt.Print("\tup version: " + version + " ...")

		m.Up()
		sqlList := m.GetSqlList()

		session := component.getDBSession()
		err = session.Begin()
		camUtils.Error.Panic(err)

		for _, sqlStr := range sqlList {
			_, err = session.Exec(sqlStr)
			if err != nil {
				panic(err)
			}
		}
		migration := new(Migration)
		migration.Version = version
		_, err := session.Insert(migration)
		if err != nil {
			_ = session.Rollback()
			camUtils.Error.Panic(err)
		} else {
			err = session.Commit()
			camUtils.Error.Panic(err)
		}

		fmt.Println(" done.")
	}
	fmt.Println("Migrate up finished.")
}

// down last database version
func (component *ConsoleComponent) MigrateDown() {
	lastVersion := component.MigrateLastVersion()
	m, has := camBase.App.GetMigrateDict()[lastVersion]
	if !has {
		if lastVersion == "" {
			fmt.Println("no version can be down.")
		} else {
			fmt.Println("not found " + lastVersion + "' struct")
		}
		return
	}
	fmt.Println("version: " + lastVersion)
	fmt.Print("Do you want to down this version ?[Y/N]:")
	if !camUtils.Console.IsPressY() {
		return
	}

	m.Down()
	sqlList := m.GetSqlList()

	var err error
	session := component.getDBSession()
	err = session.Begin()
	camUtils.Error.Panic(err)
	defer func() {
		if rec := recover(); rec != nil {
			_ = session.Rollback()
		} else {
			_ = session.Commit()
		}
	}()

	for _, sqlStr := range sqlList {
		_, err = session.Exec(sqlStr)
		if err != nil {
			panic(err)
		}
	}

	_, err = session.ID(lastVersion).Delete(Migration{})
	camUtils.Error.Panic(err)

	fmt.Println(" done.")
}

// get database session
func (component *ConsoleComponent) getDBSession() *xorm.Session {
	db := camBase.App.GetDB()
	if db == nil {
		panic("no database")
	}
	return db.NewSession()
}

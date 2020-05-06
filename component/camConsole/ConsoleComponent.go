package camConsole

import (
	"fmt"
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camUtils"
	"github.com/go-cam/cam/component"
	"github.com/go-cam/cam/plugin/camContext"
	"github.com/go-cam/cam/plugin/camRouter"
	"os"
	"sort"
	"xorm.io/xorm"
)

// command component
type ConsoleComponent struct {
	component.Component
	camRouter.RouterPlugin
	camContext.ContextPlugin

	config *ConsoleComponentConfig
}

// init
func (comp *ConsoleComponent) Init(configI camBase.ComponentConfigInterface) {
	comp.Component.Init(configI)

	var ok bool
	comp.config, ok = configI.(*ConsoleComponentConfig)
	if !ok {
		camBase.App.Fatal("ConsoleComponent", "invalid config")
		return
	}

	// init router plugin
	comp.RouterPlugin.Init(&comp.config.RouterPluginConfig)
	comp.ContextPlugin.Init(&comp.config.ContextPluginConfig)
}

// run command
// Example:
//	# go build cam.go
//	# ./cam controllerName/actionName param1 param2
func (comp *ConsoleComponent) RunAction() {
	if len(os.Args) < 2 {
		fmt.Println("please input route")
		return
	}

	route := os.Args[1]
	controller, action := comp.GetControllerAction(route)
	if controller == nil || action == nil {
		panic("route not found: " + route)
	}
	ctx := comp.NewContext()
	controller.Init()
	controller.SetContext(ctx)
	if !controller.BeforeAction(action) {
		panic("invalid call")
		return
	}
	action.Call()
	response := controller.AfterAction(action, controller.GetContext().Read())
	fmt.Println(string(response))
}

// get migrate up version list.
func (comp *ConsoleComponent) GetMigrateUpVersionList() []string {
	lastVersion := comp.MigrateLastVersion()
	var versionList []string
	for version := range camBase.App.GetMigrateDict() {
		if version <= lastVersion {
			continue
		}
		versionList = append(versionList, version)
	}
	sort.Slice(versionList, func(i, j int) bool {
		return versionList[i] < versionList[j]
	})
	return versionList
}

func (comp *ConsoleComponent) MigrateLastVersion() string {
	session := comp.getDBSession()
	exists, err := session.IsTableExist(new(Migration))
	if err != nil {
		panic(err)
	}
	if !exists {
		err = comp.createMigrateVersionTable()
		if err != nil {
			panic(err)
		}
		return ""
	}

	migration := new(Migration)
	has, err := session.Desc("version").Get(migration)
	if err != nil {
		panic(err)
	}

	version := ""
	if has {
		version = migration.Version
	}
	return version
}

// create migrations's version record table
func (comp *ConsoleComponent) createMigrateVersionTable() error {
	session := comp.getDBSession()
	migration := new(Migration)
	return session.Sync2(migration)
}

// up all database version
func (comp *ConsoleComponent) MigrateUp() {
	fmt.Println("Migrate up start.")

	lastVersion := comp.MigrateLastVersion()
	var err error
	for version, m := range camBase.App.GetMigrateDict() {
		if version <= lastVersion {
			continue
		}

		fmt.Print("\tup version: " + version + " ...")

		m.Up()
		sqlList := m.GetSqlList()

		session := comp.getDBSession()
		err = session.Begin()
		if err != nil {
			panic(err)
		}

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
			panic(err)
		} else {
			err = session.Commit()
			if err != nil {
				panic(err)
			}
		}

		fmt.Println(" done.")
	}
	fmt.Println("Migrate up finished.")
}

// down last database version
func (comp *ConsoleComponent) MigrateDown() {
	lastVersion := comp.MigrateLastVersion()
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
	session := comp.getDBSession()
	err = session.Begin()
	if err != nil {
		panic(err)
	}
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
	if err != nil {
		panic(err)
	}

	fmt.Println(" done.")
}

// get database session
func (comp *ConsoleComponent) getDBSession() *xorm.Session {
	db := camBase.App.GetDB()
	if db == nil {
		panic("no database")
	}
	return db.NewSession()
}

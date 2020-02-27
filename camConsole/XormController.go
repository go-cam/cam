package camConsole

import (
	"fmt"
	"github.com/go-cam/cam/camUtils"
)

// xorm's console controller
type XormController struct {
	ConsoleController
}

// install github.com/go-xorm/cmd/xorm
func (controller *XormController) InstallCmdXorm() {
	_ = camUtils.Console.Start("go get github.com/go-xorm/cmd/xorm")
}

// Generate ORM classes automatically according to the database
// xorm reverse
//		usage: xorm reverse [-s] driverName datasourceName tmplPath [generatedPath] [tableFilterReg]
//
//		according database's tables and columns to generate codes for Go, C++ and etc.
//
//		-s                Generated one go file for every table
//		driverName        Database driver name, now supported four: mysql mymysql sqlite3 postgres
//		datasourceName    Database connection uri, for detail infomation please visit driver's project page
//		tmplPath          Template dir for generated. the default templates dir has provide 1 template
//		generatedPath     This parameter is optional, if blank, the default value is model, then will
//		generated all codes in model dir
//		tableFilterReg    Table name filter regexp
func (controller *XormController) Generate() {
	db := controller.GetDatabaseComponent()
	dsn := db.GetDSN()
	tplDir := db.GetXormTemplateDir()
	modelsDir := db.GetXormModelDir()
	if !camUtils.File.Exists(modelsDir) {
		err := camUtils.File.Mkdir(modelsDir)
		camUtils.Error.Panic(err)
	}

	command := "xorm reverse mysql \"" + dsn + "\" \"" + tplDir + "\" \"" + modelsDir + "\""
	fmt.Println(command)
	err := camUtils.Console.Start(command)
	camUtils.Error.Panic(err)
}

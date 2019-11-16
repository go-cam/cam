package controllers

import (
	"fmt"
	"github.com/cinling/cin/core/utils"
)

// xorm's console controller
type XormController struct {
	baseConsoleController
}

// install github.com/go-xorm/cmd/xorm
func (controller *XormController) InstallCmdXorm() {
	_ = utils.Console.Start("go get github.com/go-xorm/cmd/xorm")
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
	tmlDir := db.GetXormTemplateDir()
	modelsDir := db.GetXormModelDir()
	if !utils.File.Exists(modelsDir) {
		err := utils.File.Mkdir(modelsDir)
		utils.Error.Panic(err)
	}

	command := "xorm reverse mysql \"" + dsn + "\" \"" + tmlDir + "\" \"" + modelsDir + "\""
	fmt.Println(command)
	err := utils.Console.Start(command)
	utils.Error.Panic(err)
}

package controllers

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/go-cam/cam/core/models/tpls"
	"github.com/go-cam/cam/core/utils"
	"html/template"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 数据库版本管理控制器
type MigrateController struct {
	baseConsoleController
}

// 创建一个数据库升级文件
func (controller *MigrateController) CreateBak() {
	var err error

	// 生成路径
	migrateDir := controller.GetValue("migrateDir").(string)
	if !utils.File.Exists(migrateDir) {
		err = utils.File.Mkdir(migrateDir)
		utils.Error.Panic(err)
	}

	// 生成id
	timestamp := time.Now().Unix()
	timestampStr := strconv.FormatInt(timestamp, 10)

	// 名字
	name := "new_migrate"
	if len(os.Args) >= 3 {
		name = os.Args[2]

	}
	upFilename := migrateDir + "/" + timestampStr + "_" + name + ".up.sql"
	downFilename := migrateDir + "/" + timestampStr + "_" + name + ".down.sql"

	fmt.Println("General filename...")
	fmt.Println("\t" + downFilename)
	fmt.Println("\t" + upFilename)
	fmt.Print("Do you want to create the following two file?[Y/N]:")
	input := bufio.NewScanner(os.Stdin)
	if !input.Scan() {
		return
	}
	str := strings.ToLower(input.Text())
	if str != "y" {
		return
	}

	if !utils.File.Exists(migrateDir) {
		err = utils.File.Mkdir(migrateDir)
		utils.Error.Panic(err)
	}

	err = utils.File.WriteFile(downFilename, []byte{})
	utils.Error.Panic(err)
	err = utils.File.WriteFile(upFilename, []byte{})
	utils.Error.Panic(err)

	fmt.Println("")
	fmt.Println("Done: migrations's files created.")
}

func (controller *MigrateController) Create() {
	name := controller.GetArgv(0)
	if len(name) == 0 {
		name = "new_migrate"
	}
	if !controller.validName(name) {
		fmt.Println("Illegal name: " + name + ". Chars must in [a-z][A-Z][0-9]_")
		return
	}

	var err error

	filename := controller.getFilename(name)
	fmt.Println("General filename...")
	fmt.Println("\t" + filename)
	fmt.Print("Do you want to create the following two file?[Y/N]:")
	input := bufio.NewScanner(os.Stdin)
	if !input.Scan() {
		return
	}
	str := strings.ToLower(input.Text())
	if str != "y" {
		return
	}

	db := controller.GetDatabaseComponent()
	migrateDir := db.GetMigrateDir()
	if !utils.File.Exists(migrateDir) {
		err = utils.File.Mkdir(migrateDir)
		utils.Error.Panic(err)
	}
	absFilename := migrateDir + "/" + filename
	content := controller.getMigrationContent(filename)
	err = utils.File.WriteFile(absFilename, content)
	utils.Error.Panic(err)

	fmt.Print("\nDone.")
}

// migrate up
func (controller *MigrateController) Up() {
	db := controller.GetDatabaseComponent()
	if db == nil {
		panic("no database.")
	}
	db.MigrateUp()
}

// migrate down
func (controller *MigrateController) Down() {
	db := controller.GetDatabaseComponent()
	if db == nil {
		panic("no database.")
	}
	db.MigrateDown()
}

// get migration's filename. only filename, not absolute filename
func (controller *MigrateController) getFilename(name string) string {
	// generate filename
	id := utils.Migrate.IdByDatetime()
	return "m" + id + "_" + name + ".go"
}

// orm struct's file template
func (controller *MigrateController) getTpl() string {
	return `package migrations

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/core/models"
)

func init() {
	m := new(m191118_073955_init)
	cam.App.AddMigration(m)
}

type {{ .ClassName}} struct {
	models.Migration
}

// up
func (migration *{{ .ClassName}}) Up() {
}

// down
func (migration *{{ .ClassName}}) Down() {
}
`
}

// get migration file content
func (controller *MigrateController) getMigrationContent(filename string) []byte {
	var err error
	var retBytes = []byte("")

	t, err := template.New(filename).Parse(controller.getTpl())
	utils.Error.Panic(err)

	filename = strings.TrimSuffix(filename, ".go")
	className := filename
	buf := bytes.NewBuffer(retBytes)
	data := tpls.Migration{ClassName: className}
	err = t.Execute(buf, data)
	utils.Error.Panic(err)

	return buf.Bytes()
}

// Verify that the file name is legal
func (controller *MigrateController) validName(name string) bool {
	re, err := regexp.Compile("([0-9]|[a-z]|[A-Z]|_)")
	if err != nil {
		utils.Error.Panic(err)
	}
	name = re.ReplaceAllString(name, "")
	return len(name) == 0
}

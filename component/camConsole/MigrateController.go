package camConsole

import (
	"bytes"
	"fmt"
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/base/camUtils"
	"html/template"
	"regexp"
	"strings"
)

//
type MigrateController struct {
	ConsoleController
}

// Install xorm reverse command
func (ctrl *MigrateController) Install() {
	_ = camUtils.Console.Start("go get xorm.io/reverse@v0.1.1")
}

// Gen xorm model's files
func (ctrl *MigrateController) Generate() {
	_ = camUtils.Console.Start("reverse -f " + camUtils.File.GetRunPath() + "/database/generate.yml")
}

// create migration's file
func (ctrl *MigrateController) Create() {
	name := ctrl.GetArgv(0)
	if len(name) == 0 {
		name = "new_migrate"
	}
	if !ctrl.validName(name) {
		fmt.Println("Illegal name: " + name + ". Chars must in [a-z][A-Z][0-9]_")
		return
	}

	var err error

	console := ctrl.GetConsoleComponent()
	dbDir := console.config.DatabaseDir
	migrateDir := dbDir + "/migrations"
	if !camUtils.File.Exists(migrateDir) {
		err = camUtils.File.Mkdir(migrateDir)
		if err != nil {
			panic(err)
		}
	}
	filename := ctrl.getFilename(name)
	absFilename := migrateDir + "/" + filename
	fmt.Println("General filename...")
	fmt.Println("\t" + absFilename)
	fmt.Print("Do you want to create the above files?[Y/N]:")
	if !camUtils.Console.IsPressY() {
		return
	}

	content := ctrl.getMigrationContent(filename)
	err = camUtils.File.WriteFile(absFilename, content)
	if err != nil {
		panic(err)
	}

	fmt.Print("\nDone.")
}

// migrate up
func (ctrl *MigrateController) Up() {
	db := camStatics.App.GetDB()
	if db == nil {
		panic("no database.")
	}

	versionList := ctrl.GetConsoleComponent().GetMigrateUpVersionList()
	if len(versionList) == 0 {
		fmt.Println("No new versions need to be up")
		return
	}

	fmt.Println("List of versions:")
	for _, version := range versionList {
		fmt.Println("\t" + version)
	}
	fmt.Println("Do you want to up the above version?[Y/N]:")
	if !camUtils.Console.IsPressY() {
		return
	}

	ctrl.GetConsoleComponent().MigrateUp()
}

// migrate down
func (ctrl *MigrateController) Down() {
	ctrl.GetConsoleComponent().MigrateDown()
}

// Generate code's file using database's tables
// TODO Waiting https://gitea.com/xorm/reverse has release version
//func (ctrl *MigrateController) Reverse() {
//	if !camUtils.Console.HasCommand("reverse") {
//		err := camUtils.Console.Start("go get xorm.io/reverse")
//		if err != nil {
//			panic(err)
//		}
//	}
//
//
//	err := camUtils.Console.Start("reverse -f ./../common/xorm-reverse/config.yml")
//	if err != nil {
//		panic(err)
//	}
//}

// get migration's filename. only filename, not absolute filename
func (ctrl *MigrateController) getFilename(name string) string {
	// generate filename
	id := camUtils.Migrate.IdByDatetime()
	return "m" + id + "_" + name + ".go"
}

// orm struct's file template
func (ctrl *MigrateController) getTpl() string {
	return `package migrations

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/component/camConsole"
)

func init() {
	m := new({{ .ClassName}})
	cam.App.AddMigration(m)
}

type {{ .ClassName}} struct {
	camConsole.BaseMigration
}

// up
func (m *{{ .ClassName}}) Up() {
}

// down
func (m *{{ .ClassName}}) Down() {
}
`
}

// get migration file content
func (ctrl *MigrateController) getMigrationContent(filename string) []byte {
	var err error
	var retBytes = []byte("")

	t, err := template.New(filename).Parse(ctrl.getTpl())
	if err != nil {
		panic(err)
	}

	filename = strings.TrimSuffix(filename, ".go")
	className := filename
	buf := bytes.NewBuffer(retBytes)
	data := MigrationTpl{ClassName: className}
	err = t.Execute(buf, data)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}

// Verify that the file name is legal
func (ctrl *MigrateController) validName(name string) bool {
	re, err := regexp.Compile("([0-9]|[a-z]|[A-Z]|_)")
	if err != nil {
		panic(err)
	}
	name = re.ReplaceAllString(name, "")
	return len(name) == 0
}

func (ctrl *MigrateController) Sync() {
	// TODO
}

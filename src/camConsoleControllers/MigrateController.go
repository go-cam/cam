package camConsoleControllers

import (
	"bytes"
	"fmt"
	"github.com/go-cam/cam/camModels/camModelsTpls"
	"github.com/go-cam/cam/camUtils"
	"html/template"
	"regexp"
	"strings"
)

//
type MigrateController struct {
	baseConsoleController
}

// create migration's file
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

	db := controller.GetDatabaseComponent()
	migrateDir := db.GetMigrateDir()
	if !camUtils.File.Exists(migrateDir) {
		err = camUtils.File.Mkdir(migrateDir)
		camUtils.Error.Panic(err)
	}
	filename := controller.getFilename(name)
	absFilename := migrateDir + "/" + filename
	fmt.Println("General filename...")
	fmt.Println("\t" + absFilename)
	fmt.Print("Do you want to create the above files?[Y/N]:")
	if !camUtils.Console.IsPressY() {
		return
	}
	//input := bufio.NewScanner(os.Stdin)
	//if !input.Scan() {
	//	return
	//}
	//str := strings.ToLower(input.Text())
	//if str != "y" {
	//	return
	//}

	content := controller.getMigrationContent(filename)
	err = camUtils.File.WriteFile(absFilename, content)
	camUtils.Error.Panic(err)

	fmt.Print("\nDone.")
}

// migrate up
func (controller *MigrateController) Up() {
	db := controller.GetDatabaseComponent()
	if db == nil {
		panic("no database.")
	}

	versionList := db.GetMigrateUpVersionList()
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
	id := camUtils.Migrate.IdByDatetime()
	return "m" + id + "_" + name + ".go"
}

// orm struct's file template
func (controller *MigrateController) getTpl() string {
	return `package migrations

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/camModels"
)

func init() {
	m := new({{ .ClassName}})
	cam.App.AddMigration(m)
}

type {{ .ClassName}} struct {
	camModels.Migration
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
	camUtils.Error.Panic(err)

	filename = strings.TrimSuffix(filename, ".go")
	className := filename
	buf := bytes.NewBuffer(retBytes)
	data := camModelsTpls.Migration{ClassName: className}
	err = t.Execute(buf, data)
	camUtils.Error.Panic(err)

	return buf.Bytes()
}

// Verify that the file name is legal
func (controller *MigrateController) validName(name string) bool {
	re, err := regexp.Compile("([0-9]|[a-z]|[A-Z]|_)")
	if err != nil {
		camUtils.Error.Panic(err)
	}
	name = re.ReplaceAllString(name, "")
	return len(name) == 0
}
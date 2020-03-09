package template

import (
	"fmt"
	"github.com/go-cam/cam/base/camConstants"
	"github.com/go-cam/cam/base/camUtils"
	"os"
)

type Command struct {
}

func NewCommand() *Command {
	c := new(Command)
	if !camUtils.File.Exists("cam.json") {
		c.initCamJson()
	}
	return c
}

func (c *Command) Run() {
	argLen := len(os.Args)
	if argLen < 2 {
		c.help()
		return
	}
}

func (c *Command) help() {
	fmt.Println("Cam help:")
	fmt.Println("    help        print help info")
	fmt.Println("    init        init project")
}

// init project's config: cam.json
func (c *Command) initCamJson() {
	cam := new(Cam)
	cam.Modules = []*CamModule{}

	// common
	commonModule := new(CamModule)
	commonModule.Name = "common"
	commonModule.Type = camConstants.CamModuleTypeLib
	cam.Modules = append(cam.Modules, commonModule)

	// server
	serverModule := new(CamModule)
	serverModule.Name = "server"
	serverModule.Type = camConstants.CamModuleTypeApp
	cam.Modules = append(cam.Modules, serverModule)

	jsonBytes := camUtils.Json.EncodeBeautiful(cam)
	err := camUtils.File.WriteFile("cam.json", jsonBytes)
	camUtils.Error.Panic(err)
}

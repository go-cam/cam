package cmd

import (
	"flag"
	"fmt"
	"github.com/go-cam/cam/base/camUtils"
	"os"
	"strings"
)

var (
	m          = flag.String("m", "", "Operating module")
	module     = flag.String("module", "", "Operating module")
	mt         = flag.String("mt", "application", "moduleType. [application] [library] [grpc]")
	moduleType = flag.String("moduleType", "application", "moduleType. [application] [library] [grpc]")
)

type CamManager struct {
	args []string
}

func NewCamManager() *CamManager {
	c := new(CamManager)
	if !camUtils.File.Exists("cam.json") {
		c.initCamJson()
	}
	return c
}

func (m *CamManager) Run() {
	m.loadArgs()
	flag.Parse()

	if len(m.args) == 0 {
		m.help()
		return
	}

	switch m.args[0] {
	case "g":
		fallthrough
	case "generate":
		m.generate()
	case "v":
		fallthrough
	case "version":
		fmt.Println(camUtils.C.Version())
	case "help":
		m.help()
	case "test":
		m.getCam()
	default:
		m.unknownAction()
	}
}

// load args.
//
// Example 1:
//  $ ./cam generate projectA --m=application
//  args = []string{"generate", "projectA"}
//
// Example 2:
//  $ ./cam generate --m=application projectA
//  args = []string{"generate", "projectA"}
func (m *CamManager) loadArgs() {
	m.args = []string{}

	for i, arg := range os.Args {
		if i == 0 {
			continue
		}
		if strings.HasPrefix(arg, "--") {
			continue
		}
		m.args = append(m.args, arg)
	}
}

func (m *CamManager) help() {
	fmt.Println("Cam help:")
	fmt.Println("    help        print help info")
	fmt.Println("    init        init project")
}

func (m *CamManager) unknownAction() {
	fmt.Println("unknown command: [" + m.args[0] + "]")
}

// init project's config: cam.json
func (m *CamManager) initCamJson() {
	cam := NewCam()

	// common
	cam.AddModule("common", &CamModule{Type: CamModuleTypeLibrary})
	// server
	cam.AddModule("server", &CamModule{Type: CamModuleTypeApplication})
	// server-grpc
	cam.AddModule("server-grpc", &CamModule{Type: CamModuleTypeGrpc})

	m.saveCam(cam)
}

// Get the cam.json content
func (m *CamManager) getCam() *Cam {
	path := camUtils.File.GetRunPath()
	absFilename := path + "/cam.json"
	content, err := camUtils.File.ReadFile(absFilename)
	if err != nil {
		panic(err)
	}

	cam := new(Cam)
	camUtils.Json.DecodeToObj(content, cam)

	return cam
}

func (m *CamManager) saveCam(cam *Cam) {
	jsonBytes := camUtils.Json.EncodeBeautiful(cam)
	path := camUtils.File.GetRunPath()
	err := camUtils.File.WriteFile(path+"/cam.json", jsonBytes)
	if err != nil {
		panic(err)
	}
}

// TODO
func (m *CamManager) generate() {
	if len(m.args) < 3 {
		return
	}

	switch m.args[1] {
	case "m":
		fallthrough
	case "module":
		m.generateModule()
	}
}

func (m *CamManager) getModuleType() {

}

// TODO
func (m *CamManager) generateModule() {
	switch {

	}
}

func (m *CamManager) generateGrpc() {

}

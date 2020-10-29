package cmd

// generate file class
// TODO
type generate struct {
}

func (g *generate) genApp(name string) {

}

// config/app.go
func (g *generate) genAppConfigAppGo(name string) string {
	return `
package config

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camBase"
	"test/backend/controllers"
	"test/backend/structs"
)

func GetApp() camBase.AppConfigInterface {
	config := cam.NewConfig()
	config.ComponentDict = map[string]camBase.ComponentConfigInterface{
		"http":    httpServer(),
	}
	return config
}

func httpServer() camBase.ComponentConfigInterface {
	config := cam.NewHttpConfig(10800)
	config.SessionName = "` + name + `"
	config.RecoverRoute("test/recover")
	config.SetContextStruct(&structs.HttpContextAo{})

	config.Register(&controllers.TestController{})
	return config
}
`
}

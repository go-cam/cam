package camLog

import (
	"fmt"
	"github.com/go-cam/cam/camBase"
	"github.com/go-cam/cam/camUtils"
	"reflect"
)

// log components
type LogComponent struct {
	camBase.Component
	config *LogComponentConfig

	logRootDir string
}

// on App init
func (component *LogComponent) Init(configInterface camBase.ComponentConfigInterface) {
	component.Component.Init(configInterface)

	configValue := reflect.ValueOf(configInterface)
	var config *LogComponentConfig
	if configValue.Kind() == reflect.Ptr {
		config = configValue.Interface().(*LogComponentConfig)
	} else if configValue.Kind() == reflect.Struct {
		configStruct := configValue.Interface().(LogComponentConfig)
		config = &configStruct
	} else {
		panic("illegal config")
	}
	component.config = config

	// log output path
	component.logRootDir = camUtils.File.GetRunPath() + "/runtime/log"
	if !camUtils.File.Exists(component.logRootDir) {
		err := camUtils.File.Mkdir(component.logRootDir)
		camUtils.Error.Panic(err)
	}

}

// on App start
func (component *LogComponent) Start() {
	component.Component.Start()
}

// before App destroy
func (component *LogComponent) Stop() {
	component.Component.Stop()
}

func (component *LogComponent) baseLog(logType string, title string, content string) error {
	datetime := camUtils.Time.NowDateTime()
	line := "[" + datetime + " " + logType + " " + title + "] " + content
	filename := component.logRootDir + "/App.log"

	if component.config.IsPrint {
		fmt.Println(line)
	}
	if component.config.IsWrite {
		return camUtils.File.AppendFile(filename, []byte(line+"\n"))
	}
	return nil
}

func (component *LogComponent) Debug(title string, content string) error {
	return component.baseLog("debug", title, content)
}

func (component *LogComponent) Info(title string, content string) error {
	return component.baseLog("info", title, content)
}

func (component *LogComponent) Warn(title string, content string) error {
	return component.baseLog("warning", title, content)
}

func (component *LogComponent) Error(title string, content string) error {
	return component.baseLog("error", title, content)
}

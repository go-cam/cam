package camComponents

import (
	"fmt"
	"github.com/go-cam/cam/camBase"
	"github.com/go-cam/cam/camConfigs"
	"github.com/go-cam/cam/camUtils"
	"reflect"
)

// log components
type Log struct {
	camBase.Component
	config *camConfigs.Log

	logRootDir string
}

// on App init
func (component *Log) Init(configInterface camBase.ComponentConfigInterface) {
	component.Component.Init(configInterface)

	configValue := reflect.ValueOf(configInterface)
	var config *camConfigs.Log
	if configValue.Kind() == reflect.Ptr {
		config = configValue.Interface().(*camConfigs.Log)
	} else if configValue.Kind() == reflect.Struct {
		configStruct := configValue.Interface().(camConfigs.Log)
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
func (component *Log) Start() {
	component.Component.Start()
}

// before App destroy
func (component *Log) Stop() {
	component.Component.Stop()
}

func (component *Log) baseLog(logType string, title string, content string) error {
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

func (component *Log) Info(title string, content string) error {
	return component.baseLog("info", title, content)
}

func (component *Log) Warn(title string, content string) error {
	return component.baseLog("warning", title, content)
}

func (component *Log) Error(title string, content string) error {
	return component.baseLog("error", title, content)
}

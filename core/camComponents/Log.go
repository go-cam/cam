package camComponents

import (
	"fmt"
	"github.com/go-cam/cam/core/camBase"
	"github.com/go-cam/cam/core/camConfigs"
	"github.com/go-cam/cam/core/camUtils"
	"reflect"
)

// log components
type Log struct {
	Base
	config *camConfigs.Log

	logRootDir string
}

// on app init
func (component *Log) Init(configInterface camBase.ConfigComponentInterface) {
	component.Base.Init(configInterface)

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

	// 这是日志输出路径
	component.logRootDir = camUtils.File.GetRunPath() + "/runtime/log"
	if !camUtils.File.Exists(component.logRootDir) {
		err := camUtils.File.Mkdir(component.logRootDir)
		camUtils.Error.Panic(err)
	}

}

// on app start
func (component *Log) Start() {
	component.Base.Start()
}

// before app destroy
func (component *Log) Stop() {
	component.Base.Stop()
}

func (component *Log) baseLog(logType string, title string, content string) error {
	datetime := camUtils.Time.NowDateTime()
	line := "[" + datetime + " " + logType + " " + title + "] " + content
	filename := component.logRootDir + "/app.log"

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

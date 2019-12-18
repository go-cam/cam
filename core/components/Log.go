package components

import (
	"github.com/go-cam/cam/core/base"
	"github.com/go-cam/cam/core/utils"
)

// log components
type Log struct {
	Base

	logRootDir string
}

// on app init
func (component *Log) Init(configInterface base.ConfigComponentInterface) {
	component.Base.Init(configInterface)
	component.logRootDir = utils.File.GetRunPath() + "/runtime/log"
	if !utils.File.Exists(component.logRootDir) {
		err := utils.File.Mkdir(component.logRootDir)
		utils.Error.Panic(err)
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
	datetime := utils.Time.NowDateTime()
	line := "[" + datetime + " " + logType + " " + title + "] " + content
	filename := component.logRootDir + "/app.log"
	return utils.File.AppendFile(filename, []byte(line))
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

package camLog

import (
	"errors"
	"fmt"
	"github.com/go-cam/cam/camBase"
	"github.com/go-cam/cam/camUtils"
	"reflect"
)

// log components
type LogComponent struct {
	camBase.Component
	config *LogComponentConfig

	logRootDir  string
	levelLabels map[camBase.LogLevel]string
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
	component.initLevelLabels()

}

// on App start
func (component *LogComponent) Start() {
	component.Component.Start()
}

// before App destroy
func (component *LogComponent) Stop() {
	component.Component.Stop()
}
func (component *LogComponent) base(level camBase.LogLevel, title string, content string) error {
	if !component.isBaseLevel(level) {
		return errors.New("level is not basic level")
	}

	levelLabel := component.getLevelLabels(level)

	datetime := camUtils.Time.NowDateTime()
	line := "[" + datetime + " " + levelLabel + " " + title + "] " + content
	filename := component.logRootDir + "/App.log"

	if component.isOutputLevel(level, component.config.PrintLevel) {
		fmt.Println(line)
	}
	if component.isOutputLevel(level, component.config.WriteLevel) {
		return camUtils.File.AppendFile(filename, []byte(line+"\n"))
	}
	return nil
}

func (component *LogComponent) Debug(title string, content string) error {
	return component.base(LevelDebug, title, content)
}

func (component *LogComponent) Info(title string, content string) error {
	return component.base(LevelInfo, title, content)
}

func (component *LogComponent) Warn(title string, content string) error {
	return component.base(LevelWarn, title, content)
}

func (component *LogComponent) Error(title string, content string) error {
	return component.base(LevelError, title, content)
}

// init level labels
func (component *LogComponent) initLevelLabels() {
	component.levelLabels = map[camBase.LogLevel]string{
		LevelDebug: "D",
		LevelInfo:  "I",
		LevelWarn:  "W",
		LevelError: "E",
	}
}

// get level labels
func (component *LogComponent) getLevelLabels(level camBase.LogLevel) string {
	label, has := component.levelLabels[level]
	if !has {
		return ""
	}
	return label
}

// Whether output is required for detection level
func (component *LogComponent) isOutputLevel(targetLevel camBase.LogLevel, outputLevel camBase.LogLevel) bool {
	return targetLevel&outputLevel == targetLevel
}

// Whether level is basic level (debug, info, warn, error)
func (component *LogComponent) isBaseLevel(level camBase.LogLevel) bool {
	return level == LevelDebug || level == LevelInfo || level == LevelWarn || level == LevelError
}

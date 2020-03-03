package camLog

import (
	"errors"
	"fmt"
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camConstants"
	"github.com/go-cam/cam/base/camUtils"
	"strconv"
	"strings"
	"time"
)

// log components
type LogComponent struct {
	camBase.Component
	config *LogComponentConfig

	logRootDir             string                      // file log dir
	levelLabels            map[camBase.LogLevel]string // log level label. It will output on console and file
	lastCheckFileTimestamp int64                       // last check file time
	titleMaxLen            int                         // title max len
}

// on App init
func (component *LogComponent) Init(configI camBase.ComponentConfigInterface) {
	component.Component.Init(configI)

	var ok bool
	component.config, ok = configI.(*LogComponentConfig)
	if !ok {
		camBase.App.Error("LogComponent", "invalid config")
	}

	// log output path
	component.logRootDir = camUtils.File.GetRunPath() + "/runtime/log"
	if !camUtils.File.Exists(component.logRootDir) {
		err := camUtils.File.Mkdir(component.logRootDir)
		camUtils.Error.Panic(err)
	}
	component.initLevelLabels()
	component.lastCheckFileTimestamp = 0
	component.titleMaxLen = 20

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
	spaceTitle := component.addSpaceToTitle(title)
	line := "[" + datetime + " " + levelLabel + " " + spaceTitle + "] " + content
	filename := component.getLogFilename()

	if component.isOutputLevel(level, component.config.PrintLevel) {
		fmt.Println(line)
	}
	if component.isOutputLevel(level, component.config.WriteLevel) {
		component.checkAndRenameFile()
		return camUtils.File.AppendFile(filename, []byte(line+"\n"))
	}
	return nil
}

func (component *LogComponent) Debug(title string, content string) error {
	return component.base(camConstants.LevelDebug, title, content)
}

func (component *LogComponent) Info(title string, content string) error {
	return component.base(camConstants.LevelInfo, title, content)
}

func (component *LogComponent) Warn(title string, content string) error {
	return component.base(camConstants.LevelWarn, title, content)
}

func (component *LogComponent) Error(title string, content string) error {
	return component.base(camConstants.LevelError, title, content)
}

// init level labels
func (component *LogComponent) initLevelLabels() {
	component.levelLabels = map[camBase.LogLevel]string{
		camConstants.LevelDebug: "D",
		camConstants.LevelInfo:  "I",
		camConstants.LevelWarn:  "W",
		camConstants.LevelError: "E",
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
	return level == camConstants.LevelDebug || level == camConstants.LevelInfo || level == camConstants.LevelWarn || level == camConstants.LevelError
}

// Check if the file exceeds the configured size
func (component *LogComponent) checkAndRenameFile() {
	now := time.Now().Unix()
	if now < component.lastCheckFileTimestamp+10 {
		return
	}
	component.lastCheckFileTimestamp = now

	filename := component.getLogFilename()
	fileSize := camUtils.File.Size(filename)
	if fileSize >= component.config.FileMaxSize {
		newFilename := component.logRootDir + "/app_" + strconv.FormatInt(now, 10) + ".log"
		err := camUtils.File.Rename(filename, newFilename)
		if err != nil {
			_ = component.Error("LogComponent.checkAndRenameFile", err.Error())
		}
	}
}

// get log absolute filename
func (component *LogComponent) getLogFilename() string {
	return component.logRootDir + "/app.log"
}

// add space before title
func (component *LogComponent) addSpaceToTitle(title string) string {

	titleLen := len(title)
	if titleLen > component.titleMaxLen {
		component.titleMaxLen = titleLen
	}

	strArr := make([]string, component.titleMaxLen)
	spaceNum := component.titleMaxLen - titleLen
	i := 0
	for ; i < spaceNum; i++ {
		strArr[i] = " "
	}
	titleArr := strings.Split(title, "")
	for j := 0; j < titleLen; j++ {
		index := i + j
		strArr[index] = titleArr[j]
	}

	return strings.Join(strArr, "")
}

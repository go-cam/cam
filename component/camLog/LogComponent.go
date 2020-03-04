package camLog

import (
	"errors"
	"fmt"
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camConstants"
	"github.com/go-cam/cam/base/camUtils"
	"strconv"
	"strings"
	"sync"
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
	logChan                chan bool
	mutex                  sync.Mutex
}

// on App init
func (comp *LogComponent) Init(configI camBase.ComponentConfigInterface) {
	comp.Component.Init(configI)

	var ok bool
	comp.config, ok = configI.(*LogComponentConfig)
	if !ok {
		camBase.App.Error("LogComponent", "invalid config")
	}

	// log output path
	comp.logRootDir = camUtils.File.GetRunPath() + "/runtime/log"
	if !camUtils.File.Exists(comp.logRootDir) {
		err := camUtils.File.Mkdir(comp.logRootDir)
		camUtils.Error.Panic(err)
	}
	comp.initLevelLabels()
	comp.lastCheckFileTimestamp = 0
	comp.titleMaxLen = 32
}

// on App start
func (comp *LogComponent) Start() {
	comp.Component.Start()
}

// before App destroy
func (comp *LogComponent) Stop() {
	comp.Component.Stop()
}
func (comp *LogComponent) base(level camBase.LogLevel, title string, content string) error {
	comp.mutex.Lock()
	defer func() {
		comp.mutex.Unlock()
	}()

	if !comp.isBaseLevel(level) {
		return errors.New("level is not basic level")
	}
	isPrint := comp.isOutputLevel(level, comp.config.PrintLevel)
	isWrite := comp.isOutputLevel(level, comp.config.WriteLevel)
	if !isPrint && !isWrite {
		return nil
	}

	var err error = nil
	levelLabel := comp.getLevelLabels(level)
	datetime := camUtils.Time.NowDateTime()
	spaceTitle := comp.addSpaceToTitle(title)
	line := "[ " + datetime + " " + levelLabel + " | " + spaceTitle + " ] " + content
	filename := comp.getLogFilename()

	if isPrint {
		fmt.Println(line)
	}
	if isWrite {
		comp.checkAndRenameFile()
		err = camUtils.File.AppendFile(filename, []byte(line+"\n"))
	}
	return err
}

func (comp *LogComponent) Debug(title string, content string) error {
	return comp.base(camConstants.LevelDebug, title, content)
}

func (comp *LogComponent) Info(title string, content string) error {
	return comp.base(camConstants.LevelInfo, title, content)
}

func (comp *LogComponent) Warn(title string, content string) error {
	return comp.base(camConstants.LevelWarn, title, content)
}

func (comp *LogComponent) Error(title string, content string) error {
	return comp.base(camConstants.LevelError, title, content)
}

// init level labels
func (comp *LogComponent) initLevelLabels() {
	comp.levelLabels = map[camBase.LogLevel]string{
		camConstants.LevelDebug: "DEBUG",
		camConstants.LevelInfo:  "INFO ",
		camConstants.LevelWarn:  "WARN ",
		camConstants.LevelError: "ERROR",
	}
}

// get level labels
func (comp *LogComponent) getLevelLabels(level camBase.LogLevel) string {
	label, has := comp.levelLabels[level]
	if !has {
		return ""
	}
	return label
}

// Whether output is required for detection level
func (comp *LogComponent) isOutputLevel(targetLevel camBase.LogLevel, outputLevel camBase.LogLevel) bool {
	return targetLevel&outputLevel == targetLevel
}

// Whether level is basic level (debug, info, warn, error)
func (comp *LogComponent) isBaseLevel(level camBase.LogLevel) bool {
	return level == camConstants.LevelDebug || level == camConstants.LevelInfo || level == camConstants.LevelWarn || level == camConstants.LevelError
}

// Check if the file exceeds the configured size
func (comp *LogComponent) checkAndRenameFile() {
	now := time.Now().Unix()
	if now < comp.lastCheckFileTimestamp+10 {
		return
	}
	comp.lastCheckFileTimestamp = now

	filename := comp.getLogFilename()
	fileSize := camUtils.File.Size(filename)
	if fileSize >= comp.config.FileMaxSize {
		newFilename := comp.logRootDir + "/app_" + strconv.FormatInt(now, 10) + ".log"
		err := camUtils.File.Rename(filename, newFilename)
		if err != nil {
			_ = comp.Error("LogComponent.checkAndRenameFile", err.Error())
		}
	}
}

// get log absolute filename
func (comp *LogComponent) getLogFilename() string {
	return comp.logRootDir + "/app.log"
}

// add space before title
func (comp *LogComponent) addSpaceToTitle(title string) string {

	titleLen := len(title)
	if titleLen > comp.titleMaxLen {
		comp.titleMaxLen = titleLen
	}

	spaceNum := comp.titleMaxLen - titleLen
	strArr := make([]string, spaceNum)

	return title + strings.Join(strArr, " ")
}

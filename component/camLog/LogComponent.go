package camLog

import (
	"errors"
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/base/camUtils"
	"github.com/go-cam/cam/component"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// log components
type LogComponent struct {
	component.Component
	config *LogComponentConfig

	logRootDir             string                         // file log dir
	levelLabels            map[camStatics.LogLevel]string // log level label. It will output on console and file
	lastCheckFileTimestamp int64                          // last check file time
	titleMaxLen            int                            // title max len
	consoleLogger          *log.Logger                    // console logger
	fileLogger             *log.Logger                    // file logger
	logFile                *os.File                       // log file
	fileRenameMutex        sync.Mutex                     // log file rename mutex
}

// on App init
func (comp *LogComponent) Init(configI camStatics.ComponentConfigInterface) {
	comp.Component.Init(configI)

	var ok bool
	comp.config, ok = configI.(*LogComponentConfig)
	if !ok {
		camStatics.App.Fatal("LogComponent", "invalid config")
		return
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
	comp.consoleLogger = log.New(os.Stdout, comp.config.Prefix, comp.config.Flag)
	comp.fileLogger = log.New(nil, comp.config.Prefix, comp.config.Flag)
	comp.resetFileLoggerOutput()
}

// on App start
func (comp *LogComponent) Start() {
	comp.Component.Start()
}

// before App destroy
func (comp *LogComponent) Stop() {
	defer comp.Component.Stop()
}
func (comp *LogComponent) Record(level camStatics.LogLevel, title string, content string) error {
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
	spaceTitle := comp.addSpaceToTitle(title)
	line := "[ " + levelLabel + " | " + spaceTitle + " ] " + content

	if isPrint {
		comp.consoleLogger.Println(line)
	}
	if isWrite {
		comp.checkAndRenameFile()
		comp.fileLogger.Println(line)
	}
	return err
}

// init level labels
func (comp *LogComponent) initLevelLabels() {
	comp.levelLabels = map[camStatics.LogLevel]string{
		camStatics.LevelTrace: "TRACE",
		camStatics.LevelDebug: "DEBUG",
		camStatics.LevelInfo:  "INFO ",
		camStatics.LevelWarn:  "WARN ",
		camStatics.LevelError: "ERROR",
		camStatics.LevelFatal: "FATAL",
	}
}

// get level labels
func (comp *LogComponent) getLevelLabels(level camStatics.LogLevel) string {
	label, has := comp.levelLabels[level]
	if !has {
		return ""
	}
	return label
}

// Whether output is required for detection level
func (comp *LogComponent) isOutputLevel(targetLevel camStatics.LogLevel, outputLevel camStatics.LogLevel) bool {
	return targetLevel&outputLevel == targetLevel
}

// Whether level is basic level (debug, info, warn, error)
func (comp *LogComponent) isBaseLevel(level camStatics.LogLevel) bool {
	_, has := comp.levelLabels[level]
	return has
}

// Check if the file exceeds the configured size
func (comp *LogComponent) checkAndRenameFile() {
	comp.fileRenameMutex.Lock()
	defer comp.fileRenameMutex.Unlock()

	now := time.Now().Unix()
	if now < comp.lastCheckFileTimestamp+10 {
		return
	}
	comp.lastCheckFileTimestamp = now

	// close app.log
	if comp.logFile == nil {
		return
	}
	comp.fileLogger.SetOutput(nil)
	err := comp.logFile.Close()
	if err != nil {
		_ = comp.Record(camStatics.LevelFatal, "LogComponent.checkAndRenameFile", "failed to close file. err: "+err.Error())
		return
	}

	// rename app.log
	filename := comp.getLogFilename()
	fileSize := camUtils.File.Size(filename)
	if fileSize >= comp.config.FileMaxSize {
		newFilename := comp.logRootDir + "/app_" + strconv.FormatInt(now, 10) + ".log"
		err := camUtils.File.Rename(filename, newFilename)
		if err != nil {
			_ = comp.Record(camStatics.LevelFatal, "LogComponent.checkAndRenameFile", "failed to rename. err: "+err.Error())
			return
		}
	}

	// new and set io.writer
	comp.resetFileLoggerOutput()
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

// reset fileLogger output io.writer
func (comp *LogComponent) resetFileLoggerOutput() {
	comp.createLogFile()

	var err error
	comp.logFile, err = os.OpenFile(comp.getLogFilename(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		camStatics.App.Fatal("LogComponent.getLogFileWriter", err.Error())
		return
	}

	comp.fileLogger.SetOutput(comp.logFile)
}

// create log file
func (comp *LogComponent) createLogFile() {
	logFilename := comp.getLogFilename()
	if camUtils.File.Exists(logFilename) {
		return
	}

	err := camUtils.File.WriteFile(logFilename, []byte(""))
	if err != nil {
		camStatics.App.Fatal("LogComponent.createLogFile", "can't create log file. err: "+err.Error())
	}
}

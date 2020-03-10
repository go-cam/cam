package camLog

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camConstants"
)

type LogComponentConfig struct {
	camBase.ComponentConfig

	// print log level
	// Binary switch used.
	// constant defined in constant.go and cam.constant.go
	PrintLevel camBase.LogLevel
	// write log level
	// Binary switch used
	// constant defined in constant.go and cam.constant.go
	WriteLevel camBase.LogLevel
	// log file max size
	// When the log file exceeds this size, a new file will be created. Old file will be renamed
	FileMaxSize int64
}

func NewLogConfig() *LogComponentConfig {
	config := new(LogComponentConfig)
	config.PrintLevel = camConstants.LevelTrace | camConstants.LevelDebug | camConstants.LevelInfo | camConstants.LevelWarn | camConstants.LevelError | camConstants.LevelFatal
	config.WriteLevel = camConstants.LevelInfo | camConstants.LevelWarn | camConstants.LevelError | camConstants.LevelFatal
	config.FileMaxSize = 10 * 1024 * 1024
	return config
}
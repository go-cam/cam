package camLog

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camConstants"
)

type LogComponentConfig struct {
	camBase.ComponentConfig

	// Deprecated: replace by PrintLevel
	IsPrint bool // Whether print log to console
	// Deprecated: replace by WriteLevel
	IsWrite bool // Whether write log to file

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
	config.PrintLevel = camConstants.LevelDebug | camConstants.LevelInfo | camConstants.LevelWarn | camConstants.LevelError
	config.WriteLevel = camConstants.LevelInfo | camConstants.LevelWarn | camConstants.LevelError
	config.FileMaxSize = 10 * 1024 * 1024
	return config
}

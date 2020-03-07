package cam

import (
	"github.com/go-cam/cam/base/camConstants"
)

// app
const (
	AppStatusInit  = camConstants.AppStatusInit
	AppStatusStart = camConstants.AppStatusStart
	AppStatusStop  = camConstants.AppStatusStop
)

// Log
const (
	LogLevelTrace = camConstants.LevelTrace // log level: trace
	LogLevelDebug = camConstants.LevelDebug // log level: debug
	LogLevelInfo  = camConstants.LevelInfo  // log level: info
	LogLevelWarn  = camConstants.LevelWarn  // log level: warning
	LogLevelError = camConstants.LevelError // log level: error
	LogLevelFatal = camConstants.LevelFatal // log level: fatal

	LogLevelNone    = 0                                                                                           // none level
	LogLevelSuggest = LogLevelInfo | LogLevelWarn | LogLevelError | LogLevelFatal                                 // suggest level
	LogLevelAll     = LogLevelTrace | LogLevelDebug | LogLevelInfo | LogLevelWarn | LogLevelError | LogLevelFatal // all level
)

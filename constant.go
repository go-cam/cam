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
	LogLevelDebug = camConstants.LevelDebug // debug log
	LogLevelInfo  = camConstants.LevelInfo  // info log
	LogLevelWarn  = camConstants.LevelWarn  // warning log
	LogLevelError = camConstants.LevelError // error log

	LogLevelNone    = 0                                                           // none level
	LogLevelSuggest = LogLevelInfo | LogLevelWarn | LogLevelError                 // suggest level
	LogLevelAll     = LogLevelDebug | LogLevelInfo | LogLevelWarn | LogLevelError // all level
)

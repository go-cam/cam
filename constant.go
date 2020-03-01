package cam

import "github.com/go-cam/cam/camLog"

// package camLog
const (
	LogLevelDebug = camLog.LevelDebug // debug log
	LogLevelInfo  = camLog.LevelInfo  // info log
	LogLevelWarn  = camLog.LevelWarn  // warning log
	LogLevelError = camLog.LevelError // error log

	LogLevelNone    = 0                                                           // none level
	LogLevelSuggest = LogLevelInfo | LogLevelWarn | LogLevelError                 // suggest level
	LogLevelAll     = LogLevelDebug | LogLevelInfo | LogLevelWarn | LogLevelError // all level
)

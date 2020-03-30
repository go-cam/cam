package cam

import (
	"github.com/go-cam/cam/base/camConstants"
)

// Log
const (
	LogLevelTrace   = camConstants.LevelTrace   // log level: trace
	LogLevelDebug   = camConstants.LevelDebug   // log level: debug
	LogLevelInfo    = camConstants.LevelInfo    // log level: info
	LogLevelWarn    = camConstants.LevelWarn    // log level: warning
	LogLevelError   = camConstants.LevelError   // log level: error
	LogLevelFatal   = camConstants.LevelFatal   // log level: fatal
	LogLevelNone    = camConstants.LevelNone    // none
	LogLevelSuggest = camConstants.LevelSuggest // suggest this level to write file
	LogLevelAll     = camConstants.LevelAll     // all
)

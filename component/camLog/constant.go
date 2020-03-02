package camLog

import "github.com/go-cam/cam/base/camBase"

const (
	LevelDebug camBase.LogLevel = 1      // debug log
	LevelInfo  camBase.LogLevel = 1 << 1 // info log
	LevelWarn  camBase.LogLevel = 1 << 2 // warning log
	LevelError camBase.LogLevel = 1 << 3 // error log
)

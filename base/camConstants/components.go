package camConstants

import "github.com/go-cam/cam/base/camBase"

// DatabaseComponent
const (
	DatabaseDriverMysql = "mysql"
)

// LogComponent
const (
	LevelDebug camBase.LogLevel = 1      // debug log
	LevelInfo  camBase.LogLevel = 1 << 1 // info log
	LevelWarn  camBase.LogLevel = 1 << 2 // warning log
	LevelError camBase.LogLevel = 1 << 3 // error log
)

package camConstants

import "github.com/go-cam/cam/base/camBase"

// DatabaseComponent
const (
	DatabaseDriverMysql = "mysql"
)

// LogComponent
const (
	LevelTrace camBase.LogLevel = 1               // trace log
	LevelDebug                  = LevelTrace << 1 // debug log
	LevelInfo                   = LevelDebug << 1 // info log
	LevelWarn                   = LevelInfo << 1  // warning log
	LevelError                  = LevelWarn << 1  // error log
	LevelFatal                  = LevelError << 1 // fatal log
)

package camConstants

import "github.com/go-cam/cam/base/camBase"

// BaseComponent
const (
	RecoverHandlerResultPass  camBase.RecoverHandlerResult = iota // It mean the panic was recover
	RecoverHandlerResultPanic                                     // It mean the panic was not recover and panic again
)

// DatabaseComponent
const (
	DatabaseDriverMysql = "mysql"
)

// LogComponent
const (
	LevelTrace   camBase.LogLevel = 1                                                             // trace log
	LevelDebug                    = LevelTrace << 1                                               // debug log
	LevelInfo                     = LevelDebug << 1                                               // info log
	LevelWarn                     = LevelInfo << 1                                                // warning log
	LevelError                    = LevelWarn << 1                                                // error log
	LevelFatal                    = LevelError << 1                                               // fatal log
	LevelNone                     = 0                                                             // none
	LevelSuggest                  = LevelInfo | LevelWarn | LevelError | LevelFatal               // suggest this level to write file
	LevelAll                      = LevelTrace | LevelDebug | LevelWarn | LevelError | LevelFatal // all
)

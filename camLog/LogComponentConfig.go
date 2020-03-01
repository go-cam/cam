package camLog

import "github.com/go-cam/cam/camBase"

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
}

func NewLogConfig() *LogComponentConfig {
	config := new(LogComponentConfig)
	config.PrintLevel = LevelDebug | LevelInfo | LevelWarn | LevelError
	config.WriteLevel = LevelInfo | LevelWarn | LevelError
	return config
}

package camStatics

// BaseComponent
const (
	RecoverHandlerResultPass  RecoverHandlerResult = iota // It mean the panic was recover
	RecoverHandlerResultPanic                             // It mean the panic was not recover and panic again
)

// DatabaseComponent
const (
	DatabaseDriverMysql = "mysql"
)

// LogComponent
const (
	LogLevelTrace   LogLevel = 1                                                                            // trace log
	LogLevelDebug            = LogLevelTrace << 1                                                           // debug log
	LogLevelInfo             = LogLevelDebug << 1                                                           // info log
	LogLevelWarn             = LogLevelInfo << 1                                                            // warning log
	LogLevelError            = LogLevelWarn << 1                                                            // error log
	LogLevelFatal            = LogLevelError << 1                                                           // fatal log
	LogLevelNone             = 0                                                                            // none
	LogLevelSuggest          = LogLevelInfo | LogLevelWarn | LogLevelError | LogLevelFatal                  // suggest this level to write file
	LogLevelAll              = LogLevelTrace | LogLevelDebug | LogLevelWarn | LogLevelError | LogLevelFatal // all
)

// ValidationComponent
const (
	ModeInterface ValidMode = 1
	ModeTag                 = ModeInterface << 1
	ModeBoth                = ModeInterface | ModeTag
)

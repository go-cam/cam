package camLog

import (
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/component"
	"log"
)

type LogComponentConfig struct {
	component.ComponentConfig

	// print log level
	// Binary switch used.
	// constant defined in constant.go and cam.constant.go
	PrintLevel camStatics.LogLevel
	// write log level
	// Binary switch used
	// constant defined in constant.go and cam.constant.go
	WriteLevel camStatics.LogLevel
	// log file max size
	// When the log file exceeds this size, a new file will be created. Old file will be renamed
	FileMaxSize int64
	// log prefix
	Prefix string
	// log flag
	//
	// See: log.Flag
	//  Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	//  Ltime                         // the time in the local time zone: 01:23:23
	//  Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	//  Llongfile                     // full file name and line number: /a/b/c/d.go:23
	//  Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	//  LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	//  Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	//  LstdFlags     = Ldate | Ltime // initial values for the standard logger
	Flag int
}

func NewLogConfig() *LogComponentConfig {
	config := new(LogComponentConfig)
	config.Component = &LogComponent{}
	config.PrintLevel = camStatics.LevelAll
	config.WriteLevel = camStatics.LevelSuggest
	config.FileMaxSize = 10 * 1024 * 1024
	config.Prefix = "[cam] "
	config.Flag = log.LstdFlags
	return config
}

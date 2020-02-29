package camLog

import "github.com/go-cam/cam/camBase"

type LogComponentConfig struct {
	camBase.ComponentConfig

	IsPrint bool // Whether print log to console
	IsWrite bool // Whether write log to file
}

func NewLogConfig() *LogComponentConfig {
	config := new(LogComponentConfig)
	config.IsPrint = true
	config.IsWrite = true
	return config
}

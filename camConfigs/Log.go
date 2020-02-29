package camConfigs

import "github.com/go-cam/cam/camBase"

type Log struct {
	camBase.ComponentConfig

	IsPrint bool // Whether print log to console
	IsWrite bool // Whether write log to file
}

func NewLogConfig() *Log {
	config := new(Log)
	config.IsPrint = true
	config.IsWrite = true
	return config
}

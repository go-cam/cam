package camConfigs

type Log struct {
	BaseConfig

	IsPrint bool // Whether print log to console
	IsWrite bool // Whether write log to file
}

func NewLogConfig() *Log {
	config := new(Log)
	config.IsPrint = false
	config.IsWrite = true
	return config
}

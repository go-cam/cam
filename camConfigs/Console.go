package camConfigs

// console config
type Console struct {
	BaseConfig
	RouterPlugin
}

// new console config
func NewConsoleConfig() *Console {
	config := new(Console)
	config.RouterPlugin.Init()
	return config
}

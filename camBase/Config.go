package camBase

// base config
type Config struct {
	ConfigComponentInterface
	Component ComponentInterface // Instance of corresponding component
}

// get component instance
func (config *Config) GetComponent() ComponentInterface {
	return config.Component
}

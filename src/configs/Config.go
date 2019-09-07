package configs

// 配置实例。用于封装配置
type Config struct {
	params        map[string]interface{}        // 参数。自定义的配置参数
	componentDict map[string]ComponentInterface // 组件
}

// 创建配置对象
func NewConfig() *Config {
	config := new(Config)
	config.componentDict = map[string]ComponentInterface{}
	config.params = map[string]interface{}{}
	return config
}

// 设置配置参数
func (config *Config) SetParams(params map[string]interface{}) {
	config.params = params
}

// 添加一个组件
func (config *Config) AddComponents(name string, componentConfig ComponentInterface) {
	config.componentDict[name] = componentConfig
}

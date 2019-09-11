package models

import (
	"cin/src/base"
)

// 配置实例。用于封装配置
type Config struct {
	BaseModel
	Params              map[string]interface{}                   // 参数。自定义的配置参数
	ComponentDict       map[string]base.ConfigComponentInterface // 组件。组件名 => 配置
}

// 创建配置对象
func NewConfig() *Config {
	config := new(Config)
	config.ComponentDict = map[string]base.ConfigComponentInterface{}
	config.Params = map[string]interface{}{}
	return config
}
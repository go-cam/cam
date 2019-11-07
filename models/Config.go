package models

import (
	"github.com/cinling/cin/base"
)

// 配置实例。用于封装配置
type Config struct {
	BaseModel
	Params        map[string]interface{}                   // 参数。自定义的配置参数
	ComponentDict map[string]base.ConfigComponentInterface // 组件。组件名 => 配置
}

// 初始化数据
func (config *Config) Init() {
	config.ComponentDict = map[string]base.ConfigComponentInterface{}
	config.Params = map[string]interface{}{}
}

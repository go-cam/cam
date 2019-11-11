package configs

import (
	base2 "github.com/cinling/cin/core/base"
)

// 基础配置
type BaseConfig struct {
	base2.ConfigComponentInterface
	Component base2.ComponentInterface // 需要初始化的组件实例
}

// 获取配置需要初始化的组件
func (config *BaseConfig) GetComponent() base2.ComponentInterface {
	return config.Component
}

package configs

import "github.com/cinling/cin/base"

// 基础配置
type BaseConfig struct {
	base.ConfigComponentInterface
	Component base.ComponentInterface // 需要初始化的组件实例
}

// 获取配置需要初始化的组件
func (config *BaseConfig) GetComponent() base.ComponentInterface {
	return config.Component
}
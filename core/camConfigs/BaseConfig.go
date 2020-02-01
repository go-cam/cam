package camConfigs

import (
	"github.com/go-cam/cam/core/camBase"
)

// 基础配置
type BaseConfig struct {
	camBase.ConfigComponentInterface
	Component camBase.ComponentInterface // 需要初始化的组件实例
}

// 获取配置需要初始化的组件
func (config *BaseConfig) GetComponent() camBase.ComponentInterface {
	return config.Component
}

package config

import (
	"github.com/go-cam/cam"
)

// 获取默认配置
func GetApp() *core.Config {
	config := core.NewConfig()
	return config
}

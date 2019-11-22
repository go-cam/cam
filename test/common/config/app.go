package config

import (
	"github.com/cinling/cam"
)

// 获取默认配置
func GetApp() *cin.Config {
	config := cin.NewConfig()
	return config
}

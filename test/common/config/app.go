package config

import (
	"github.com/go-cam/cam"
)

// 获取默认配置
func GetApp() *cam.Config {
	config := cam.NewConfig()
	return config
}

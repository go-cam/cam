package controllers

import (
	"github.com/go-cam/cam"
)

// 测试控制器
type TestController struct {
	cin.BaseController
}

func (controller *TestController) Test() []byte {
	return []byte("test")
}

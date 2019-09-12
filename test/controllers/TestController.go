package controllers

import cin "cin/src"

// 测试控制器
type TestController struct {
	cin.Controller
}

func (controller *TestController) Test() string {
	return "test"
}
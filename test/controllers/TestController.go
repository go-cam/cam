package controllers

import cin "cin"

// 测试控制器
type TestController struct {
	cin.Controller
}

func (controller *TestController) Test() string {
	return "test"
}
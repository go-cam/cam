package controllers

import (
	cin "cin"
	"cin/utils"
	"fmt"
)

// 测试控制器
type TestController struct {
	cin.BaseController
}

func (controller *TestController) Test() []byte {
	context := controller.GetContext()
	session := context.GetSession()
	value := session.Get("key")
	fmt.Println(value)
	uuid := utils.String.UUID()
	fmt.Print(uuid)
	return []byte("test")
}
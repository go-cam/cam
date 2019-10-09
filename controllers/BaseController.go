package controllers

import (
	"cin/base"
)

// 所有 handler 的基类。主要处理问题是：统一接口、数据库管理
type BaseController struct {
	base.ControllerInterface

	context base.ContextInterface
}

// 初始化
func (controller *BaseController) Init() {

}

// 执行动作前执行的方法
// 如果返回 false 将会返回一个错误
func (controller *BaseController) BeforeAction(action string) bool {
	return true
}

// 执行动作后执行的方法
// 过滤返回结果
func (controller *BaseController) AfterAction(action string, response []byte) []byte {
	return response
}

// TODO
func (controller *BaseController) Get(param string) interface{} {
	return nil
}

// 设置上下文对象
func (controller *BaseController) SetContext(context base.ContextInterface) {
	controller.context = context
}
// 获取上下文对象
func (controller *BaseController) GetContext() base.ContextInterface {
	return controller.context
}

package controllers

import (
	"cin/base"
)

// 所有 handler 的基类。主要处理问题是：统一接口、数据库管理
type BaseHandler struct {
	base.HandlerInterface

	context base.ContextInterface
}

// 初始化
func (handler *BaseHandler) Init() {

}

// 执行动作前执行的方法
// 如果返回 false 将会返回一个错误
func (handler *BaseHandler) BeforeAction(action string) bool {
	return true
}

// 执行动作后执行的方法
// 过滤返回结果
func (handler *BaseHandler) AfterAction(action string, response []byte) []byte {
	return response
}

// TODO
func (handler *BaseHandler) Get(param string) interface{} {
	return nil
}

// 设置上下文对象
func (handler *BaseHandler) SetContext(context base.ContextInterface) {
	handler.context = context
}
// 获取上下文对象
func (handler *BaseHandler) GetContext() base.ContextInterface {
	return handler.context
}

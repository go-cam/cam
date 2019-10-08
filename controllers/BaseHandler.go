package controllers

import "cin/base"

// 所有 handler 的基类。主要处理问题是：统一接口、数据库管理
type BaseHandler struct {
	base.HandlerInterface
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
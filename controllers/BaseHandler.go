package controllers

// 所有handler的基类（用于统一接口）
type HandlerInterface interface {
	// 初始化方法
	Init()
	// 执行动作前执行的方法
	BeforeAction(action string) bool
	// 执行动作后执行的方法
	AfterAction(action string, response string) string

	// 获取请求中的参数
	Get(param string) interface{}
	// 获取参数并转为字符串
	GetString(param string) string
	// 获取参数并转为 int64
	GetInt(param string) int64
}

// 所有 handler 的基类。主要处理问题是：统一接口、数据库管理
type BaseHandler struct {
	HandlerInterface
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
func (handler *BaseHandler) AfterAction(action string, response string) string {
	return response
}
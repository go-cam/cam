package base

// 框架内接口

// 组件类型配置接口
type ConfigComponentInterface interface {
	GetComponent() ComponentInterface
}

// 组件接口
type ComponentInterface interface {
	// 初始化
	Init(configInterface ConfigComponentInterface)
	// 开始
	Start()
	// 停止
	Stop()
}

// 所有handler的基类（用于统一接口）
type HandlerInterface interface {
	// 初始化方法
	Init()
	// 执行动作前执行的方法
	BeforeAction(action string) bool
	// 执行动作后执行的方法
	AfterAction(action string, response []byte) []byte

	// 获取请求中的参数
	Get(param string) interface{}

	// 设置上下文对象
	SetContext(context ContextInterface)
	// 获取上下文对象
	GetContext() ContextInterface
}

// 上下文接口
type ContextInterface interface {
	// 获取 session
	GetSession() SessionInterface
}

// session 接口
type SessionInterface interface {
	// 获取 sessionId
	GetSessionId() string
	// 设置值
	Set(key interface{}, value interface{})
	// 获取值
	Get(key interface{}) interface{}
	// 销毁session
	Destroy()
}
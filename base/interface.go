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

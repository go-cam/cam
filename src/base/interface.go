package base

// 框架内接口

// 组件类型配置接口
type ConfigComponentInterface interface {
	GetComponent() ComponentInterface
}

// 组件接口
type ComponentInterface interface {
	// 生命周期方法
	Init(configInterface ConfigComponentInterface)
	Start(configDict map[string]interface{})
	Run(configDict map[string]interface{})
	Stop(configDict map[string]interface{})
	Destroy(configDict map[string]interface{})

	// 打印日志
	Log(message string)
}

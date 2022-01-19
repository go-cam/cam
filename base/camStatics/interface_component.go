package camStatics

type IComponentConfig interface {
	NewComponent() IComponent
	GetRecoverHandler() RecoverHandler
}

type IComponent interface {
	Init(configInterface IComponentConfig)
	Start()
	Stop()
	SetApp(app IApplication)
}

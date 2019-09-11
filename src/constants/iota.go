package constants

// 存状态常量。不需要知道具体的值
import "cin/src/base"

const (
	ApplicationStatusInit base.ApplicationStatus = iota
	ApplicationStatusStart
	ApplicationStatusRun
	ApplicationStatusStop
	ApplicationStatusDestroy

	WebsocketServerModeAutoHandler base.WebsocketServerMode = iota
	WebsocketServerModeCustom

	ControllerTypeWebsocket base.WebsocketServerMode = iota
	ControllerTypeHttp
)

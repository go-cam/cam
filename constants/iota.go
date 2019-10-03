package constants

// 存状态常量。不需要知道具体的值
import "cin/base"

const (
	ApplicationStatusInit base.ApplicationStatus = iota
	ApplicationStatusStart
	ApplicationStatusStop

	WebsocketServerModeAutoHandler base.WebsocketServerMode = iota
	WebsocketServerModeCustom

	ControllerTypeWebsocket base.WebsocketServerMode = iota
	ControllerTypeHttp
)

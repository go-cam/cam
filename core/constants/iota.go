package constants

// 存状态常量。不需要知道具体的值
import (
	base2 "github.com/cinling/cin/core/base"
)

const (
	ApplicationStatusInit base2.ApplicationStatus = iota
	ApplicationStatusStart
	ApplicationStatusStop

	WebsocketServerModeAutoHandler base2.WebsocketServerMode = iota
	WebsocketServerModeCustom

	ControllerTypeWebsocket base2.WebsocketServerMode = iota
	ControllerTypeHttp
)

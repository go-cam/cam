package constants

// 存状态常量。不需要知道具体的值
import "cin/src/alias"

const (
	ApplicationStatusInit    alias.ApplicationStatus = iota
	ApplicationStatusStart
	ApplicationStatusRun
	ApplicationStatusStop
	ApplicationStatusDestroy

	WebsocketServerModeAutoHandler alias.WebsocketServerMode = iota
	WebsocketServerModeCustom
)

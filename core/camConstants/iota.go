package camConstants

// 存状态常量。不需要知道具体的值
import (
	"github.com/go-cam/cam/core/camBase"
)

const (
	ApplicationStatusInit camBase.ApplicationStatus = iota
	ApplicationStatusStart
	ApplicationStatusStop

	WebsocketServerModeAutoHandler camBase.WebsocketServerMode = iota
	WebsocketServerModeCustom

	ControllerTypeWebsocket camBase.WebsocketServerMode = iota
	ControllerTypeHttp
)

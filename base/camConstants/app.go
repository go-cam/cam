package camConstants

import "github.com/go-cam/cam/base/camBase"

const (
	AppStatusBeforeInit camBase.ApplicationStatus = iota
	AppStatusBeforeStart
	AppStatusAfterStart
	AppStatusBeforeStop
	AppStatusAfterStop
)

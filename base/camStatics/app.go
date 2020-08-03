package camStatics

const (
	AppStatusBeforeInit ApplicationStatus = iota
	AppStatusBeforeStart
	AppStatusAfterStart
	AppStatusBeforeStop
	AppStatusAfterStop
)

package camConstants

import (
	"github.com/go-cam/cam/base/camBase"
)

const (
	ApplicationStatusInit camBase.ApplicationStatus = iota
	ApplicationStatusStart
	ApplicationStatusStop

	CamModuleTypeApp = iota
	CamModuleTypeLib
)

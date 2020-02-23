package camConstants

import (
	"github.com/go-cam/cam/core/camBase"
)

const (
	ApplicationStatusInit camBase.ApplicationStatus = iota
	ApplicationStatusStart
	ApplicationStatusStop

	CamModuleTypeApp = iota
	CamModuleTypeLib
)

package camModels

import (
	"github.com/go-cam/cam/core/camBase"
)

type CamModule struct {
	Name string                `json:"name"`
	Type camBase.CamModuleType `json:"type"`
}

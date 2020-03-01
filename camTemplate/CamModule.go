package camTemplate

import (
	"github.com/go-cam/cam/camBase"
)

type CamModule struct {
	Name string                `json:"name"`
	Type camBase.CamModuleType `json:"type"`
}
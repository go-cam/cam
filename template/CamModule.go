package template

import "github.com/go-cam/cam/base/camBase"

type CamModule struct {
	Name string                `json:"name"`
	Type camBase.CamModuleType `json:"type"`
}

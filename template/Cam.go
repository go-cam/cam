package template

import "github.com/go-cam/cam/base/camUtils"

type Cam struct {
	Version string                `json:"version"`
	Modules map[string]*CamModule `json:"modules"`
}

func NewCam() *Cam {
	cam := new(Cam)
	cam.Version = camUtils.C.Version()
	cam.Modules = map[string]*CamModule{}
	return cam
}

func (cam Cam) AddModule(name string, module *CamModule) {
	cam.Modules[name] = module
}

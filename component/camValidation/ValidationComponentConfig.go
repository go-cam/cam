package camValidation

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camConstants"
	"github.com/go-cam/cam/component"
)

// validation component config
type ValidationComponentConfig struct {
	component.ComponentConfig

	// valid mode
	Mode camBase.ValidMode
	// custom valid handler dict
	CustomValidDict map[string]camBase.ValidHandler
	// stop valid when has first error
	StopWhenFirstErr bool
}

// new ValidationComponentConfig instance
func NewValidationConfig() *ValidationComponentConfig {
	conf := new(ValidationComponentConfig)
	conf.Component = &ValidationComponent{}
	conf.Mode = camConstants.ModeInterface
	conf.CustomValidDict = map[string]camBase.ValidHandler{}
	conf.StopWhenFirstErr = true
	return conf
}

// add custom valid handler
func (conf *ValidationComponentConfig) AddValidHandler(name string, handler camBase.ValidHandler) {
	conf.CustomValidDict[name] = handler
}

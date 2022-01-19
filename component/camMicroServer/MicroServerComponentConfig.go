package camMicroServer

import (
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/component"
)

func NewMicroServerComponentConfig(port uint16) *MicroServerComponentConfig {
	c := new(MicroServerComponentConfig)
	c.port = port
	c.Component = new(MicroServerComponent)
	return c
}

type MicroServerComponentConfig struct {
	component.ComponentConfig
	camStatics.IMicroServerComponentConfig
	port uint16
}

func (c *MicroServerComponentConfig) Port() uint16 {
	return c.port
}

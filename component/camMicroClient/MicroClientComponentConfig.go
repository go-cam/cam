package camMicroClient

import "github.com/go-cam/cam/component"

type MicroClientComponentConfig struct {
	component.ComponentConfig
	ServerAddress string
}

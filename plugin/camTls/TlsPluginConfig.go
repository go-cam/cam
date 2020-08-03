package camTls

import (
	"github.com/go-cam/cam/base/camStatics"
)

// tls plugin config
type TlsPluginConfig struct {
	camStatics.PluginConfigInterface

	IsTlsOn     bool   // enable SSL mode
	TlsPort     uint16 // SSL listen port
	TlsCertFile string // absolute cert's filename
	TlsKeyFile  string // absolute key's filename
	TlsOnly     bool   // whether only SSl mode
}

// init config
func (config *TlsPluginConfig) Init() {
	config.IsTlsOn = false
	config.TlsPort = 0
	config.TlsCertFile = ""
	config.TlsKeyFile = ""
	config.TlsOnly = false
}

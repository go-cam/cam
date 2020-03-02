package camPluginSsl

import "github.com/go-cam/cam/base/camBase"

// ssl plugin config
type SslPluginConfig struct {
	camBase.PluginConfigInterface

	IsSslOn     bool   // enable SSL mode
	SslPort     uint16 // SSL listen port
	SslCertFile string // absolute cert's filename
	SslKeyFile  string // absolute key's filename
	SslOnly     bool   // whether only SSl mode
}

// init config
func (config *SslPluginConfig) Init() {
	config.IsSslOn = false
	config.SslPort = 0
	config.SslCertFile = ""
	config.SslKeyFile = ""
	config.SslOnly = false
}

package camConfigs

// SslPlugin config
type SslPlugin struct {
	IsSslOn     bool   // enable SSL mod
	SslPort     uint16 // SSL listen port
	SslCertFile string // absolute cert's filename
	SslKeyFile  string // absolute key's filename
	IsSslOnly   bool   // Whether only SSL mode is enabled
}

// init params
func (plugin *SslPlugin) InitSslPlugin() {
	plugin.IsSslOn = false
	plugin.SslPort = 0
	plugin.SslCertFile = ""
	plugin.SslKeyFile = ""
	plugin.IsSslOnly = false
}

// listen SslPlugin
func (plugin *SslPlugin) ListenSsl(port uint16, certFile string, keyFile string) *SslPlugin {
	plugin.IsSslOn = true
	plugin.SslPort = port
	plugin.SslCertFile = certFile
	plugin.SslKeyFile = keyFile
	return plugin
}

// only SSl mode.
func (plugin *SslPlugin) SslOnly() *SslPlugin {
	plugin.IsSslOnly = true
	return plugin
}

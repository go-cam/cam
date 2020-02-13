package camConfigs

// PluginSsl config
type PluginSsl struct {
	IsSslOn     bool   // enable SSL mod
	SslPort     uint16 // SSL listen port
	SslCertFile string // absolute cert's filename
	SslKeyFile  string // absolute key's filename
	IsSslOnly   bool   // Whether only SSL mode is enabled
}

// 初始化阐述
func (plugin *PluginSsl) InitPluginSsl() {
	plugin.IsSslOn = false
	plugin.SslPort = 0
	plugin.SslCertFile = ""
	plugin.SslKeyFile = ""
	plugin.IsSslOnly = false
}

// listen PluginSsl
func (plugin *PluginSsl) ListenSsl(port uint16, certFile string, keyFile string) *PluginSsl {
	plugin.IsSslOn = true
	plugin.SslPort = port
	plugin.SslCertFile = certFile
	plugin.SslKeyFile = keyFile
	return plugin
}

// only SSl mode.
func (plugin *PluginSsl) SslOnly() *PluginSsl {
	plugin.IsSslOnly = true
	return plugin
}

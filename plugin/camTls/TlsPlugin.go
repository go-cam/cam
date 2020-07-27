package camTls

import "github.com/go-cam/cam/base/camBase"

type TlsPlugin struct {
	camBase.PluginInterface
	conf       *TlsPluginConfig
	handler    func()
	tlsHandler func()
}

func (plugin *TlsPlugin) Init(conf *TlsPluginConfig) {
	plugin.conf = conf
	plugin.handler = func() {}
	plugin.tlsHandler = func() {}
}

func (plugin *TlsPlugin) SetListenHandler(handler func(), tlsHandler func()) {
	plugin.handler = handler
	plugin.tlsHandler = tlsHandler
}

func (plugin *TlsPlugin) StartListenServer() {
	if !plugin.conf.TlsOnly {
		go plugin.handler()
	}
	if plugin.conf.IsTlsOn {
		go plugin.tlsHandler()
	}
}

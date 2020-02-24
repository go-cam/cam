package config

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/camBase"
)

// 获取默认配置
func GetApp() *core.Config {
	config := core.NewConfig()
	config.ComponentDict = map[string]camBase.ConfigComponentInterface{
		"ws":      websocketServer(),
		"http":    httpServer(),
		"db":      core.NewDatabaseConfig("mysql", "127.0.0.1", "3306", "cam", "root", "123456"),
		"console": core.NewConsoleConfig(),
	}
	return config
}

func websocketServer() camBase.ConfigComponentInterface {
	sslCert := core.App.GetEvn("SSL_CERT")
	sslKey := core.App.GetEvn("SSL_KEY")
	return core.NewWebsocketServerConfig(20010).ListenSsl(20011, sslCert, sslKey)
}

func httpServer() camBase.ConfigComponentInterface {
	sslCert := core.App.GetEvn("SSL_CERT")
	sslKey := core.App.GetEvn("SSL_KEY")
	return core.NewHttpServerConfig(20000).SetSessionName("test").ListenSsl(20001, sslCert, sslKey)
}

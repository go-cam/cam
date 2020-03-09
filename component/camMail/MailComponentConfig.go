package camMail

import "github.com/go-cam/cam/base/camBase"

type MailComponentConfig struct {
	camBase.ComponentConfig

	Email              string // master mail
	Nickname           string // nickname
	Password           string // password
	Host               string // Host. mail server Host. Example: "smtp.qq.com"
	Port               uint16 // post
	Ssl                bool   // whether ssl mode
	DefaultContentType string // default: "Content-Type: text/plain; charset=UTF-8"
}

// new mail component config
func NewMailConfig(email string, password string, host string) *MailComponentConfig {
	config := new(MailComponentConfig)
	config.Component = &MailComponent{}
	config.Email = email
	config.Nickname = ""
	config.Password = password
	config.Host = host
	config.Port = 465
	config.Ssl = true
	config.DefaultContentType = "Content-Type: text/plain; charset=UTF-8"
	return config
}

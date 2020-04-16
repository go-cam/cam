package camMail

import (
	"crypto/tls"
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/component"
	"net/smtp"
	"strconv"
)

type MailComponent struct {
	component.Component

	config *MailComponentConfig

	addr string
}

func (comp *MailComponent) Init(configI camBase.ComponentConfigInterface) {
	comp.Component.Init(configI)

	var ok bool
	comp.config, ok = configI.(*MailComponentConfig)
	if !ok {
		camBase.App.Error("MailComponent", "invalid config")
	}

	comp.addr = comp.config.Host + ":" + strconv.FormatInt(int64(comp.config.Port), 10)
}

// on App start
func (comp *MailComponent) Start() {
	comp.Component.Start()
}

// before App destroy
func (comp *MailComponent) Stop() {
	defer comp.Component.Stop()
}

// send mail
func (comp *MailComponent) Send(subject string, body string, to ...string) error {

	// If you want to customize your mail, it is not recommended to refer to the files in cammail. It is recommended to refer to github.com/jordan-wright/email directly
	e := NewEmail()
	e.From = comp.config.Email
	e.To = to
	e.Subject = subject
	e.Text = []byte(body)
	auth := smtp.PlainAuth("", comp.config.Email, comp.config.Password, comp.config.Host)

	if comp.config.Ssl {
		return e.SendWithTLS(comp.addr, auth, &tls.Config{ServerName: comp.config.Host})
	} else {
		return e.Send(comp.addr, auth)
	}
}

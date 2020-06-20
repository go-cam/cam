package camHttp

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camUtils"
	"time"
)

type Store interface {
	SetSessionOption(opt *SessionOption)
	// Create one if not exists
	Get(sessId string) (HttpSessionInterface, error)
	Save(sessI HttpSessionInterface) error
	Del(sessI HttpSessionInterface) error
}

type SessionStoreManager struct {
	store Store
	// session Option
	option *SessionOption
}

func NewSessionStoreManager(store Store, option *SessionOption) *SessionStoreManager {
	m := new(SessionStoreManager)
	m.store = store
	m.option = m.makeSessionOption(option)
	return m
}

// set default value if opt field was empty
func (m *SessionStoreManager) makeSessionOption(opt *SessionOption) *SessionOption {
	if opt.CookieSessionIdName == "" {
		opt.CookieSessionIdName = "SessionID"
	}
	if opt.Expires == 0 {
		opt.Expires = 14 * 24 * time.Hour
	}
	return opt
}

// get session or new by Context
func (m *SessionStoreManager) GetSession(ctx camBase.HttpContextInterface) (HttpSessionInterface, error) {
	sessId := ctx.GetCookieValue(m.option.CookieSessionIdName)
	if sessId == "" {
		sessId = m.generateSessionId()
		ctx.SetCookieValue(m.option.CookieSessionIdName, sessId)
	}
	return m.store.Get(sessId)
}

func (m *SessionStoreManager) generateSessionId() string {
	return camUtils.String.Random(32)
}

type SessionOption struct {
	// session id name stored in cookie
	CookieSessionIdName string
	// session expires
	Expires time.Duration
}

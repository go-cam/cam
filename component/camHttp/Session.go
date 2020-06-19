package camHttp

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camUtils"
)

type Store interface {
	// Create one if not exists
	Get(sessId string) (HttpSessionInterface, error)
	Save(sessI HttpSessionInterface) error
	Del(sessI HttpSessionInterface) error
}

type SessionStoreManager struct {
	store Store
	// session id name stored in cookie
	cookieSessionIdName string
}

func NewSessionStoreManager(cookieSessionIdName string, store Store) *SessionStoreManager {
	m := new(SessionStoreManager)
	m.store = store
	m.cookieSessionIdName = cookieSessionIdName
	return m
}

func (m *SessionStoreManager) GetSession(ctx camBase.HttpContextInterface) (HttpSessionInterface, error) {
	sessId := ctx.GetCookieValue(m.cookieSessionIdName)
	if sessId == "" {
		sessId = m.generateSessionId()
		ctx.SetCookieValue(m.cookieSessionIdName, sessId)
	}
	return m.store.Get(sessId)
}

func (m *SessionStoreManager) generateSessionId() string {
	return camUtils.String.Random(32)
}

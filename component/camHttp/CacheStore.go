package camHttp

import (
	"errors"
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camUtils"
)

type cacheStore struct {
	cachePrefix string
	option      *SessionOption
}

func NewCacheStore(cachePrefix string) *cacheStore {
	store := new(cacheStore)
	store.cachePrefix = cachePrefix
	return store
}

func (store *cacheStore) getCacheKey(sessId string) string {
	return store.cachePrefix + sessId
}

func (store *cacheStore) SetSessionOption(opt *SessionOption) {
	store.option = opt
}

func (store *cacheStore) Get(sessId string) (HttpSessionInterface, error) {
	key := store.getCacheKey(sessId)
	i := camBase.App.GetCache().Get(key)
	if i == nil {
		i = "{}"
	}
	str, ok := i.(string)
	if !ok {
		return nil, errors.New("session's cache value was not type with string")
	}
	var values = map[string]interface{}{}
	camUtils.Json.DecodeToObj([]byte(str), &values)

	sess := newHttpSession(sessId, values)
	sess.SetDestroyHandler(func() {
		err := store.Del(sess)
		if err != nil {
			camBase.App.Error("HttpSession.destroyHandler", "err:"+err.Error())
		}
	})

	return sess, nil
}

func (store *cacheStore) Save(sessI HttpSessionInterface) error {
	key := store.getCacheKey(sessI.GetSessionId())
	str := camUtils.Json.EncodeStr(sessI.Values())
	return camBase.App.GetCache().SetDuration(key, str, store.option.Expires)
}

func (store *cacheStore) Del(sessI HttpSessionInterface) error {
	key := store.getCacheKey(sessI.GetSessionId())
	return camBase.App.GetCache().Del(key)
}

package camMicroServer

import "sync"

type microAppDictDict struct {
	dict *sync.Map
}

func (d *microAppDictDict) getMicroAppInfoDict(name string) *sync.Map {
	dict, _ := d.dict.LoadOrStore(name, &sync.Map{})
	return dict.(*sync.Map)
}

func (d *microAppDictDict) getMicroAppInfo(name string, address string) *microAppInfo {
	serverDict := d.getMicroAppInfoDict(name)
	s, ok := serverDict.Load(address)
	if !ok {
		return nil
	}
	return s.(*microAppInfo)
}

func (d *microAppDictDict) put(microAppInfo *microAppInfo) {
	d.getMicroAppInfoDict(microAppInfo.name).Store(microAppInfo.address, microAppInfo)
}

func (d *microAppDictDict) del(name string, address string) {
	d.getMicroAppInfoDict(name).Delete(address)
}

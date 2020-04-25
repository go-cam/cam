package camCache

// In order to keep the structure as much as possible
type CacheAo struct {
	Value interface{} `json:"v"`
}

func NewCacheAo(value interface{}) *CacheAo {
	ao := new(CacheAo)
	ao.Value = value
	return ao
}

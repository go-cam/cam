package camUtils

import (
	"encoding/json"
)

type JsonUtil struct {
}

var Json = new(JsonUtil)

// encode to bytes
func (util *JsonUtil) Encode(v interface{}) []byte {
	jsonBytes, _ := json.Marshal(v)
	return jsonBytes
}

// encode to bytes and beautify json
func (util *JsonUtil) EncodeBeautiful(v interface{}) []byte {
	jsonBytes, _ := json.MarshalIndent(v, "", "\t")
	return jsonBytes
}

// encode to string
func (util *JsonUtil) EncodeStr(v interface{}) string {
	return string(util.Encode(v))
}

// encode to string and beautify json
func (util *JsonUtil) EncodeStrBeautiful(v interface{}) string {
	return string(util.EncodeBeautiful(v))
}

// decode to struct
func (util *JsonUtil) DecodeToObj(bytes []byte, obj interface{}) {
	_ = json.Unmarshal(bytes, obj)
}

package camUtils

import (
	"encoding/json"
	"github.com/tidwall/gjson"
)

type JsonUtil struct {
}

var Json = new(JsonUtil)

// encode to bytes
func (util *JsonUtil) Encode(v interface{}) []byte {
	jsonBytes, _ := json.Marshal(v)
	return jsonBytes
}

// encode to string
func (util *JsonUtil) EncodeStr(v interface{}) string {
	return string(util.Encode(v))
}

// decode to struct
func (util *JsonUtil) DecodeToObj(bytes []byte, obj interface{}) {
	_ = json.Unmarshal(bytes, obj)
}

// Deprecated:
// decode to gjson sturct
func (util *JsonUtil) DecodeStr(jsonStr string) gjson.Result {
	return gjson.Parse(jsonStr)
}

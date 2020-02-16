package camUtils

import (
	"encoding/json"
	"github.com/tidwall/gjson"
)

type JsonUtil struct {
}

var Json = new(JsonUtil)

// 将数据转为 字节流
func (util *JsonUtil) Encode(v interface{}) []byte {
	jsonBytes, _ := json.Marshal(v)
	return jsonBytes
}

// 将数据转为 字符流
func (util *JsonUtil) EncodeStr(v interface{}) string {
	return string(util.Encode(v))
}

// 将json数据转为对象
func (util *JsonUtil) DecodeToObj(bytes []byte, obj interface{}) {
	_ = json.Unmarshal(bytes, obj)
}

// 解析json字符串
func (util *JsonUtil) DecodeStr(jsonStr string) gjson.Result {
	return gjson.Parse(jsonStr)
}

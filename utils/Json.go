package utils

import (
	"encoding/json"
	"github.com/tidwall/gjson"
)

type jsonUtil struct {

}

var Json = new(jsonUtil)

// 将数据转为 字节流
func (util *jsonUtil) Encode(v interface{}) []byte {
	jsonBytes, _ := json.Marshal(v)
	return jsonBytes
}

// 将数据转为 字符流
func (util *jsonUtil) EncodeStr(v interface{}) string {
	return string(util.Encode(v))
}

// 将json数据转为对象
func (util *jsonUtil) DecodeToObj(bytes []byte, obj interface{}) {
	_ = json.Unmarshal(bytes, obj)
}

// 解析json字符串
func (util *jsonUtil) DecodeStr(jsonStr string) gjson.Result {
	return gjson.Parse(jsonStr)
}
package utils

import "encoding/json"

// 将数据转为 字节流
func Encode(v interface{}) []byte {
	jsonBytes, _ := json.Marshal(v)
	return jsonBytes
}

// 将数据转为 字符流
func EncodeStr(v interface{}) string {
	return string(Encode(v))
}

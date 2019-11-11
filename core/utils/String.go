package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// 字符串工具
type stringUtil struct {
}

// 字符串工具实例
var String = new(stringUtil)

// 生成uuid
func (util *stringUtil) UUID() string {
	timestamp := time.Now().UnixNano()
	hex := fmt.Sprintf("%X", timestamp)
	var splice []string
	splice = append(splice, hex[0:4])
	splice = append(splice, hex[4:8])
	splice = append(splice, hex[8:12])
	splice = append(splice, hex[12:])
	splice = append(splice, util.Random(4))
	splice = append(splice, util.Random(4))
	return strings.Join(splice, "-")
}

// 获取随机字符串
func (util *stringUtil) Random(size int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

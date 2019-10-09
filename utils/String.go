package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// 字符串工具
type stringUtil struct {

}

// 字符串工具实例
var String = new(stringUtil)

// 将 url 转为 驼峰命名。如：get-user-list => GetUserList
func (util *stringUtil) UrlToHump(url string) string {
	words := strings.Split(url, "-")
	hump := ""
	for _, word := range words {
		if len(word) > 0 {
			firstStr := strings.ToUpper(word[0: 1])
			otherStr := word[1:]
			hump += firstStr + otherStr
		}
	}
	return hump
}

// 将驼峰命名转为url模式。如：GetUserList => get-user-list
func (util *stringUtil) HumpToUrl(hump string) string {
	// 将大写字母前面添加 - 。如：GetUserList => -Get-User-List
	reg := regexp.MustCompile(	`([A-Z]])`)
	reg.ReplaceAllString(hump, "-${n}")
	// 将所有字母转为小写
	hump = strings.ToLower(hump)
	// 去除 - 前缀
	hump = strings.TrimPrefix(hump, "-")
	return hump
}

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
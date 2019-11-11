package utils

import (
	"strings"
)

// url 工具
type urlUtil struct {
}

var Url = new(urlUtil)

// 切割url
// 如： /test/test/abc/222?name=aa&age=cc  将切割成： ["test", "test", "abc", "222", "name=aa&age=cc"]
func (util *urlUtil) SplitUrl(url string) []string {
	dirs := strings.Split(url, "/")
	length := len(dirs)
	lastDirs := strings.Split(dirs[length-1], "?")
	// 去收首元素和末元素 （首元素是空字符串，末元素已解析为 lastDirs ）
	dirs = dirs[:length-1]
	dirs = dirs[1:]

	for _, dir := range lastDirs {
		dirs = append(dirs, dir)
	}

	return dirs
}

// 将 url 转为 驼峰命名。如：get-user-list => GetUserList
func (util *urlUtil) UrlToHump(url string) string {
	words := strings.Split(url, "-")
	hump := ""
	for _, word := range words {
		if len(word) > 0 {
			firstStr := strings.ToUpper(word[0:1])
			otherStr := word[1:]
			hump += firstStr + otherStr
		}
	}
	return hump
}

// 将驼峰命名转为url模式。如：GetUserList => get-user-list
func (util *urlUtil) HumpToUrl(hump string) string {
	data := make([]byte, 0, len(hump)*2)
	j := false
	num := len(hump)
	for i := 0; i < num; i++ {
		d := hump[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '-')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

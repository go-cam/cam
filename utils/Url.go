package utils

import (
	"regexp"
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
	lastDirs := strings.Split(dirs[length - 1], "?")
	// 去收首元素和末元素 （首元素是空字符串，末元素已解析为 lastDirs ）
	dirs = dirs[:length - 1]
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
			firstStr := strings.ToUpper(word[0: 1])
			otherStr := word[1:]
			hump += firstStr + otherStr
		}
	}
	return hump
}

// 将驼峰命名转为url模式。如：GetUserList => get-user-list
func (util *urlUtil) HumpToUrl(hump string) string {
	// 将大写字母前面添加 - 。如：GetUserList => -Get-User-List
	reg := regexp.MustCompile(	`([A-Z]])`)
	reg.ReplaceAllString(hump, "-${n}")
	// 将所有字母转为小写
	hump = strings.ToLower(hump)
	// 去除 - 前缀
	hump = strings.TrimPrefix(hump, "-")
	return hump
}
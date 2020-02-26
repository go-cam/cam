package camUtils

import (
	"strings"
)

// utl util
type UrlUtil struct {
}

var Url = new(UrlUtil)

// split url
// Example: /test/test/abc/222?name=aa&age=cc  split to:  ["test", "test", "abc", "222"]
func (util *UrlUtil) SplitUrl(url string) []string {
	// Delete question mark and string after
	questionMarkIndex := strings.Index(url, "?")
	if questionMarkIndex != -1 {
		url = url[:questionMarkIndex]
	}

	dirs := strings.Split(url, "/")
	length := len(dirs)
	lastDirs := strings.Split(dirs[length-1], "?")
	dirs = dirs[:length-1]
	dirs = dirs[1:]

	for _, dir := range lastDirs {
		dirs = append(dirs, dir)
	}

	return dirs
}

// url to Hump.
// Example: get-user-list => GetUserList
func (util *UrlUtil) UrlToHump(url string) string {
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

// Hump to url.
// Example: GetUserList => get-user-list
func (util *UrlUtil) HumpToUrl(hump string) string {
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

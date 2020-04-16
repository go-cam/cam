package camUtils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// string util
type StringUtil struct {
}

var String = new(StringUtil)

// generate UUID
func (util *StringUtil) UUID() string {
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

// get random string
func (util *StringUtil) Random(size int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// underline to hump
// example: get_user_list => GetUserList
func (util *StringUtil) UnderToHump(url string) string {
	words := strings.Split(url, "_")
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

// hump to underline
// example:  GetUserList => get_user_list
func (util *StringUtil) HumpToUnder(hump string) string {
	data := make([]byte, 0, len(hump)*2)
	j := false
	num := len(hump)
	for i := 0; i < num; i++ {
		d := hump[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// Fill the number with 0
// example:
//		FillZero("12", 2) => "12"
//		FillZero("9", 2) => "09"
//		FillZero("129", 2) => "129"
func (util *StringUtil) FillZero(num string, digit int) string {
	for fillNum := len(num) - digit; fillNum > 0; fillNum-- {
		num = "0" + num
	}
	return num
}

// uint64 to string
func (util *StringUtil) Uint64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}

// int64 to string
func (util *StringUtil) Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

// float64 to string
func (util *StringUtil) Float64ToString(num float64) string {
	return strconv.FormatFloat(num, 'f', 16, 64)
}

// string to int64
func (util *StringUtil) StringToInt64(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return num
}

// string to uint64
func (util *StringUtil) StringToUint64(str string) uint64 {
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0
	}
	return num
}

// string to float64
func (util *StringUtil) StringToFloat64(str string) float64 {
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return num
}

package utils

import "time"

type timeUtil struct {
}

var Time = new(timeUtil)

// get now datetime.
// Example: 2006-01-02 15:04:05
func (util *timeUtil) NowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// get now date.
// Example: 2006-01-02
func (util *timeUtil) NowDate() string {
	return time.Now().Format("2006-01-02")
}

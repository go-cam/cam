package camUtils

import "time"

type TimeUtil struct {
}

var Time = new(TimeUtil)

// get now datetime.
// Example: 2006-01-02 15:04:05
func (util *TimeUtil) NowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// get now date.
// Example: 2006-01-02
func (util *TimeUtil) NowDate() string {
	return time.Now().Format("2006-01-02")
}

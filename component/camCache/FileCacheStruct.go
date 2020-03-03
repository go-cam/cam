package camCache

import "time"

// file cache application object
type FileCacheAo struct {
	Value         interface{} `json:"v"`
	DeadTimestamp float64     `json:"dt"` // timestamp in seconds
}

// new object
func NewFileCacheAo(value interface{}, duration time.Duration) *FileCacheAo {
	ao := new(FileCacheAo)
	ao.Value = value
	ao.DeadTimestamp = float64(time.Now().Unix()) + duration.Seconds()
	return ao
}

// is dead
func (ao *FileCacheAo) IsDead() bool {
	return float64(time.Now().Unix()) > ao.DeadTimestamp
}

// is dead. Accurate to nanosecond
func (ao *FileCacheAo) IsDeadNano() bool {
	now := time.Now()

	second := float64(now.Unix())
	nanoSecond := float64(now.Nanosecond()) / 1000000000
	timestamp := second + nanoSecond

	return timestamp > ao.DeadTimestamp
}

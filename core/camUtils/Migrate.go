package camUtils

import (
	"strconv"
	"time"
)

var Migrate = new(migrateUtil)

// migrate util
type migrateUtil struct {
}

// generate migration's file id by datetime using system timezone
func (util *migrateUtil) IdByDatetime() string {
	now := time.Now()
	return now.Format("060102_150405")
}

// generate migration's file id by timestamp
func (util *migrateUtil) IdByTimestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

package camUtils

import (
	"strconv"
	"time"
)

var Migrate = new(MigrateUtil)

// migrate util
type MigrateUtil struct {
}

// generate migration's file id by datetime using system timezone
func (util *MigrateUtil) IdByDatetime() string {
	now := time.Now()
	return now.Format("060102_150405")
}

// generate migration's file id by timestamp
func (util *MigrateUtil) IdByTimestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

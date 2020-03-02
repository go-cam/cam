package camConsole

import "time"

type Migration struct {
	Version string    `xorm:"pk notnull"`
	DoneAt  time.Time `xorm:"timestamp created notnull"`
}

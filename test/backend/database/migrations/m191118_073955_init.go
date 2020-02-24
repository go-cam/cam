package migrations

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/camModels"
)

func init() {
	m := new(m191118_073955_init)
	core.App.AddMigration(m)
}

type m191118_073955_init struct {
	camModels.Migration
}

// up
func (m *m191118_073955_init) Up() {
	m.Exec(`
CREATE TABLE user (
	id INTEGER NOT NULL AUTO_INCREMENT COMMENT 'user id',
	account VARCHAR(50) COMMENT 'user name',
	password VARCHAR(64) COMMENT 'password',
	delete_at timestamp comment 'delete timestamp',
	update_at timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update timestamp',
	create_at timestamp DEFAULT CURRENT_TIMESTAMP COMMENT 'create timestamp',
	PRIMARY KEY(id)
) ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci`)
}

// down
func (m *m191118_073955_init) Down() {
	m.Exec(`DROP TABLE user`)
}

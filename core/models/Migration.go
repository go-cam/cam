package models

import "github.com/cinling/cin/core/base"

// base migration struct
type Migration struct {
	base.MigrationInterface

	sqlList []string
}

// exec sql
func (model *Migration) Exec(sql string) {
	model.sqlList = append(model.sqlList, sql)
}

// get sql list
func (model *Migration) GetSqlList() []string {
	return model.sqlList
}

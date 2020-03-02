package camConsole

import "github.com/go-cam/cam/base/camBase"

// base migration struct
type BaseMigration struct {
	camBase.MigrationInterface

	sqlList []string
}

// exec sql
func (model *BaseMigration) Exec(sql string) {
	model.sqlList = append(model.sqlList, sql)
}

// get sql list
func (model *BaseMigration) GetSqlList() []string {
	return model.sqlList
}

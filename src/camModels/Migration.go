package camModels

import "github.com/go-cam/cam/camBase"

// base migration struct
type Migration struct {
	camBase.MigrationInterface

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

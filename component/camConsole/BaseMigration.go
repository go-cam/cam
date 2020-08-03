package camConsole

import (
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/base/camStructs"
)

// base migration struct
type BaseMigration struct {
	camStatics.MigrationInterface

	sqlList []string
}

// exec sql
func (m *BaseMigration) Exec(sql string) {
	m.sqlList = append(m.sqlList, sql)
}

// get sql list
func (m *BaseMigration) GetSqlList() []string {
	return m.sqlList
}

// TODO Create table
func (m *BaseMigration) CreateTable(tableName string, columns []*camStructs.MysqlColumnBuilder) {

}

// TODO Create index
func (m *BaseMigration) CreateIndex(indexName string, tableName string, columnNames... string) {

}

// TODO Add column for table
func (m *BaseMigration) AddColumn(tableName string, column *camStructs.MysqlColumnBuilder) {

}

// TODO Alter table
func (m *BaseMigration) AlterColumn(tableName string, column *camStructs.MysqlColumnBuilder) {

}

// TODO Drop column
func (m *BaseMigration) DropColumn(tableName string, columnName string) {

}

// TODO Drop index
func (m *BaseMigration) DropIndex(indexName string, tableName string) {}


// TODO Drop table
func (m *BaseMigration) DropTable(tableName string) {

}

func (m *BaseMigration) PrimaryKey() *camStructs.MysqlColumnBuilder {
	col := m.Integer().
}

// type: integer
func (m *BaseMigration) Integer() *camStructs.MysqlColumnBuilder {
	col := camStructs.NewMysqlColumnBuilder()
	col.Type = camStatics.MysqlColumnTypeInteger
	return col
}

package camStructs

import (
	"github.com/go-cam/cam/base/camStatics"
	"strings"
)

type MysqlBuilder struct {

}

func NewMysqlBuilder() *MysqlBuilder {
	m := new(MysqlBuilder)
	return m
}

// Create table
func (m *MysqlBuilder) CreateTable(tableName string, columnList []camStatics.MysqlColumnBuilderInterface, options string) string {
	tpl := "" +
	"CREATE TABLE `${tableName}` (\n" +
	"	${columnRows}\n"+
	") ${options};"

	var columnArr []string
	for _, col := range columnList {
		columnArr = append(columnArr, col.ToSql())
	}
	for _, col := range columnList {
		sql := col.GetKeyPartSql()
		if sql != "" {
			columnArr = append(columnArr, sql)
		}
	}
	columnRows := strings.Join(columnArr, ",\n")

	return m.tplToSql(tpl, map[string]string{
		"tableName": tableName,
		"columnRows": columnRows,
		"options": options,
	})
}

// Rename table
func (m *MysqlBuilder) RenameTable(oldTableName, newTableName string) string {
	return "RENAME TABLE `" + oldTableName + "` TO `" + newTableName + "`;"
}

// Create index
func (m *MysqlBuilder) CreateIndex(indexName, tableName string, columnNames... string) string {
	if len(columnNames) == 0 {
		camStatics.App.Trace("MysqlBuilder.CreateIndex()", "At least one columnName")
	}
	for key, columnName := range columnNames {
		columnNames[key] = "`" + columnName + "`"
	}
	columnStr := strings.Join(columnNames, ",")

	return "ALTER TABLE `" + tableName + "` ADD INDEX `" + indexName + "` (" + columnStr + ");"
}

// Create unique
func (m *MysqlBuilder) CreateUnique(indexName, tableName string, columnNames... string) string {
	if len(columnNames) == 0 {
		camStatics.App.Trace("MysqlBuilder.CreateIndex()", "At least one columnName")
	}
	for key, columnName := range columnNames {
		columnNames[key] = "`" + columnName + "`"
	}
	columnStr := strings.Join(columnNames, ",")

	return "ALTER TABLE `" + tableName + "` ADD UNIQUE `" + indexName + "` (" + columnStr + ");"
}

// Create foreign key
func (m *MysqlBuilder) CreateForeignKey(name, table string, columns []string, refTable string, refColumns []string) string {
	for key, col := range columns {
		columns[key] = "`" + col + "`"
	}
	columnStr := strings.Join(columns, ", ")
	for key, col := range refColumns {
		refColumns[key] = "`" + col + "`"
	}
	refColumnStr := strings.Join(refColumns, ", ")

	return "ALTER TABLE `" + table + "` ADD CONSTRAINT `" + name + "` FOREIGN KEY(" + columnStr + ") REFERENCES `" + refTable + "`(" + refColumnStr + ");"
}

// Add column
func (m *MysqlBuilder) AddColumn(tableName string, column camStatics.MysqlColumnBuilderInterface) string {
	return "ALTER TABLE `" + tableName + "` ADD " + column.ToSql()
}

// Rename column
func (m *MysqlBuilder) RenameColumn(tableName, oldName, newName string) string {
	return "ALTER TABLE `" + tableName + "` RENAME COLUMN `" + oldName + "` TO `" + newName + "`;"
}

// Alter Column
func (m *MysqlBuilder) AlterColumn(tableName, oldName, newName string, columnI camStatics.MysqlColumnBuilderInterface) string {
	column, ok := columnI.(*MysqlColumnBuilder)
	if !ok {
		panic("Invalid column type")
	}
	column.Name = newName
	return "ALTER TABLE `" + tableName + "` CHANGE `" + oldName + "` " + column.ToSql() + ";"
}

// Drop column
func (m *MysqlBuilder) DropColumn(tableName, columnName string) string {
	return "ALTER TABLE `" + tableName + "` DROP COLUMN `" + columnName + "`;"
}

// Drop index
func (m *MysqlBuilder) DropIndex(tableName, indexName string) string {
	return "ALTER TABLE `" + tableName + "` DROP INDEX `" + indexName + "`;"
}

// Drop table
func (m *MysqlBuilder) DropTable(tableName string) string {
	return "DROP TABLE `" + tableName + "`;"
}

// Table options
func (m *MysqlBuilder) Option(comment, engine, charset, collate string) string {
	return "CHARACTER SET " + charset + " COLLATE " + collate + " ENGINE=" + engine + " COMMENT = '" + comment + "'"
}

func (m *MysqlBuilder) tplToSql(tpl string, params map[string]string) string {
	for key, value := range params {
		tpl = strings.Replace(tpl, "${" + key + "}", value, 1)
	}
	return tpl
}

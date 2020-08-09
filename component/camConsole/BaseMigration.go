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
	m.addSql(sql)
}

// get sql list
func (m *BaseMigration) GetSqlList() []string {
	return m.sqlList
}

// Create table
//
// Param options:  a string after create table.
// Example: CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci ENGINE=InnoDB COMMENT = 'table comment'
func (m *BaseMigration) CreateTable(tableName string, columnList []camStatics.MysqlColumnBuilderInterface, options string) {

	sql := camStructs.NewMysqlBuilder().CreateTable(tableName, columnList, options)
	m.addSql(sql)
}

// Table column
// In order to keep table columns sort
func (m *BaseMigration) Column(name string, columnI camStatics.MysqlColumnBuilderInterface) camStatics.MysqlColumnBuilderInterface {
	column, ok := columnI.(*camStructs.MysqlColumnBuilder)
	if !ok {
		panic("Invalid column type")
	}
	column.Name = name
	return column
}

// Rename table
func (m *BaseMigration) RenameTable(oldTableName, newTableName string) {
	sql := camStructs.NewMysqlBuilder().RenameTable(oldTableName, newTableName)
	m.addSql(sql)
}

// Create index
func (m *BaseMigration) CreateIndex(indexName, tableName string, columnNames... string) {
	sql := camStructs.NewMysqlBuilder().CreateIndex(indexName, tableName, columnNames...)
	m.addSql(sql)
}

// Create unique
func (m *BaseMigration) CreateUnique(indexName, tableName string, columnNames... string) {
	sql := camStructs.NewMysqlBuilder().CreateUnique(indexName, tableName, columnNames...)
	m.addSql(sql)
}

// Create foreign key
func (m *BaseMigration) CreateForeignKey(name, table string, columns []string, refTable string, refColumns []string) {
	sql := camStructs.NewMysqlBuilder().CreateForeignKey(name, table, columns, refTable, refColumns)
	m.addSql(sql)
}

// Add column for table
func (m *BaseMigration) AddColumn(tableName, columnName string, columnI camStatics.MysqlColumnBuilderInterface) {
	column, ok := columnI.(*camStructs.MysqlColumnBuilder)
	if !ok {
		panic("Invalid column type")
	}
	column.Name = columnName
	sql := camStructs.NewMysqlBuilder().AddColumn(tableName, column)
	m.addSql(sql)
}

// Rename column
func (m *BaseMigration) RenameColumn(tableName, oldName, newName string) {
	sql := camStructs.NewMysqlBuilder().RenameColumn(tableName, oldName, newName)
	m.addSql(sql)
}

// Alter table
func (m *BaseMigration) AlterColumn(tableName, oldName, newName string, column camStatics.MysqlColumnBuilderInterface) {
	sql := camStructs.NewMysqlBuilder().AlterColumn(tableName, oldName, newName, column)
	m.addSql(sql)
}

// Drop column
func (m *BaseMigration) DropColumn(tableName, columnName string) {
	sql := camStructs.NewMysqlBuilder().DropColumn(tableName, columnName)
	m.addSql(sql)
}

// Drop index
func (m *BaseMigration) DropIndex(tableName, indexName string) {
	sql := camStructs.NewMysqlBuilder().DropIndex(tableName, indexName)
	m.addSql(sql)
}


// Drop table
func (m *BaseMigration) DropTable(tableName string) {
	sql := camStructs.NewMysqlBuilder().DropTable(tableName)
	m.addSql(sql)
}

// Set int type primary key
func (m *BaseMigration) IntPrimaryKey() camStatics.MysqlColumnBuilderInterface {
	return m.Int().Unsigned().NotNull().AutoIncrement().PrimaryKey()
}

// Set bigint type primary key
func (m *BaseMigration) BigintPrimaryKey() camStatics.MysqlColumnBuilderInterface {
	return m.Bigint().Unsigned().NotNull().AutoIncrement().PrimaryKey()
}

func (m *BaseMigration) Tinyint(sizes... int) camStatics.MysqlColumnBuilderInterface {
	col := camStructs.NewMysqlColumnBuilder()
	col.Type = camStatics.MysqlColumnTypeTinyint
	m.autoSetSize(col, sizes)
	return col
}


func (m *BaseMigration) Smallint(sizes... int) camStatics.MysqlColumnBuilderInterface {
	col := camStructs.NewMysqlColumnBuilder()
	col.Type = camStatics.MysqlColumnTypeSmallint
	m.autoSetSize(col, sizes)
	return col
}


func (m *BaseMigration) Mediumint(sizes... int) camStatics.MysqlColumnBuilderInterface {
	col := camStructs.NewMysqlColumnBuilder()
	col.Type = camStatics.MysqlColumnTypeMediumint
	m.autoSetSize(col, sizes)
	return col
}


func (m *BaseMigration) Int(sizes... int) camStatics.MysqlColumnBuilderInterface {
	col := camStructs.NewMysqlColumnBuilder()
	col.Type = camStatics.MysqlColumnTypeInteger
	m.autoSetSize(col, sizes)
	return col
}


func (m *BaseMigration) Bigint(sizes... int) camStatics.MysqlColumnBuilderInterface {
	col := camStructs.NewMysqlColumnBuilder()
	col.Type = camStatics.MysqlColumnTypeBigint
	m.autoSetSize(col, sizes)
	return col
}


func (m *BaseMigration) Float(sizes... int) camStatics.MysqlColumnBuilderInterface {
	col := camStructs.NewMysqlColumnBuilder()
	col.Type = camStatics.MysqlColumnTypeFloat
	m.autoSetSize(col, sizes)
	return col
}


func (m *BaseMigration) Double(sizes... int) camStatics.MysqlColumnBuilderInterface {
	col := camStructs.NewMysqlColumnBuilder()
	col.Type = camStatics.MysqlColumnTypeDouble
	m.autoSetSize(col, sizes)
	return col
}


func (m *BaseMigration) Decimal(sizes... int) camStatics.MysqlColumnBuilderInterface {
	col := camStructs.NewMysqlColumnBuilder()
	col.Type = camStatics.MysqlColumnTypeDecimal
	m.autoSetSize(col, sizes)
	return col
}


func (m *BaseMigration) Varchar(sizes... int) camStatics.MysqlColumnBuilderInterface {
	col := camStructs.NewMysqlColumnBuilder()
	col.Type = camStatics.MysqlColumnTypeVarchar
	m.autoSetSize(col, sizes)
	return col
}


func (m *BaseMigration) Char(sizes... int) camStatics.MysqlColumnBuilderInterface {
	col := camStructs.NewMysqlColumnBuilder()
	col.Type = camStatics.MysqlColumnTypeChar
	m.autoSetSize(col, sizes)
	return col
}


func (m *BaseMigration) Tinytext() camStatics.MysqlColumnBuilderInterface {
	col := camStructs.NewMysqlColumnBuilder()
	col.Type = camStatics.MysqlColumnTypeTinytext
	return col
}


func (m *BaseMigration) Text() camStatics.MysqlColumnBuilderInterface {
	col := camStructs.NewMysqlColumnBuilder()
	col.Type = camStatics.MysqlColumnTypeText
	return col
}


func (m *BaseMigration) Mediumtext() camStatics.MysqlColumnBuilderInterface {
	col := camStructs.NewMysqlColumnBuilder()
	col.Type = camStatics.MysqlColumnTypeMediumtext
	return col
}


func (m *BaseMigration) Longtext() camStatics.MysqlColumnBuilderInterface {
	col := camStructs.NewMysqlColumnBuilder()
	col.Type = camStatics.MysqlColumnTypeLongText
	return col
}


func (m *BaseMigration) Binary(sizes... int) camStatics.MysqlColumnBuilderInterface {
	col := camStructs.NewMysqlColumnBuilder()
	col.Type = camStatics.MysqlColumnTypeBinary
	m.autoSetSize(col, sizes)
	return col
}


func (m *BaseMigration) Varbinary(sizes... int) camStatics.MysqlColumnBuilderInterface {
	col := camStructs.NewMysqlColumnBuilder()
	col.Type = camStatics.MysqlColumnTypeVarbinary
	m.autoSetSize(col, sizes)
	return col
}


func (m *BaseMigration) Tinyblob() camStatics.MysqlColumnBuilderInterface {
	col := camStructs.NewMysqlColumnBuilder()
	col.Type = camStatics.MysqlColumnTypeTinyblob
	return col
}


func (m *BaseMigration) Blob() camStatics.MysqlColumnBuilderInterface {
	col := camStructs.NewMysqlColumnBuilder()
	col.Type = camStatics.MysqlColumnTypeBlob
	return col
}


func (m *BaseMigration) Mediumblob() camStatics.MysqlColumnBuilderInterface {
	col := camStructs.NewMysqlColumnBuilder()
	col.Type = camStatics.MysqlColumnTypeMediumblob
	return col
}


func (m *BaseMigration) Longblob() camStatics.MysqlColumnBuilderInterface {
	col := camStructs.NewMysqlColumnBuilder()
	col.Type = camStatics.MysqlColumnTypeLongblob
	return col
}

func (m *BaseMigration) Datetime() camStatics.MysqlColumnBuilderInterface {
	col := camStructs.NewMysqlColumnBuilder()
	col.Type = camStatics.MysqlColumnTypeDatetime
	return col
}


func (m *BaseMigration) Timestamp() camStatics.MysqlColumnBuilderInterface {
	col := camStructs.NewMysqlColumnBuilder()
	col.Type = camStatics.MysqlColumnTypeTimestamp
	return col
}

func (m *BaseMigration) DefaultOption(comment string) string {
	return camStructs.NewMysqlBuilder().Option(comment, "InnoDB", "utf8mb4", "utf8mb4_unicode_ci")
}

func (m *BaseMigration) Option(comment, engine, charset, collate string) string {
	return camStructs.NewMysqlBuilder().Option(comment, engine, charset, collate)
}

func (m *BaseMigration) autoSetSize(col *camStructs.MysqlColumnBuilder, sizes []int) {
	switch len(sizes) {
	case 0:
	case 1:
		col.Size = sizes[0]
	case 2:
		col.Size = sizes[0]
		col.Size2 = sizes[1]
	default:
		panic("Invalid size nums")
	}
}

func (m *BaseMigration) addSql(sql string) {
	m.sqlList = append(m.sqlList, sql)
}

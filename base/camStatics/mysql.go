package camStatics

const (
	MysqlNullableTypeDefault MysqlNullableType = iota
	MysqlNullableTypeNull
	MysqlNullableTypeNotNull
)

const (
	MysqlColumnTypeTinyint    MysqlColumnType = "TINYINT"
	MysqlColumnTypeSmallint   MysqlColumnType = "SMALLINT"
	MysqlColumnTypeMediumint  MysqlColumnType = "MEDIUMINT"
	MysqlColumnTypeInteger    MysqlColumnType = "INT"
	MysqlColumnTypeBigint     MysqlColumnType = "BIGINT"
	MysqlColumnTypeFloat      MysqlColumnType = "FLOAT"
	MysqlColumnTypeDouble     MysqlColumnType = "DOUBLE"
	MysqlColumnTypeDecimal    MysqlColumnType = "DECIMAL"
	MysqlColumnTypeChar       MysqlColumnType = "CHAR"
	MysqlColumnTypeVarchar    MysqlColumnType = "VARCHAR"
	MysqlColumnTypeTinytext   MysqlColumnType = "TINYTEXT"
	MysqlColumnTypeText       MysqlColumnType = "TEXT"
	MysqlColumnTypeMediumtext MysqlColumnType = "MEDIUMTEXT"
	MysqlColumnTypeLongText   MysqlColumnType = "LONGTEXT"
	MysqlColumnTypeBinary     MysqlColumnType = "BINARY"
	MysqlColumnTypeVarbinary  MysqlColumnType = "VARBINARY"
	MysqlColumnTypeTinyblob   MysqlColumnType = "TINYBLOB"
	MysqlColumnTypeBlob       MysqlColumnType = "BLOB"
	MysqlColumnTypeMediumblob MysqlColumnType = "MEDIUMBLOB"
	MysqlColumnTypeLongblob   MysqlColumnType = "LONGBLOB"
	MysqlColumnTypeDatetime   MysqlColumnType = "DATETIME"
	MysqlColumnTypeTimestamp  MysqlColumnType = "TIMESTAMP"
)

const (
	MysqlKeyTypeNone       MysqlKeyType = ""
	MysqlKeyTypePrimaryKey MysqlKeyType = "PRIMARY KEY"
	MysqlKeyTypeIndex      MysqlKeyType = "KEY"
	MysqlKeyTypeUnique     MysqlKeyType = "UNIQUE"
)

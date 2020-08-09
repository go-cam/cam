package camStructs

import (
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/base/camUtils"
)

// Migration column builder
type MysqlColumnBuilder struct {
	Name          string
	Type          camStatics.MysqlColumnType
	Size          int
	Size2         int
	unsigned      bool
	nullable      camStatics.MysqlNullableType
	defaultValue  interface{}
	autoIncrement bool
	comment       string
	after         string
	keyType       camStatics.MysqlKeyType
}

// New MysqlColumnBuilder
func NewMysqlColumnBuilder() *MysqlColumnBuilder {
	col := new(MysqlColumnBuilder)
	col.unsigned = false
	col.nullable = camStatics.MysqlNullableTypeDefault
	col.Size = -1
	col.Size2 = -1
	col.defaultValue = nil
	col.autoIncrement = false
	col.comment = ""
	col.after = ""
	col.keyType = camStatics.MysqlKeyTypeNone
	return col
}

// Set primary key
func (col *MysqlColumnBuilder) PrimaryKey() camStatics.MysqlColumnBuilderInterface {
	col.keyType = camStatics.MysqlKeyTypePrimaryKey
	return col
}

// Set index key
func (col *MysqlColumnBuilder) Index() camStatics.MysqlColumnBuilderInterface {
	col.keyType = camStatics.MysqlKeyTypeIndex
	return col
}

// Set unique key
func (col *MysqlColumnBuilder) Unique() camStatics.MysqlColumnBuilderInterface {
	col.keyType = camStatics.MysqlKeyTypeUnique
	return col
}

// Set unsigned
func (col *MysqlColumnBuilder) Unsigned() camStatics.MysqlColumnBuilderInterface {
	col.unsigned = true
	return col
}

// Set not nullable
func (col *MysqlColumnBuilder) NotNull() camStatics.MysqlColumnBuilderInterface {
	col.nullable = camStatics.MysqlNullableTypeNotNull
	return col
}

// Set nullable
func (col *MysqlColumnBuilder) Null() camStatics.MysqlColumnBuilderInterface {
	col.nullable = camStatics.MysqlNullableTypeNull
	return col
}

// Set default value
func (col *MysqlColumnBuilder) Default(value interface{}) camStatics.MysqlColumnBuilderInterface {
	col.defaultValue = value
	return col
}

// Set auto-increment
func (col *MysqlColumnBuilder) AutoIncrement() camStatics.MysqlColumnBuilderInterface {
	col.autoIncrement = true
	return col
}

// Set comment
func (col *MysqlColumnBuilder) Comment(comment string) camStatics.MysqlColumnBuilderInterface {
	col.comment = comment
	return col
}

// Set after column
func (col *MysqlColumnBuilder) After(name string) camStatics.MysqlColumnBuilderInterface {
	col.after = name
	return col
}

// To row sql
func (col *MysqlColumnBuilder) ToSql() string {
	return col.namePart() + col.typePart() + col.unsignedPart() + col.nullPart() + col.autoIncrementPart() + col.defaultPart() + col.commentPart() + col.afterPart()
}

// Get key part sql.
// Example: PRIMARY KEY、INDEX
// Only used on Create table
func (col *MysqlColumnBuilder) GetKeyPartSql() string {
	if col.keyType != camStatics.MysqlKeyTypeNone {
		return string(col.keyType) + " (`" + col.Name + "`)"
	}

	return ""
}

func (col *MysqlColumnBuilder) namePart() string {
	return " `" + col.Name + "`"
}

func (col *MysqlColumnBuilder) typePart() string {
	typeStr := " " + string(col.Type)

	if col.Size > 0 && col.Size2 > 0 {
		typeStr += "(" + camUtils.C.IntToString(col.Size) + ", " + camUtils.C.IntToString(col.Size2) + ")"
	} else if col.Size > 0 {
		typeStr += "(" + camUtils.C.IntToString(col.Size) + ")"
	}
	return typeStr
}

func (col *MysqlColumnBuilder) unsignedPart() string {
	if col.unsigned {
		return " UNSIGNED"
	}
	return ""
}

func (col *MysqlColumnBuilder) nullPart() string {
	if col.nullable == camStatics.MysqlNullableTypeNull {
		return " NULL"
	} else if col.nullable == camStatics.MysqlNullableTypeNotNull {
		return " NOT NULL"
	} else {
		return ""
	}
}

func (col *MysqlColumnBuilder) defaultPart() string {
	if col.defaultValue == nil {
		return ""
	}

	switch col.defaultValue.(type) {
	case int:
		return " DEFAULT " + camUtils.C.IntToString(col.defaultValue.(int))
	case int64:
		return " DEFAULT " + camUtils.C.Int64ToString(col.defaultValue.(int64))
	case uint64:
		return " DEFAULT " + camUtils.C.Uint64ToString(col.defaultValue.(uint64))
	case float32:
		return " DEFAULT " + camUtils.C.Float32ToString(col.defaultValue.(float32))
	case float64:
		return " DEFAULT " + camUtils.C.Float64ToString(col.defaultValue.(float64))
	case string:
		return " DEFAULT \"" + col.defaultValue.(string) + "\""
	default:
		panic("Invalid default value type. only support string、int、int64、uint64、float32、float64")
	}
}

func (col *MysqlColumnBuilder) autoIncrementPart() string {
	if col.autoIncrement {
		return " AUTO_INCREMENT"
	}
	return ""
}

func (col *MysqlColumnBuilder) commentPart() string {
	if col.comment != "" {
		return " COMMENT '" + col.comment + "'"
	}
	return ""
}

func (col *MysqlColumnBuilder) afterPart() string {
	if col.after != "" {
		return " AFTER `" + col.after + "`"
	}
	return ""
}

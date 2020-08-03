package camStructs

import "github.com/go-cam/cam/base/camUtils"

type nullableType = int

const(
	nullableTypeDefault nullableType = iota
	nullableTypeNull
	nullableTypeNotNull
)

// Migration column builder
type MysqlColumnBuilder struct {
	Name          string
	Type          string
	Len           uint16
	Unsigned      bool
	Nullable      nullableType
	AutoIncrement bool
	Comment       string
	After         string
}

// new Migration
func NewMysqlColumnBuilder() *MysqlColumnBuilder {
	col := new(MysqlColumnBuilder)
	col.Nullable = nullableTypeDefault
	col.Len = 0
	col.Unsigned = false
	col.Comment = ""
	col.After = ""
	return col
}

// set not null
func (col *MysqlColumnBuilder) NotNull() *MysqlColumnBuilder {
	col.Nullable = false
	return col
}

func (col *MysqlColumnBuilder) Null() *MysqlColumnBuilder {
	col.Nullable = true
	return col
}

func (col *MysqlColumnBuilder) ToSql() string {
	return col.namePart() + col.typePart() + col.unsignedPart() + col.nullPart() + col.autoIncrementPart() + col.commentPart() + col.afterPart()
}

func (col *MysqlColumnBuilder) namePart() string {
	return " `" + col.Name + "`"
}

func (col *MysqlColumnBuilder) typePart() string {
	if col.Len > 0 {
		return " " + col.Type + "(" + camUtils.C.Uint16ToString(col.Len) + ")"
	}
	return " " + col.Type
}

func (col *MysqlColumnBuilder) unsignedPart() string {
	return ""
}

func (col *MysqlColumnBuilder) nullPart() string {
	if col.Nullable {
		return " NULL"
	} else {
		return " NOT NULL"
	}
}

func (col *MysqlColumnBuilder) autoIncrementPart() string {
	if col.AutoIncrement {
		return " AUTO_INCREMENT"
	}
	return ""
}

func (col *MysqlColumnBuilder) commentPart() string {
	if col.Comment != "" {
		return " COMMENT '" + col.Comment + "'"
	}
	return ""
}

func (col *MysqlColumnBuilder) afterPart() string {
	if col.After != "" {
		return " AFTER `" + col.After + "`"
	}
	return ""
}

package camUtils

// Common utils
type CUtil struct {
}

var C = new(CUtil)

// Framework version
func (c *CUtil) Version() string {
	return "v0.5.0-alpha.0-dev"
}

func (c *CUtil) Uint64ToString(num uint64) string {
	return String.Uint64ToString(num)
}

func (c *CUtil) UintToString(num uint) string {
	return c.Uint64ToString(uint64(num))
}

func (c *CUtil) Uint32ToString(num uint32) string {
	return c.Uint64ToString(uint64(num))
}

func (c *CUtil) Uint16ToString(num uint16) string {
	return c.Uint64ToString(uint64(num))
}

func (c *CUtil) Uint8ToString(num uint8) string {
	return c.Uint64ToString(uint64(num))
}

func (c *CUtil) Int64ToString(num int64) string {
	return String.Int64ToString(num)
}

func (c *CUtil) IntToString(num int) string {
	return c.Int64ToString(int64(num))
}

func (c *CUtil) Int32ToString(num int32) string {
	return c.Int64ToString(int64(num))
}

func (c *CUtil) Int16ToString(num int16) string {
	return c.Int64ToString(int64(num))
}

func (c *CUtil) Int8ToString(num int8) string {
	return c.Int64ToString(int64(num))
}

func (c *CUtil) Float64ToString(num float64) string {
	return String.Float64ToString(num)
}

func (c *CUtil) Float32ToString(num float32) string {
	return c.Float64ToString(float64(num))
}

func (c *CUtil) InterfaceToString(v interface{}) string {
	switch v.(type) {
	case float32:
		return c.Float32ToString(v.(float32))
	case float64:
		return c.Float64ToString(v.(float64))
	case int64:
		return c.Int64ToString(v.(int64))
	case int32:
		return c.Int32ToString(v.(int32))
	case int16:
		return c.Int16ToString(v.(int16))
	case int8:
		return c.Int8ToString(v.(int8))
	case int:
		return c.IntToString(v.(int))
	case uint64:
		return c.Uint64ToString(v.(uint64))
	case uint32:
		return c.Uint32ToString(v.(uint32))
	case uint16:
		return c.Uint16ToString(v.(uint16))
	case uint8:
		return c.Uint8ToString(v.(uint8))
	case uint:
		return c.UintToString(v.(uint))
	case bool:
		str := "false"
		if v.(bool) {
			str = "true"
		}
		return str
	case []byte:
		return string(v.([]byte))
	case string:
		return v.(string)
	default:
		return ""
	}
}

func (c *CUtil) StringToInt64(str string) int64 {
	return String.StringToInt64(str)
}

func (c *CUtil) StringToInt32(str string) int32 {
	return int32(c.StringToInt64(str))
}

func (c *CUtil) StringToInt16(str string) int16 {
	return int16(c.StringToInt64(str))
}

func (c *CUtil) StringToInt8(str string) int8 {
	return int8(c.StringToInt64(str))
}

func (c *CUtil) StringToInt(str string) int {
	return int(c.StringToInt64(str))
}

func (c *CUtil) StringToUint64(str string) uint64 {
	return String.StringToUint64(str)
}

func (c *CUtil) StringToUint32(str string) uint32 {
	return uint32(c.StringToUint64(str))
}

func (c *CUtil) StringToUint16(str string) uint16 {
	return uint16(c.StringToUint64(str))
}

func (c *CUtil) StringToUint8(str string) uint8 {
	return uint8(c.StringToUint64(str))
}

func (c *CUtil) StringToUint(str string) uint {
	return uint(c.StringToUint64(str))
}

func (c *CUtil) StringToFloat64(str string) float64 {
	return String.StringToFloat64(str)
}

func (c *CUtil) StringToFloat32(str string) float32 {
	return float32(c.StringToFloat64(str))
}

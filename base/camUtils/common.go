package camUtils

import "strconv"

// Common utils
type CUtil struct {
}

var C = new(CUtil)

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

func (c *CUtil) StringToInt64(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return num
}

func (c *CUtil) StringToInt(str string) int {
	return int(c.StringToInt64(str))
}

func (c *CUtil) StringToUint64(str string) uint64 {
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0
	}
	return num
}

func (c *CUtil) StringToUint(str string) uint {
	return uint(c.StringToUint64(str))
}

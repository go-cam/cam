package camUtils

import "reflect"

// reflect tool
type ReflectUtil struct {
}

var Reflect = new(ReflectUtil)

// get struct name
func (util *ReflectUtil) GetStructName(v interface{}) string {
	t := reflect.TypeOf(v)
	return t.Elem().Name()
}

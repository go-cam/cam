package camUtils

import "reflect"

// reflect tool
type ReflectUtil struct {
}

var Reflect = new(ReflectUtil)

// get struct name
func (util *ReflectUtil) GetStructName(i interface{}) string {
	t := reflect.TypeOf(i)
	return t.Elem().Name()
}

// reflect to element.
// If reflect value is point, change it to element
func (util *ReflectUtil) ValueOfElem(i interface{}) reflect.Value {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}

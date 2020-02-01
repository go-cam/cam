package camUtils

import "reflect"

// 反射工具实例
var Reflect = new(reflectUtil)

// 反射工具
type reflectUtil struct {
}

// 通过反射获取类的名字
func (util *reflectUtil) GetStructName(v interface{}) string {
	t := reflect.TypeOf(v)
	return t.Elem().Name()
}

package camUtils

var Error = new(ErrorUtil)

// 错误处理工具
type ErrorUtil struct {
}

// 如果错误不为 nil 则抛出异常
func (util *ErrorUtil) Panic(err error) {
	if err != nil {
		panic(err)
	}
}

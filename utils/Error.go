package utils

var Error = new(errorUtil)

// 错误处理工具
type errorUtil struct {
}

// 如果错误不为 nil 则抛出异常
func (util *errorUtil) Panic(err error) {
	if err != nil {
		panic(err)
	}
}
package camUtils

var Error = new(ErrorUtil)

// error util
type ErrorUtil struct {
}

// call panic() if error not nil
// Deprecated: Remove on v0.5.x
func (util *ErrorUtil) Panic(err error) {
	if err != nil {
		panic(err)
	}
}

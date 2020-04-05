package camValidation

import (
	"errors"
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camUtils"
	"reflect"
	"regexp"
)

type ValidRule struct {
}

var Rule = new(ValidRule)

// valid email
func (comp *ValidRule) Email(value reflect.Value) error {
	email := value.String()

	reg := regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
	if !reg.MatchString(email) {
		return errors.New("invalid email")
	}
	return nil
}

// length limit. only support string value
func (comp *ValidRule) Length(min int, max int) camBase.ValidHandler {
	return func(value reflect.Value) error {
		if value.Kind() != reflect.String {
			return errors.New("length only support string type. not support <" + value.Type().String() + " Type>")
		}
		strLen := len(value.String())
		if strLen < min {
			return errors.New("less then " + camUtils.C.IntToString(min) + " character")
		}
		if strLen > max {
			return errors.New("more then " + camUtils.C.IntToString(max) + " character")
		}

		return nil
	}
}

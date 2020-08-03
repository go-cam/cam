package camValidation

import (
	"errors"
	"github.com/go-cam/cam/base/camStatics"
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

// limit string length
// length limit. only support string value
func (comp *ValidRule) Length(min int, max int) camStatics.ValidHandler {
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

// Limit value maximum
func (comp *ValidRule) Max(max float64) camStatics.ValidHandler {
	return func(value reflect.Value) error {
		switch value.Kind() {
		case reflect.Float32:
			fallthrough
		case reflect.Float64:
			f := value.Float()
			if f > max {
				return errors.New("more then " + camUtils.C.Float64ToString(max))
			}
		case reflect.Int8:
			fallthrough
		case reflect.Int16:
			fallthrough
		case reflect.Int32:
			fallthrough
		case reflect.Int:
			fallthrough
		case reflect.Int64:
			i := value.Int()
			if i > int64(max) {
				return errors.New("more then " + camUtils.C.Float64ToString(max))
			}
		case reflect.Uint8:
			fallthrough
		case reflect.Uint16:
			fallthrough
		case reflect.Uint32:
			fallthrough
		case reflect.Uint:
			fallthrough
		case reflect.Uint64:
			i := value.Uint()
			if i > uint64(max) {
				return errors.New("more then " + camUtils.C.Float64ToString(max))
			}
		case reflect.String:
			f := camUtils.C.StringToFloat64(value.String())
			if f > max {
				return errors.New("more then " + camUtils.C.Float64ToString(max))
			}
		}

		return nil
	}
}

// Limit value minimum
func (comp *ValidRule) Min(min float64) camStatics.ValidHandler {
	return func(value reflect.Value) error {
		switch value.Kind() {
		case reflect.Float32:
			fallthrough
		case reflect.Float64:
			f := value.Float()
			if f < min {
				return errors.New("less then " + camUtils.C.Float64ToString(min))
			}
		case reflect.Int8:
			fallthrough
		case reflect.Int16:
			fallthrough
		case reflect.Int32:
			fallthrough
		case reflect.Int:
			fallthrough
		case reflect.Int64:
			i := value.Int()
			if i < int64(min) {
				return errors.New("less then " + camUtils.C.Float64ToString(min))
			}
		case reflect.Uint8:
			fallthrough
		case reflect.Uint16:
			fallthrough
		case reflect.Uint32:
			fallthrough
		case reflect.Uint:
			fallthrough
		case reflect.Uint64:
			i := value.Uint()
			if i < uint64(min) {
				return errors.New("less then " + camUtils.C.Float64ToString(min))
			}
		case reflect.String:
			f := camUtils.C.StringToFloat64(value.String())
			if f < min {
				return errors.New("less then " + camUtils.C.Float64ToString(min))
			}
		}

		return nil
	}
}

// number range
// support string and number of float or int
func (comp *ValidRule) Range(min float64, max float64) camStatics.ValidHandler {
	return func(value reflect.Value) error {
		switch value.Kind() {
		case reflect.Float32:
			fallthrough
		case reflect.Float64:
			f := value.Float()
			if f < min {
				return errors.New("less then " + camUtils.C.Float64ToString(min))
			} else if f > max {
				return errors.New("more then " + camUtils.C.Float64ToString(max))
			}
		case reflect.Int8:
			fallthrough
		case reflect.Int16:
			fallthrough
		case reflect.Int32:
			fallthrough
		case reflect.Int:
			fallthrough
		case reflect.Int64:
			i := value.Int()
			if i < int64(min) {
				return errors.New("less then " + camUtils.C.Float64ToString(min))
			} else if i > int64(max) {
				return errors.New("more then " + camUtils.C.Float64ToString(max))
			}
		case reflect.Uint8:
			fallthrough
		case reflect.Uint16:
			fallthrough
		case reflect.Uint32:
			fallthrough
		case reflect.Uint:
			fallthrough
		case reflect.Uint64:
			i := value.Uint()
			if i < uint64(min) {
				return errors.New("less then " + camUtils.C.Float64ToString(min))
			} else if i > uint64(max) {
				return errors.New("more then " + camUtils.C.Float64ToString(max))
			}
		case reflect.String:
			f := camUtils.C.StringToFloat64(value.String())
			if f < min {
				return errors.New("less then " + camUtils.C.Float64ToString(min))
			} else if f > max {
				return errors.New("more then " + camUtils.C.Float64ToString(max))
			}
		}

		return nil
	}
}

// value can not be empty
func (comp *ValidRule) Require(value reflect.Value) error {
	switch value.Kind() {
	case reflect.Slice:
		fallthrough
	case reflect.Map:
		fallthrough
	case reflect.String:
		if value.Len() == 0 {
			return errors.New("can not be empty")
		}
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		if value.Float() == 0 {
			return errors.New("can not be zero")
		}
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int:
		fallthrough
	case reflect.Int64:
		if value.Int() == 0 {
			return errors.New("can not be zero")
		}
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint:
		fallthrough
	case reflect.Uint64:
		if value.Uint() == 0 {
			return errors.New("can not be zero")
		}
	case reflect.Chan:
		fallthrough
	case reflect.Func:
		fallthrough
	case reflect.Interface:
		fallthrough
	case reflect.Ptr:
		if value.IsNil() {
			return errors.New("can not be nil")
		}
	default:
		return errors.New("not support type")
	}

	return nil
}

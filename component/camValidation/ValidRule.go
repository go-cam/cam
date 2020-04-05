package camValidation

import (
	"errors"
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

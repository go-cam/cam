package camValidation

import (
	"errors"
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camConstants"
	"github.com/go-cam/cam/component"
	"reflect"
)

// validation component
type ValidationComponent struct {
	camBase.ValidationComponentInterface
	component.Component

	conf      *ValidationComponentConfig
	validDict map[string]camBase.ValidHandler
}

// init
func (comp *ValidationComponent) Init(conf camBase.ComponentConfigInterface) {
	comp.Component.Init(conf)

	var ok bool
	comp.conf, ok = conf.(*ValidationComponentConfig)
	if !ok {
		camBase.App.Fatal("ValidationComponent", "invalid config")
		return
	}

	comp.initValidDict()
}

// start
func (comp *ValidationComponent) Start() {
	comp.Component.Start()
}

// stop
func (comp *ValidationComponent) Stop() {
	defer comp.Component.Stop()
}

// init valid dict
func (comp *ValidationComponent) initValidDict() {
	comp.validDict = map[string]camBase.ValidHandler{
		"email": Rule.Email,
	}
}

// valid struct
func (comp *ValidationComponent) Valid(v interface{}) map[string][]error {
	errs := map[string][]error{}

	if comp.conf.Mode == camConstants.ModeInterface {
		errs = comp.validInterface(v)
	}

	return errs
}

// valid interface
func (comp *ValidationComponent) validInterface(v interface{}) map[string][]error {
	errDict := map[string][]error{}

	valid, ok := v.(camBase.ValidInterface)
	if !ok {
		// not implement camBase.ValidInterface, pass
		return errDict
	}
	rules := valid.Rules()
	rValue := reflect.ValueOf(v)
	if rValue.Kind() == reflect.Ptr {
		rValue = rValue.Elem()
	}
	if rValue.Kind() != reflect.Struct {
		comp.addError(&errDict, "", errors.New("value of validation is not a struct"))
		return errDict
	}

	for _, rule := range rules {
		fieldNames := rule.Fields()
		handlers := rule.Handlers()

		for _, fieldName := range fieldNames {
			field := rValue.FieldByName(fieldName)

			for _, handler := range handlers {
				err := handler(field)
				if err != nil {
					errMsg := fieldName + ": " + err.Error()
					comp.addError(&errDict, fieldName, errors.New(errMsg))
					if comp.conf.StopWhenFirstErr {
						return errDict
					}
				}
			}

			// validation each child member
			if comp.conf.Each {
				errs := comp.eachValid(&field)
				for _, err := range errs {
					comp.addError(&errDict, fieldName, err)
				}
			}
		}
	}

	return errDict
}

// get valid handler list
func (comp *ValidationComponent) getValidHandlers(names []string) ([]camBase.ValidHandler, error) {
	var validHandlers []camBase.ValidHandler

	for _, name := range names {
		validHandler, has := comp.validDict[name]
		if !has {
			return nil, errors.New(name + ": invalid ValidHandler name")
		}
		validHandlers = append(validHandlers, validHandler)
	}

	return validHandlers, nil
}

// add error handler
func (comp *ValidationComponent) addError(errDict *map[string][]error, fieldName string, err error) {
	_, has := (*errDict)[fieldName]
	if !has {
		(*errDict)[fieldName] = []error{}
	}
	(*errDict)[fieldName] = append((*errDict)[fieldName], err)
}

// validation children
func (comp *ValidationComponent) eachValid(value *reflect.Value) []error {
	var errs []error

	var field reflect.Value
	if (*value).Kind() == reflect.Ptr {
		field = (*value).Elem()
	} else {
		field = *value
	}

	switch field.Kind() {
	case reflect.Slice:
		length := field.Len()
		for i := 0; i < length; i++ {
			childField := field.Index(i)
			comp.eachValid(&childField)
		}
	case reflect.Map:
		keyValues := field.MapKeys()
		for _, keyValue := range keyValues {
			childField := field.MapIndex(keyValue)
			comp.eachValid(&childField)
		}
	case reflect.Struct:
		errDict := comp.Valid(field)
		for _, childErrs := range errDict {
			errs = append(errs, childErrs...)
		}
	}

	return errs
}

package validator

import (
	"database/sql/driver"
	"fmt"
	"reflect"
)

type RequiredValidator struct {
	definition RequiredValidatorDefinition
}

type RequiredValidatorDefinition struct {
	Required []string `json:"pattern"`
}

type RequiredValidationError struct {
	Input      interface{}                 `json:"input"`
	Definition RequiredValidatorDefinition `json:"definition"`
}

func (r RequiredValidationError) Error() string {
	return fmt.Sprintf("input sturct does not satisfy required values '%v'\n", r.Definition.Required)
}

func NewRequiredValidator(definition RequiredValidatorDefinition) (RequiredValidator, error) {
	required := definition.Required
	len := len(required)
	if len == 0 {
		return RequiredValidator{}, EmptyError{"the required value should have at least one element"}
	}

	for i := 0; i < len-1; i++ {
		key := required[i]
		for j := i + 1; j < len; j++ {
			if required[j] == key {
				return RequiredValidator{}, DuplicationError{"the required value should not be duplicated"}
			}
		}
	}

	return RequiredValidator{definition}, nil
}

func (r RequiredValidator) Validate(input interface{}) error {
	var v reflect.Value
	if reflect.TypeOf(input).Kind() != reflect.Ptr {
		v = reflect.ValueOf(input)
	} else {
		v = reflect.ValueOf(input).Elem()
	}
	if v.Kind() != reflect.Struct {
		return &InvalidFieldTypeError{
			Definition: r.definition,
			Input:      input,
		}
	}

	for _, key := range r.definition.Required {
		e := v.FieldByName(key)
		if !e.IsValid() {
			return &InvalidFieldTypeError{
				Definition: r.definition,
				Input:      input,
			}
		}
		n, ok := e.Interface().(driver.Valuer)
		if !ok {
			return &InvalidFieldTypeError{
				Definition: r.definition,
				Input:      input,
			}
		}
		v, err := n.Value()
		if err != nil {
			return &InvalidFieldTypeError{
				Definition: r.definition,
				Input:      input,
			}
		}
		if v == nil {
			return &RequiredValidationError{
				Definition: r.definition,
				Input:      input,
			}
		}
	}

	return nil
}

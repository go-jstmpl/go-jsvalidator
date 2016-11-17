package validator

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/gocraft/dbr"
)

var (
	RequiredDefinitionEmptyError       = errors.New("the required value should have at least one element")
	RequiredDefinitionDuplicationError = errors.New("the required value should not be duplicated")
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
	return fmt.Sprintf("input struct does not satisfy required values '%v'\n", r.Definition.Required)
}

func NewRequiredValidator(definition RequiredValidatorDefinition) (RequiredValidator, error) {
	required := definition.Required
	len := len(required)
	if len == 0 {
		return RequiredValidator{}, RequiredDefinitionEmptyError
	}

	for i := 0; i < len-1; i++ {
		key := required[i]
		for j := i + 1; j < len; j++ {
			if required[j] == key {
				return RequiredValidator{}, RequiredDefinitionDuplicationError
			}
		}
	}

	return RequiredValidator{definition}, nil
}

func (r RequiredValidator) Validate(input interface{}) error {
	if input == nil {
		return &InvalidTypeError{
			Definition: r.definition,
			Input:      input,
		}
	}

	var v reflect.Value
	if reflect.TypeOf(input).Kind() != reflect.Ptr {
		v = reflect.ValueOf(input)
	} else {
		v = reflect.ValueOf(input).Elem()
	}
	if v.Kind() != reflect.Struct {
		return &InvalidTypeError{
			Definition: r.definition,
			Input:      input,
		}
	}

	for _, key := range r.definition.Required {
		e := v.FieldByName(key)
		if e.Kind() == reflect.Ptr {
			if e.IsNil() {
				return &RequiredValidationError{
					Definition: r.definition,
					Input:      input,
				}
			}
			e = v.FieldByName(key).Elem()
		}

		i := e.Interface()
		switch j := i.(type) {
		case dbr.NullString:
			if !j.Valid {
				return &RequiredValidationError{
					Definition: r.definition,
					Input:      input,
				}
			}
			continue
		case dbr.NullTime:
			if !j.Valid {
				return &RequiredValidationError{
					Definition: r.definition,
					Input:      input,
				}
			}
			continue
		case dbr.NullInt64:
			if !j.Valid {
				return &RequiredValidationError{
					Definition: r.definition,
					Input:      input,
				}
			}
			continue
		case dbr.NullFloat64:
			if !j.Valid {
				return &RequiredValidationError{
					Definition: r.definition,
					Input:      input,
				}
			}
			continue
		case dbr.NullBool:
			if !j.Valid {
				return &RequiredValidationError{
					Definition: r.definition,
					Input:      input,
				}
			}
			continue
		default:
			continue
		}
	}
	return nil
}

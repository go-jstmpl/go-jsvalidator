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

// Validate reports whether input is valid against required keys.
// Fields of Input struct must be exported.
func (r RequiredValidator) Validate(input interface{}) error {
	v, ok := convertToConcreteValue(reflect.ValueOf(input))
	if !ok {
		return &InvalidTypeError{
			Definition: r.definition,
			Input:      input,
		}
	}
	if v.Kind() != reflect.Struct {
		return &InvalidTypeError{
			Definition: r.definition,
			Input:      input,
		}
	}
	for _, key := range r.definition.Required {
		e, ok := getFieldByName(v, key)
		if !ok {
			return &InvalidFieldTypeError{
				Definition: r.definition,
				Input:      input,
			}
		}
		c, ok := convertToConcreteValue(e)
		if !ok {
			return &InvalidFieldTypeError{
				Definition: r.definition,
				Input:      input,
			}
		}
		i := c.Interface()
		ok = isValid(i)
		if !ok {
			return &RequiredValidationError{
				Definition: r.definition,
				Input:      input,
			}
		}
	}
	return nil
}

// convertToConcreteValue returns a concrete value that stored in the pointer.
// The ok return value reports whether conversion was successful.
func convertToConcreteValue(input reflect.Value) (value reflect.Value, ok bool) {
	if input.Kind() != reflect.Ptr {
		return input, true
	}
	if input.IsNil() {
		return reflect.Value{}, false
	}
	return input.Elem(), true
}

// getFieldByName returns the struct field with the given name.
// The ok return value reports whether field with key was found in value.
func getFieldByName(v reflect.Value, key string) (f reflect.Value, ok bool) {
	field := v.FieldByName(key)
	if (field == reflect.Value{}) {
		return reflect.Value{}, false
	}
	return field, true
}

// isValid reports whether i is valid.
// The argument i will always convert to dbr.Null* type or primitive type
// by type switch in Validate of RequiredValidator.
func isValid(i interface{}) (ok bool) {
	switch j := i.(type) {
	case dbr.NullString:
		return j.Valid
	case dbr.NullInt64:
		return j.Valid
	case dbr.NullFloat64:
		return j.Valid
	case dbr.NullBool:
		return j.Valid
	case dbr.NullTime:
		return j.Valid
	default:
		return true
	}
}

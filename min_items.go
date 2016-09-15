package validator

import (
	"fmt"
	"reflect"
)

type MinItemsValidator struct {
	definition MinItemsValidatorDefinition
}

type MinItemsValidatorDefinition struct {
	MinItems int `json:"Min_items"`
}

type MinItemsValidationError struct {
	Definition MinItemsValidatorDefinition `json:"definition"`
	Input      interface{}                 `json:"input"`
}

func (i MinItemsValidationError) Error() string {
	l, _ := i.Input.([]interface{})
	return fmt.Sprintf("the number of input items should be less than MinItems:'%d' but actual value of input has '%d' items",
		i.Definition.MinItems, len(l))
}

func NewMinItemsValidator(definition MinItemsValidatorDefinition) (MinItemsValidator, error) {
	if definition.MinItems < 0 {
		return MinItemsValidator{}, NoLengthError{"the value of MinItems should be greater than, or equal to, 0"}
	}
	return MinItemsValidator{definition}, nil
}

func (i MinItemsValidator) Validate(input interface{}) error {
	switch reflect.TypeOf(input).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(input)
		if s.Len() >= i.definition.MinItems {
			return nil
		}
	default:
		return TypeError{"input should be slice"}
	}

	return &MinItemsValidationError{
		i.definition,
		input,
	}
}

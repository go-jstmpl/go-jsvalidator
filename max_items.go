package validator

import (
	"fmt"
	"reflect"
)

type MaxItemsValidator struct {
	definition MaxItemsValidatorDefinition
}

type MaxItemsValidatorDefinition struct {
	MaxItems int `json:"max_items"`
}

type MaxItemsValidationError struct {
	Definition MaxItemsValidatorDefinition `json:"definition"`
	Input      interface{}                 `json:"input"`
}

func (i MaxItemsValidationError) Error() string {
	l, _ := i.Input.([]interface{})
	return fmt.Sprintf("the number of input items should be less than MaxItems:'%d' but actual value of input has '%d' items",
		i.Definition.MaxItems, len(l))
}

func NewMaxItemsValidator(definition MaxItemsValidatorDefinition) (MaxItemsValidator, error) {
	if definition.MaxItems < 0 {
		return MaxItemsValidator{}, NoLengthError{"the value of maxItems should be greater than, or equal to, 0"}
	}
	return MaxItemsValidator{definition}, nil
}

func (i MaxItemsValidator) Validate(input interface{}) error {
	switch reflect.TypeOf(input).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(input)
		if s.Len() <= i.definition.MaxItems {
			return nil
		}
	default:
		return TypeError{"input should be slice"}
	}

	return &MaxItemsValidationError{
		i.definition,
		input,
	}
}

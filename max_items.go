package validator

import (
	"fmt"
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
	switch input.(type) {
	case []int:
		l, _ := input.([]int)
		if len(l) <= i.definition.MaxItems {
			return nil
		}
	case []string:
		l, _ := input.([]string)
		if len(l) <= i.definition.MaxItems {
			return nil
		}
	case []float64:
		l, _ := input.([]float64)
		if len(l) <= i.definition.MaxItems {
			return nil
		}
	}

	return &MaxItemsValidationError{
		i.definition,
		input,
	}
}

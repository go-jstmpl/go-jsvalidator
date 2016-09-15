package validator

import (
	"fmt"
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
	switch input.(type) {
	case []int:
		l, _ := input.([]int)
		if len(l) >= i.definition.MinItems {
			return nil
		}
	case []string:
		l, _ := input.([]string)
		if len(l) >= i.definition.MinItems {
			return nil
		}
	case []float64:
		l, _ := input.([]float64)
		if len(l) >= i.definition.MinItems {
			return nil
		}
	}

	return &MinItemsValidationError{
		i.definition,
		input,
	}
}

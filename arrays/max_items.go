package arrays

import (
	"errors"
	"fmt"
)

var MaxItemsDefinitionNoLengthError = errors.New("the value of MaxItems should be greater than or equal to 0")

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

func (err MaxItemsValidationError) Error() string {
	return fmt.Sprintf("the length of %v should be less than %d",
		err.Input, err.Definition.MaxItems)
}

func NewMaxItemsValidator(definition MaxItemsValidatorDefinition) (MaxItemsValidator, error) {
	if definition.MaxItems < 0 {
		return MaxItemsValidator{}, MaxItemsDefinitionNoLengthError
	}
	return MaxItemsValidator{definition}, nil
}

func (i MaxItemsValidator) Validate(input interface{}) error {
	slice, err := toSlice(input)
	if err != nil {
		return err
	}

	if len(slice) <= i.definition.MaxItems {
		return nil
	}
	return &MaxItemsValidationError{
		i.definition,
		input,
	}
}

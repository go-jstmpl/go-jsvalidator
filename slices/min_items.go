package slices

import "fmt"

type MinItemsValidator struct {
	definition MinItemsValidatorDefinition
}

type MinItemsValidatorDefinition struct {
	MinItems int `json:"min_items"`
}

type MinItemsValidationError struct {
	Definition MinItemsValidatorDefinition `json:"definition"`
	Input      interface{}                 `json:"input"`
}

func (err MinItemsValidationError) Error() string {
	return fmt.Sprintf("the length of %v should be less than %d",
		err.Input, err.Definition.MinItems)
}

func NewMinItemsValidator(definition MinItemsValidatorDefinition) (MinItemsValidator, error) {
	if definition.MinItems < 0 {
		return MinItemsValidator{}, &NoLengthError{}
	}
	return MinItemsValidator{definition}, nil
}

func (i MinItemsValidator) Validate(input interface{}) error {
	slice, err := toSlice(input)
	if err != nil {
		return err
	}

	if len(slice) >= i.definition.MinItems {
		return nil
	}
	return &MinItemsValidationError{
		i.definition,
		input,
	}
}

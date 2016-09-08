package validator

import (
	"fmt"
)

type IntMinItemsValidator struct {
	definition IntMinItemsValidatorDefinition
}

type IntMinItemsValidatorDefinition struct {
	MinItems int `json:"max_items"`
}

type IntMinItemsValidationError struct {
	Definition IntMinItemsValidatorDefinition `json:"definition"`
	Input      []int                          `json:"input"`
}

func (i IntMinItemsValidationError) Error() string {
	return fmt.Sprintf("the number of input items should be greater than MinItems:'%d' but actual value of input has '%d' items \n",
		i.Definition.MinItems, len(i.Input))
}

func NewIntMinItemsValidator(definition IntMinItemsValidatorDefinition) (IntMinItemsValidator, error) {
	if definition.MinItems < 0 {
		return IntMinItemsValidator{}, NoLengthError{"the value of maxItems should be greater than, or equal to, 0"}
	}
	return IntMinItemsValidator{definition}, nil
}

func (i IntMinItemsValidator) Validate(input []int) error {
	if len(input) >= i.definition.MinItems {
		return nil
	}
	return &IntMinItemsValidationError{
		i.definition,
		input,
	}
}

type FloatMinItemsValidator struct {
	definition FloatMinItemsValidatorDefinition
}

type FloatMinItemsValidatorDefinition struct {
	MinItems int `json:"max_items"`
}

type FloatMinItemsValidationError struct {
	Definition FloatMinItemsValidatorDefinition `json:"definition"`
	Input      []float64                        `json:"input"`
}

func (f FloatMinItemsValidationError) Error() string {
	return fmt.Sprintf("the number of input items should be greater than MinItems:'%d' but actual value of input has '%d' items \n",
		f.Definition.MinItems, len(f.Input))
}

func NewFloatMinItemsValidator(definition FloatMinItemsValidatorDefinition) (FloatMinItemsValidator, error) {
	if definition.MinItems < 0 {
		return FloatMinItemsValidator{}, NoLengthError{"the value of maxItems should be greater than, or equal to, 0"}
	}
	return FloatMinItemsValidator{definition}, nil
}

func (i FloatMinItemsValidator) Validate(input []float64) error {
	if len(input) >= i.definition.MinItems {
		return nil
	}
	return &FloatMinItemsValidationError{
		i.definition,
		input,
	}
}

type StringMinItemsValidator struct {
	definition StringMinItemsValidatorDefinition
}

type StringMinItemsValidatorDefinition struct {
	MinItems int `json:"max_items"`
}

type StringMinItemsValidationError struct {
	Definition StringMinItemsValidatorDefinition `json:"definition"`
	Input      []string                          `json:"input"`
}

func (s StringMinItemsValidationError) Error() string {
	return fmt.Sprintf("the number of input items should be greater than MinItems:'%d' but actual value of input has '%d' items \n",
		s.Definition.MinItems, len(s.Input))
}

func NewStringMinItemsValidator(definition StringMinItemsValidatorDefinition) (StringMinItemsValidator, error) {
	if definition.MinItems < 0 {
		return StringMinItemsValidator{}, NoLengthError{"the value of maxItems should be greater than, or equal to, 0"}
	}

	return StringMinItemsValidator{definition}, nil
}

func (i StringMinItemsValidator) Validate(input []string) error {
	if len(input) >= i.definition.MinItems {
		return nil
	}
	return &StringMinItemsValidationError{
		i.definition,
		input,
	}
}

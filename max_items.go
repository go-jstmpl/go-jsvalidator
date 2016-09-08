package validator

import (
	"fmt"
)

type IntMaxItemsValidator struct {
	definition IntMaxItemsValidatorDefinition
}

type IntMaxItemsValidatorDefinition struct {
	MaxItems int `json:"max_items"`
}

type IntMaxItemsValidationError struct {
	Definition IntMaxItemsValidatorDefinition `json:"definition"`
	Input      []int                          `json:"input"`
}

func (i IntMaxItemsValidationError) Error() string {
	return fmt.Sprintf("the number of input items should be less than MaxItems:'%d' but actual value of input has '%d' items",
		i.Definition.MaxItems, len(i.Input))
}

func NewIntMaxItemsValidator(definition IntMaxItemsValidatorDefinition) (IntMaxItemsValidator, error) {
	if definition.MaxItems < 0 {
		return IntMaxItemsValidator{}, NoLengthError{"the value of maxItems should be greater than, or equal to, 0"}
	}
	return IntMaxItemsValidator{definition}, nil
}

func (i IntMaxItemsValidator) Validate(input []int) error {
	if len(input) <= i.definition.MaxItems {
		return nil
	}
	return &IntMaxItemsValidationError{
		i.definition,
		input,
	}
}

type FloatMaxItemsValidator struct {
	definition FloatMaxItemsValidatorDefinition
}

type FloatMaxItemsValidatorDefinition struct {
	MaxItems int `json:"max_items"`
}

type FloatMaxItemsValidationError struct {
	Definition FloatMaxItemsValidatorDefinition `json:"definition"`
	Input      []float64                        `json:"input"`
}

func (f FloatMaxItemsValidationError) Error() string {
	return fmt.Sprintf("the number of input items should be less than MaxItems:'%d' but actual value of input has '%d' items",
		f.Definition.MaxItems, len(f.Input))
}

func NewFloatMaxItemsValidator(definition FloatMaxItemsValidatorDefinition) (FloatMaxItemsValidator, error) {
	if definition.MaxItems < 0 {
		return FloatMaxItemsValidator{}, NoLengthError{"the value of maxItems should be greater than, or equal to, 0"}
	}
	return FloatMaxItemsValidator{definition}, nil
}

func (i FloatMaxItemsValidator) Validate(input []float64) error {
	if len(input) <= i.definition.MaxItems {
		return nil
	}
	return &FloatMaxItemsValidationError{
		i.definition,
		input,
	}
}

type StringMaxItemsValidator struct {
	definition StringMaxItemsValidatorDefinition
}

type StringMaxItemsValidatorDefinition struct {
	MaxItems int `json:"max_items"`
}

type StringMaxItemsValidationError struct {
	Definition StringMaxItemsValidatorDefinition `json:"definition"`
	Input      []string                          `json:"input"`
}

func (s StringMaxItemsValidationError) Error() string {
	return fmt.Sprintf("the number of input items should be less than MaxItems:'%d' but actual value of input has '%d' items",
		s.Definition.MaxItems, len(s.Input))
}

func NewStringMaxItemsValidator(definition StringMaxItemsValidatorDefinition) (StringMaxItemsValidator, error) {
	if definition.MaxItems < 0 {
		return StringMaxItemsValidator{}, NoLengthError{"the value of maxItems should be greater than, or equal to, 0"}
	}

	return StringMaxItemsValidator{definition}, nil
}

func (i StringMaxItemsValidator) Validate(input []string) error {
	if len(input) <= i.definition.MaxItems {
		return nil
	}
	return &StringMaxItemsValidationError{
		i.definition,
		input,
	}
}

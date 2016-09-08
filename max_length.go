package validator

import (
	"fmt"
	"unicode/utf8"
)

type MaxLengthValidator struct {
	definition MaxLengthValidatorDefinition
}

type MaxLengthValidatorDefinition struct {
	MaxLength int `json:"max_length"`
}

type MaxLengthValidationError struct {
	Definition MaxLengthValidatorDefinition `json:"definition"`
	Input      string                       `json:"input"`
}

func (m MaxLengthValidationError) Error() string {
	return fmt.Sprintf("should be less than, or equal to, %d charactors but actual value has %d charactors\n",
		m.Definition.MaxLength, utf8.RuneCountInString(m.Input))
}

func NewMaxLengthValidator(definition MaxLengthValidatorDefinition) (MaxLengthValidator, error) {
	if definition.MaxLength < 0 {
		return MaxLengthValidator{}, NoLengthError{"the max length should be greater than, or equal to, 0"}
	}
	return MaxLengthValidator{definition}, nil
}

func (m MaxLengthValidator) Validate(input string) error {
	if utf8.RuneCountInString(input) <= m.definition.MaxLength {
		return nil
	}
	return &MaxLengthValidationError{
		m.definition,
		input,
	}
}

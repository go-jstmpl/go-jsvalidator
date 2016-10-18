package strings

import (
	"fmt"
	"unicode/utf8"
)

type MinLengthValidator struct {
	definition MinLengthValidatorDefinition
}

type MinLengthValidatorDefinition struct {
	MinLength int `json:"min_length"`
}

type MinLengthValidationError struct {
	Definition MinLengthValidatorDefinition `json:"definition"`
	Input      string                       `json:"input"`
}

func (m MinLengthValidationError) Error() string {
	return fmt.Sprintf("should be greater than, or equal to, %d charactors but actual value has %d charactors\n",
		m.Definition.MinLength, utf8.RuneCountInString(m.Input))
}

func NewMinLengthValidator(definition MinLengthValidatorDefinition) (MinLengthValidator, error) {
	if definition.MinLength < 0 {
		return MinLengthValidator{}, NoLengthError{"the minlength should be greater than, or equal to, 0"}
	}

	return MinLengthValidator{definition}, nil
}

func (m MinLengthValidator) Validate(input string) error {
	if utf8.RuneCountInString(input) >= m.definition.MinLength {
		return nil
	}
	return &MinLengthValidationError{
		m.definition,
		input,
	}
}

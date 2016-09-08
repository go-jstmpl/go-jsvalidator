package validator

import (
	"fmt"
)

type IntMinimumValidator struct {
	definition IntMinimumValidatorDefinition
}

type IntMinimumValidatorDefinition struct {
	Minimum   int  `json:"minimum"`
	Exclusive bool `json:"exclusive"`
}

type IntMinimumValidationError struct {
	Definition IntMinimumValidatorDefinition `json:"definition"`
	Input      int                           `json:"input"`
}

func (i IntMinimumValidationError) Error() string {
	return fmt.Sprintf("should be greater than %d but actual value is %d with option exlusive %t\n",
		i.Definition.Minimum, i.Input, i.Definition.Exclusive)
}

func NewIntMinimumValidator(definition IntMinimumValidatorDefinition) (IntMinimumValidator, error) {
	return IntMinimumValidator{definition}, nil
}

func (m IntMinimumValidator) Validate(input int) error {
	if !m.definition.Exclusive {
		if input >= m.definition.Minimum {
			return nil
		}
		return &IntMinimumValidationError{
			m.definition,
			input,
		}
	}
	if input > m.definition.Minimum {
		return nil
	}
	return &IntMinimumValidationError{
		m.definition,
		input,
	}
}

type FloatMinimumValidator struct {
	definition FloatMinimumValidatorDefinition
}

type FloatMinimumValidatorDefinition struct {
	Minimum   float64 `json:"minimum"`
	Exclusive bool    `json:"exclusive"`
}

type FloatMinimumValidationError struct {
	Definition FloatMinimumValidatorDefinition `json:"definition"`
	Input      float64                         `json:"input"`
}

func (f FloatMinimumValidationError) Error() string {
	return fmt.Sprintf("should be greater than %g but actual value is %g with option exlusive %t\n",
		f.Definition.Minimum, f.Input, f.Definition.Exclusive)
}

func NewFloatMinimumValidator(definition FloatMinimumValidatorDefinition) (FloatMinimumValidator, error) {
	return FloatMinimumValidator{definition}, nil
}

func (m FloatMinimumValidator) Validate(input float64) error {
	if !m.definition.Exclusive {
		if input >= m.definition.Minimum {
			return nil
		}
		return &FloatMinimumValidationError{
			m.definition,
			input,
		}
	}

	if input > m.definition.Minimum {
		return nil
	}
	return &FloatMinimumValidationError{
		m.definition,
		input,
	}
}

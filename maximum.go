package validator

import (
	"fmt"
)

type IntMaximumValidator struct {
	definition IntMaximumValidatorDefinition
}

type IntMaximumValidatorDefinition struct {
	Maximum   int  `json:"maximum"`
	Exclusive bool `json:"exclusive"`
}

type IntMaximumValidationError struct {
	Definition IntMaximumValidatorDefinition `json:"definition"`
	Input      int                           `json:"input"`
}

func (i IntMaximumValidationError) Error() string {
	return fmt.Sprintf("should be less than %d but actual value is %d with option exlusive %t\n",
		i.Definition.Maximum, i.Input, i.Definition.Exclusive)
}

func NewIntMaximumValidator(definition IntMaximumValidatorDefinition) (IntMaximumValidator, error) {
	return IntMaximumValidator{definition}, nil
}

func (m IntMaximumValidator) Validate(input int) error {
	if !m.definition.Exclusive {
		if input <= m.definition.Maximum {
			return nil
		}
		return &IntMaximumValidationError{
			m.definition,
			input,
		}
	}

	if input < m.definition.Maximum {
		return nil
	}
	return &IntMaximumValidationError{
		m.definition,
		input,
	}
}

type FloatMaximumValidator struct {
	definition FloatMaximumValidatorDefinition
}

type FloatMaximumValidatorDefinition struct {
	Maximum   float64 `json:"maximum"`
	Exclusive bool    `json:"exclusive"`
}

type FloatMaximumValidationError struct {
	Definition FloatMaximumValidatorDefinition `json:"definition"`
	Input      float64                         `json:"input"`
}

func (f FloatMaximumValidationError) Error() string {
	return fmt.Sprintf("should be less than %g but actual value is %g with option exlusive %t\n",
		f.Definition.Maximum, f.Input, f.Definition.Exclusive)
}

func NewFloatMaximumValidator(definition FloatMaximumValidatorDefinition) (FloatMaximumValidator, error) {
	return FloatMaximumValidator{definition}, nil
}

func (m FloatMaximumValidator) Validate(input float64) error {
	if !m.definition.Exclusive {
		if input <= m.definition.Maximum {
			return nil
		}
		return &FloatMaximumValidationError{
			m.definition,
			input,
		}
	}
	if input < m.definition.Maximum {
		return nil
	}
	return &FloatMaximumValidationError{
		m.definition,
		input,
	}
}

package numbers

import "fmt"

type MinimumValidatorDefinition struct {
	Minimum   float64 `json:"minimum"`
	Exclusive bool    `json:"exclusive"`
}

type MinimumValidationError struct {
	Definition MinimumValidatorDefinition `json:"definition"`
	Input      float64                    `json:"input"`
}

func (err MinimumValidationError) Error() string {
	if err.Definition.Exclusive {
		return fmt.Sprintf("the value %f should be greater than %f", err.Input, err.Definition.Minimum)
	}
	return fmt.Sprintf("the value %f should be greater than or equal to %f", err.Input, err.Definition.Minimum)
}

type MinimumValidator struct {
	definition MinimumValidatorDefinition
}

func NewMinimumValidator(definition MinimumValidatorDefinition) (MinimumValidator, error) {
	return MinimumValidator{definition}, nil
}

func (m MinimumValidator) Validate(input float64) error {
	if !m.definition.Exclusive {
		if input >= m.definition.Minimum {
			return nil
		}
		return &MinimumValidationError{
			m.definition,
			input,
		}
	}

	if input > m.definition.Minimum {
		return nil
	}
	return &MinimumValidationError{
		m.definition,
		input,
	}
}

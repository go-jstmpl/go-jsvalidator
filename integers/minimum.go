package integers

import "fmt"

type MinimumValidatorDefinition struct {
	Minimum   int  `json:"minimum"`
	Exclusive bool `json:"exclusive"`
}

type MinimumValidationError struct {
	Definition MinimumValidatorDefinition `json:"definition"`
	Input      int                        `json:"input"`
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

func (m MinimumValidator) Validate(input int) error {
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

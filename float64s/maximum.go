package float64s

import "fmt"

type MaximumValidator struct {
	definition MaximumValidatorDefinition
}

type MaximumValidatorDefinition struct {
	Maximum   float64 `json:"maximum"`
	Exclusive bool    `json:"exclusive"`
}

type MaximumValidationError struct {
	Definition MaximumValidatorDefinition `json:"definition"`
	Input      float64                    `json:"input"`
}

func (err MaximumValidationError) Error() string {
	if err.Definition.Exclusive {
		return fmt.Sprintf("the value %f should be less than %f", err.Input, err.Definition.Maximum)
	}
	return fmt.Sprintf("the value %f should be less than or equal to %f", err.Input, err.Definition.Maximum)
}

func NewMaximumValidator(definition MaximumValidatorDefinition) (MaximumValidator, error) {
	return MaximumValidator{definition}, nil
}

func (m MaximumValidator) Validate(input float64) error {
	if !m.definition.Exclusive {
		if input <= m.definition.Maximum {
			return nil
		}
		return &MaximumValidationError{
			m.definition,
			input,
		}
	}

	if input < m.definition.Maximum {
		return nil
	}
	return &MaximumValidationError{
		m.definition,
		input,
	}
}

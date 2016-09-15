package validator

import (
	"fmt"
)

type MinimumValidator struct {
	definition MinimumValidatorDefinition
}

type MinimumValidatorDefinition struct {
	Minimum   interface{} `json:"minimum"`
	Exclusive bool        `json:"exclusive"`
}

type MinimumValidationError struct {
	Definition MinimumValidatorDefinition `json:"definition"`
	Input      interface{}                `json:"input"`
}

func (i MinimumValidationError) Error() string {
	return fmt.Sprintf("should be greater than %+v but actual value is %+v with option exlusive %t\n",
		i.Definition.Minimum, i.Input, i.Definition.Exclusive)
}

func NewMinimumValidator(definition MinimumValidatorDefinition) (MinimumValidator, error) {
	return MinimumValidator{definition}, nil
}

func (m MinimumValidator) Validate(input interface{}) error {
	switch input := input.(type) {
	case int:
		min, ok := m.definition.Minimum.(int)
		if !ok {
			return TypeError{"input and maximum should be same type"}
		}

		if !m.definition.Exclusive {
			if input >= min {
				return nil
			}
			return &MinimumValidationError{
				m.definition,
				input,
			}
		}
		if input > min {
			return nil
		}
	case float64:
		min, ok := m.definition.Minimum.(float64)
		if !ok {
			return TypeError{"input and maximum should be same type"}
		}

		if !m.definition.Exclusive {
			if input >= min {
				return nil
			}
			return &MinimumValidationError{
				m.definition,
				input,
			}
		}
		if input > min {
			return nil
		}
	default:
		return TypeError{"should be int or float64"}
	}

	return &MinimumValidationError{
		m.definition,
		input,
	}
}

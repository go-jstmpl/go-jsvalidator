package validator

import (
	"fmt"
)

type MaximumValidator struct {
	definition MaximumValidatorDefinition
}

type MaximumValidatorDefinition struct {
	Maximum   interface{} `json:"maximum"`
	Exclusive bool        `json:"exclusive"`
}

type MaximumValidationError struct {
	Definition MaximumValidatorDefinition `json:"definition"`
	Input      interface{}                `json:"input"`
}

func (i MaximumValidationError) Error() string {
	return fmt.Sprintf("should be less than %+v but actual value is %+v with option exlusive %t\n",
		i.Definition.Maximum, i.Input, i.Definition.Exclusive)
}

func NewMaximumValidator(definition MaximumValidatorDefinition) (MaximumValidator, error) {
	return MaximumValidator{definition}, nil
}

func (m MaximumValidator) Validate(input interface{}) error {
	switch input := input.(type) {
	case int:
		max, ok := m.definition.Maximum.(int)
		if !ok {
			return TypeError{"input and maximum should be same type"}
		}

		if !m.definition.Exclusive {
			if input <= max {
				return nil
			}
			return &MaximumValidationError{
				m.definition,
				input,
			}
		}

		if input < max {
			return nil
		}
	case float64:
		max, ok := m.definition.Maximum.(float64)
		if !ok {
			return TypeError{"input and maximum should be same type"}
		}

		if !m.definition.Exclusive {
			if input <= max {
				return nil
			}
			return &MaximumValidationError{
				m.definition,
				input,
			}
		}

		if input < max {
			return nil
		}
	default:
		return TypeError{"should be int or float64"}
	}

	return &MaximumValidationError{
		m.definition,
		input,
	}
}

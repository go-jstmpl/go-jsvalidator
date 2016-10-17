package validator

import (
	"fmt"
)

type EnumValidator struct {
	definition EnumValidatorDefinition
}

type EnumValidatorDefinition struct {
	Enumerate interface{} `json:"enum"`
}

type EnumValidationError struct {
	Definition EnumValidatorDefinition `json:"definition"`
	Input      interface{}             `json:"input"`
}

func (i EnumValidationError) Error() string {
	return fmt.Sprintf("input value '%d' does not listed in enumerate '%v'\n", i.Input, i.Definition.Enumerate)
}

func NewEnumValidator(definition EnumValidatorDefinition) (EnumValidator, error) {
	enumerate := definition.Enumerate
	switch enum := enumerate.(type) {
	case []int:
		len := len(enum)
		if len == 0 {
			return EnumValidator{}, EmptyError{"the enumerate should have at least one element"}
		}

		for i := 0; i < len-1; i++ {
			e := enum[i]
			for j := i + 1; j < len; j++ {
				if enum[j] == e {
					return EnumValidator{}, DuplicationError{"the elements of enumerate should not be duplicated"}
				}
			}
		}
	case []float64:
		len := len(enum)
		if len == 0 {
			return EnumValidator{}, EmptyError{"the enumerate should have at least one element"}
		}

		for i := 0; i < len-1; i++ {
			e := enum[i]
			for j := i + 1; j < len; j++ {
				if enum[j] == e {
					return EnumValidator{}, DuplicationError{"the elements of enumerate should not be duplicated"}
				}
			}
		}
	case []string:
		len := len(enum)
		if len == 0 {
			return EnumValidator{}, EmptyError{"the enumerate should have at least one element"}
		}

		for i := 0; i < len-1; i++ {
			e := enum[i]
			for j := i + 1; j < len; j++ {
				if enum[j] == e {
					return EnumValidator{}, DuplicationError{"the elements of enumerate should not be duplicated"}
				}
			}
		}
	default:
		return EnumValidator{}, TypeError{"the enumerate should be []int, []float64 or []string"}
	}

	return EnumValidator{definition}, nil

}

func (i EnumValidator) Validate(input interface{}) error {
	switch input := input.(type) {
	case int:
		enum, ok := i.definition.Enumerate.([]int)
		if !ok {
			return TypeError{"input and element of enumerate should be same type"}
		}
		for _, e := range enum {
			if input == e {
				return nil
			}
		}
	case float64:
		enum, ok := i.definition.Enumerate.([]float64)
		if !ok {
			return TypeError{"input and element of enumerate should be same type"}
		}
		for _, e := range enum {
			if input == e {
				return nil
			}
		}
	case string:
		enum, ok := i.definition.Enumerate.([]string)
		if !ok {
			return TypeError{"input and element of enumerate should be same type"}
		}
		for _, e := range enum {
			if input == e {
				return nil
			}
		}
	default:
		return TypeError{"input should be same type as element of enumerate"}
	}

	return &EnumValidationError{
		i.definition,
		input,
	}
}

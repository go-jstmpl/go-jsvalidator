package validator

import (
	"fmt"
)

type IntEnumValidator struct {
	definition IntEnumValidatorDefinition
}

type IntEnumValidatorDefinition struct {
	Enumerate []int `json:"enum"`
}

type IntEnumValidationError struct {
	Definition IntEnumValidatorDefinition `json:"definition"`
	Input      int                        `json:"input"`
}

func (i IntEnumValidationError) Error() string {
	return fmt.Sprintf("input value '%d' does not listed in enumerate '%v'\n", i.Input, i.Definition.Enumerate)
}

func NewIntEnumValidator(definition IntEnumValidatorDefinition) (IntEnumValidator, error) {
	enumerate := definition.Enumerate
	len := len(enumerate)
	if len == 0 {
		return IntEnumValidator{}, EmptyError{"the enumerate should have at least one element"}
	}

	for i := 0; i < len-1; i++ {
		e := enumerate[i]
		for j := i + 1; j < len; j++ {
			if enumerate[j] == e {
				return IntEnumValidator{}, DuplicationError{"the elements of enumerate should not be duplicated"}
			}
		}
	}

	return IntEnumValidator{definition}, nil

}

func (i IntEnumValidator) Validate(input int) error {
	for _, e := range i.definition.Enumerate {
		if input == e {
			return nil
		}
	}
	return &IntEnumValidationError{
		i.definition,
		input,
	}
}

type StringEnumValidator struct {
	definition StringEnumValidatorDefinition
}

type StringEnumValidatorDefinition struct {
	Enumerate []string `json:"enum"`
}

type StringEnumValidationError struct {
	Definition StringEnumValidatorDefinition `json:"definition"`
	Input      string                        `json:"input"`
}

func (s StringEnumValidationError) Error() string {
	return fmt.Sprintf("input value '%s' does not listed in enumerate '%v' \n", s.Input, s.Definition.Enumerate)
}

func NewStringEnumValidator(definition StringEnumValidatorDefinition) (StringEnumValidator, error) {
	enumerate := definition.Enumerate
	len := len(enumerate)
	if len == 0 {
		return StringEnumValidator{}, EmptyError{"the enumerate should have at least one element"}
	}

	for i := 0; i < len-1; i++ {
		e := enumerate[i]
		for j := i + 1; j < len; j++ {
			if enumerate[j] == e {
				return StringEnumValidator{}, DuplicationError{"the elements of enumerate should not be duplicated"}
			}
		}
	}

	return StringEnumValidator{definition}, nil
}

func (s StringEnumValidator) Validate(input string) error {
	for _, e := range s.definition.Enumerate {
		if input == e {
			return nil
		}
	}
	return &StringEnumValidationError{
		s.definition,
		input,
	}
}

type FloatEnumValidator struct {
	definition FloatEnumValidatorDefinition
}

type FloatEnumValidatorDefinition struct {
	Enumerate []float64 `json:"enum"`
}

type FloatEnumValidationError struct {
	Definition FloatEnumValidatorDefinition `json:"definition"`
	Input      float64                      `json:"input"`
}

func (s FloatEnumValidationError) Error() string {
	return fmt.Sprintf("input value '%g' does not listed in enumerate '%v'\n", s.Input, s.Definition.Enumerate)
}

func NewFloatEnumValidator(definition FloatEnumValidatorDefinition) (FloatEnumValidator, error) {
	enumerate := definition.Enumerate
	len := len(enumerate)
	if len == 0 {
		return FloatEnumValidator{}, EmptyError{"the enumerate should have at least one element"}
	}

	for i := 0; i < len-1; i++ {
		e := enumerate[i]
		for j := i + 1; j < len; j++ {
			if enumerate[j] == e {
				return FloatEnumValidator{}, DuplicationError{"the elements of enumerate should not be duplicated"}
			}
		}
	}
	return FloatEnumValidator{definition}, nil

}

func (f FloatEnumValidator) Validate(input float64) error {
	for _, e := range f.definition.Enumerate {
		if input == e {
			return nil
		}
	}
	return &FloatEnumValidationError{
		f.definition,
		input,
	}
}

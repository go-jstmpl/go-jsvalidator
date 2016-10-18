package integers

import (
	"errors"
	"fmt"
)

var (
	EnumDefinitionEmptyError       = errors.New("the Enum should have at least one element")
	EnumDefinitionDuplicationError = errors.New("the elements of Enum shouldn't be duplicated")
)

type EnumValidator struct {
	definition EnumValidatorDefinition
}

type EnumValidatorDefinition struct {
	Enum []int `json:"enum"`
}

type EnumValidationError struct {
	Definition EnumValidatorDefinition `json:"definition"`
	Input      int                     `json:"input"`
}

func (err EnumValidationError) Error() string {
	return fmt.Sprintf("input value %d doesn't exist in %v", err.Input, err.Definition.Enum)
}

func NewEnumValidator(def EnumValidatorDefinition) (EnumValidator, error) {
	len := len(def.Enum)
	if len == 0 {
		return EnumValidator{}, EnumDefinitionEmptyError
	}

	for i := 0; i < len-1; i++ {
		e := def.Enum[i]
		for j := i + 1; j < len; j++ {
			if def.Enum[j] == e {
				return EnumValidator{}, EnumDefinitionDuplicationError
			}
		}
	}

	return EnumValidator{def}, nil
}

func (v EnumValidator) Validate(input int) error {
	for _, e := range v.definition.Enum {
		if input == e {
			return nil
		}
	}
	return &EnumValidationError{
		v.definition,
		input,
	}
}
